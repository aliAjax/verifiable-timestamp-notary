package notary

import "time"

type ClaimStatus string

const (
	StatusPending     ClaimStatus = "pending"
	StatusAggregating ClaimStatus = "aggregating"
	StatusNotarized   ClaimStatus = "notarized"
	StatusRevoked     ClaimStatus = "revoked"
	StatusExpired     ClaimStatus = "expired"
)

type Claim struct {
	ID             string      `json:"id"`
	Digest         string      `json:"digest"`
	Algorithm      string      `json:"algorithm"`
	Kind           string      `json:"kind"`
	ExternalRef    string      `json:"external_ref"`
	PolicyVersion  string      `json:"policy_version"`
	IdempotencyKey string      `json:"idempotency_key"`
	Status         ClaimStatus `json:"status"`
	CreatedAt      time.Time   `json:"created_at"`
	UpdatedAt      time.Time   `json:"updated_at"`
	NotarizedAt    time.Time   `json:"notarized_at"`
	ProofID        string      `json:"proof_id"`
	Version        int64       `json:"version"`
}
type TimestampRequest struct {
	Digest         string `json:"digest"`
	Algorithm      string `json:"algorithm"`
	Kind           string `json:"kind"`
	ExternalRef    string `json:"external_ref"`
	PolicyVersion  string `json:"policy_version"`
	IdempotencyKey string `json:"idempotency_key"`
}
type Signer struct {
	ID                string        `json:"id"`
	Name              string        `json:"name"`
	Endpoint          string        `json:"endpoint"`
	CertificateDigest string        `json:"certificate_digest"`
	Enabled           bool          `json:"enabled"`
	ClockOffset       time.Duration `json:"clock_offset"`
	KeyVersion        string        `json:"key_version"`
	LastSeen          time.Time     `json:"last_seen"`
}
type SignerKey struct {
	ID, SignerID, Version, Algorithm string
	Status                           string
	NotBefore, NotAfter              time.Time
	RevokedAt                        *time.Time
}
type Commitment struct {
	SignerID, BatchID, Nonce, Signature string
	At                                  time.Time
	ClockOffset                         time.Duration
}
type Batch struct {
	ID                  string
	Leaves              []string
	Root                string
	Status              string
	CreatedAt, ClosedAt time.Time
	CheckpointID        string
}
type Checkpoint struct {
	ID, Root, PreviousID, PreviousHash string
	Sequence                           uint64
	At                                 time.Time
	Signatures                         []Signature
}
type Signature struct {
	SignerID, KeyVersion, Algorithm, Value string
	At                                     time.Time
}
type Proof struct {
	ID                 string             `json:"id"`
	ClaimID            string             `json:"claim_id"`
	Digest             string             `json:"digest"`
	Algorithm          string             `json:"algorithm"`
	Kind               string             `json:"kind"`
	PolicyVersion      string             `json:"policy_version"`
	Timestamp          time.Time          `json:"timestamp"`
	BatchID            string             `json:"batch_id"`
	MerkleRoot         string             `json:"merkle_root"`
	CheckpointID       string             `json:"checkpoint_id"`
	Path               []PathNode         `json:"path"`
	Signatures         []Signature        `json:"signatures"`
	CertificateDigests []string           `json:"certificate_digests"`
	Verification       VerificationResult `json:"verification"`
	RevokedAt          *time.Time         `json:"revoked_at,omitempty"`
	Version            int64              `json:"version"`
}
type PathNode struct {
	Hash string
	Left bool
}
type VerificationResult struct {
	Valid      bool
	Checks     []CheckResult
	VerifiedAt time.Time
	Reason     string
}
type CheckResult struct {
	Name   string
	Passed bool
	Detail string
}
type AuditEvent struct {
	ID, Type, Subject, PrevHash, Hash string
	At                                time.Time
	Data                              map[string]string
}
type Config struct {
	Quorum        int
	MaxClockSkew  time.Duration
	BatchSize     int
	BatchInterval time.Duration
	ProofTTL      time.Duration
}

func DefaultConfig() Config {
	return Config{Quorum: 2, MaxClockSkew: 30 * time.Second, BatchSize: 100, BatchInterval: 30 * time.Second, ProofTTL: 365 * 24 * time.Hour}
}
