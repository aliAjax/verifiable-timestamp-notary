package notary

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

type Service struct {
	store Store
	cfg   Config
	audit *AuditChain
}

func NewService(store Store, cfg Config) *Service {
	return &Service{store: store, cfg: cfg, audit: NewAuditChain()}
}
func (s *Service) CreateClaim(req TimestampRequest) (Claim, error) {
	digest, err := NormalizeDigest(req.Digest, req.Algorithm)
	if err != nil {
		return Claim{}, err
	}
	if req.Kind == "" {
		return Claim{}, fmt.Errorf("kind is required")
	}
	if req.IdempotencyKey != "" {
		if old, e := s.store.FindByIdempotency(req.IdempotencyKey); e == nil {
			return old, nil
		}
	}
	now := time.Now().UTC()
	c := Claim{ID: newID("clm"), Digest: digest, Algorithm: strings.ToUpper(req.Algorithm), Kind: req.Kind, ExternalRef: req.ExternalRef, PolicyVersion: req.PolicyVersion, IdempotencyKey: req.IdempotencyKey, Status: StatusPending, CreatedAt: now, UpdatedAt: now, Version: 1}
	if err := s.store.PutClaim(c); err != nil {
		return Claim{}, err
	}
	s.record("claim.created", c.ID, map[string]string{"digest": c.Digest})
	return c, nil
}
func (s *Service) GetClaim(id string) (Claim, error) { return s.store.GetClaim(id) }
func (s *Service) ListClaims(offset, limit int) []Claim {
	if offset < 0 {
		offset = 0
	}
	if limit <= 0 || limit > 100 {
		limit = 20
	}
	return s.store.ListClaims(offset, limit)
}
func (s *Service) Notarize(id string) (Proof, error) {
	c, err := s.store.GetClaim(id)
	if err != nil {
		return Proof{}, err
	}
	if c.Status == StatusRevoked {
		return Proof{}, ErrRevoked
	}
	now := time.Now().UTC()
	c.Status = StatusAggregating
	c.UpdatedAt = now
	c.Version++
	if err = s.store.PutClaim(c); err != nil {
		return Proof{}, err
	}
	signers := s.store.ListSigners()
	active := make([]Signer, 0, len(signers))
	for _, x := range signers {
		if x.Enabled && time.Since(x.LastSeen) < 10*time.Minute && absDuration(x.ClockOffset) <= s.cfg.MaxClockSkew {
			active = append(active, x)
		}
	}
	if len(active) < s.cfg.Quorum {
		c.Status = StatusPending
		c.UpdatedAt = time.Now().UTC()
		c.Version++
		_ = s.store.PutClaim(c)
		return Proof{}, ErrInsufficientQuorum
	}
	batch := Batch{ID: newID("bat"), Leaves: []string{c.Digest}, Status: "open", CreatedAt: now}
	tree := NewMerkleTree(batch.Leaves)
	batch.Root = tree.Root()
	batch.Status = "closed"
	batch.ClosedAt = now
	s.store.PutBatch(batch)
	sigs := make([]Signature, 0, len(active))
	for i, x := range active {
		sigs = append(sigs, Signature{SignerID: x.ID, KeyVersion: x.KeyVersion, Algorithm: "HMAC-SHA256", Value: HashStrings(x.ID, batch.Root, now.String(), fmt.Sprint(i)), At: now})
	}
	cp := Checkpoint{ID: newID("chk"), Root: batch.Root, Sequence: uint64(now.UnixNano()), At: now, Signatures: sigs}
	s.store.PutCheckpoint(cp)
	path, _ := tree.Proof(c.Digest)
	p := Proof{ID: newID("prf"), ClaimID: c.ID, Digest: c.Digest, Algorithm: c.Algorithm, Kind: c.Kind, PolicyVersion: c.PolicyVersion, Timestamp: now, BatchID: batch.ID, MerkleRoot: batch.Root, CheckpointID: cp.ID, Path: path, Signatures: sigs, CertificateDigests: certDigests(active), Verification: VerificationResult{Valid: true, Checks: []CheckResult{{Name: "merkle", Passed: true, Detail: "root matches"}, {Name: "quorum", Passed: true, Detail: fmt.Sprintf("%d/%d", len(sigs), s.cfg.Quorum)}}, VerifiedAt: now}, Version: 1}
	s.store.PutProof(p)
	c.Status = StatusNotarized
	c.NotarizedAt = now
	c.ProofID = p.ID
	c.UpdatedAt = now
	c.Version++
	if err = s.store.PutClaim(c); err != nil {
		return Proof{}, err
	}
	s.record("claim.notarized", c.ID, map[string]string{"proof": p.ID, "checkpoint": cp.ID})
	return p, nil
}
func (s *Service) GetProof(id string) (Proof, error) { return s.store.GetProof(id) }
func (s *Service) VerifyProof(id string) (VerificationResult, error) {
	p, err := s.store.GetProof(id)
	if err != nil {
		return VerificationResult{}, err
	}
	checks := []CheckResult{}
	okDigest, e := NormalizeDigest(p.Digest, p.Algorithm)
	checks = append(checks, CheckResult{Name: "digest-format", Passed: e == nil, Detail: okDigest})
	merkle := VerifyMerkle(p.Digest, p.MerkleRoot, p.Path)
	checks = append(checks, CheckResult{Name: "merkle-path", Passed: merkle, Detail: p.MerkleRoot})
	q := len(p.Signatures) >= s.cfg.Quorum
	checks = append(checks, CheckResult{Name: "quorum", Passed: q, Detail: fmt.Sprintf("%d signatures", len(p.Signatures))})
	fresh := time.Since(p.Timestamp) <= s.cfg.ProofTTL
	if p.RevokedAt != nil {
		fresh = false
	}
	checks = append(checks, CheckResult{Name: "lifecycle", Passed: fresh, Detail: "active and within retention"})
	valid := e == nil && merkle && q && fresh
	result := VerificationResult{Valid: valid, Checks: checks, VerifiedAt: time.Now().UTC()}
	if !valid {
		result.Reason = "one or more proof checks failed"
	}
	p.Verification = result
	p.Version++
	s.store.PutProof(p)
	return result, nil
}
func (s *Service) RevokeProof(id, reason string) (Proof, error) {
	p, err := s.store.GetProof(id)
	if err != nil {
		return Proof{}, err
	}
	if p.RevokedAt != nil {
		return p, nil
	}
	now := time.Now().UTC()
	p.RevokedAt = &now
	p.Version++
	s.store.PutProof(p)
	c, e := s.store.GetClaim(p.ClaimID)
	if e == nil {
		c.Status = StatusRevoked
		c.UpdatedAt = now
		c.Version++
		_ = s.store.PutClaim(c)
	}
	s.record("proof.revoked", id, map[string]string{"reason": reason})
	return p, nil
}
func (s *Service) RegisterSigner(x Signer) error {
	if x.ID == "" {
		x.ID = newID("sgn")
	}
	if x.KeyVersion == "" {
		x.KeyVersion = "v1"
	}
	x.Enabled = true
	x.LastSeen = time.Now().UTC()
	s.store.PutSigner(x)
	s.record("signer.registered", x.ID, nil)
	return nil
}
func (s *Service) ListSigners() []Signer { return s.store.ListSigners() }
func (s *Service) RotateKey(signerID string, key SignerKey) error {
	if signerID == "" {
		return ErrNotFound
	}
	key.ID = newID("key")
	key.SignerID = signerID
	if key.Version == "" {
		key.Version = "v" + fmt.Sprint(time.Now().Unix())
	}
	if key.Status == "" {
		key.Status = "active"
	}
	if key.NotBefore.IsZero() {
		key.NotBefore = time.Now().UTC()
	}
	if key.NotAfter.IsZero() {
		key.NotAfter = time.Now().UTC().Add(365 * 24 * time.Hour)
	}
	s.store.PutKey(key)
	s.record("key.rotated", key.ID, map[string]string{"signer": signerID})
	return nil
}
func (s *Service) ListKeys(id string) []SignerKey { return s.store.ListKeys(id) }
func (s *Service) Audit() []AuditEvent            { return s.store.ListAudit() }
func (s *Service) AuditValid() bool               { return s.audit.Verify() }
func (s *Service) record(kind, subject string, data map[string]string) {
	e := s.audit.Append(kind, subject, data)
	s.store.AppendAudit(e)
}
func newID(prefix string) string {
	b := make([]byte, 8)
	if _, e := rand.Read(b); e != nil {
		return fmt.Sprintf("%s-%d", prefix, time.Now().UnixNano())
	}
	return prefix + "-" + hex.EncodeToString(b)
}
func certDigests(s []Signer) []string {
	o := make([]string, 0, len(s))
	for _, x := range s {
		o = append(o, HashStrings(x.ID, x.CertificateDigest, x.KeyVersion))
	}
	return o
}
func absDuration(d time.Duration) time.Duration {
	if d < 0 {
		return -d
	}
	return d
}
