package notary

import (
	"sync"
	"time"
)

type Store interface {
	PutClaim(Claim) error
	GetClaim(string) (Claim, error)
	FindByIdempotency(string) (Claim, error)
	ListClaims(int, int) []Claim
	PutBatch(Batch)
	GetBatch(string) (Batch, error)
	PutCheckpoint(Checkpoint)
	GetCheckpoint(string) (Checkpoint, error)
	PutProof(Proof)
	GetProof(string) (Proof, error)
	PutSigner(Signer)
	ListSigners() []Signer
	PutKey(SignerKey)
	ListKeys(string) []SignerKey
	AppendAudit(AuditEvent)
	ListAudit() []AuditEvent
}
type MemoryStore struct {
	mu          sync.RWMutex
	claims      map[string]Claim
	idem        map[string]string
	batches     map[string]Batch
	checkpoints map[string]Checkpoint
	proofs      map[string]Proof
	signers     map[string]Signer
	keys        map[string]SignerKey
	audit       []AuditEvent
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{claims: map[string]Claim{}, idem: map[string]string{}, batches: map[string]Batch{}, checkpoints: map[string]Checkpoint{}, proofs: map[string]Proof{}, signers: map[string]Signer{}, keys: map[string]SignerKey{}}
}
func (m *MemoryStore) PutClaim(c Claim) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if old, ok := m.claims[c.ID]; ok && old.Version != c.Version-1 {
		return ErrConflict
	}
	m.claims[c.ID] = c
	if c.IdempotencyKey != "" {
		m.idem[c.IdempotencyKey] = c.ID
	}
	return nil
}
func (m *MemoryStore) GetClaim(id string) (Claim, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	c, ok := m.claims[id]
	if !ok {
		return Claim{}, ErrNotFound
	}
	return c, nil
}
func (m *MemoryStore) FindByIdempotency(k string) (Claim, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	id, ok := m.idem[k]
	if !ok {
		return Claim{}, ErrNotFound
	}
	return m.claims[id], nil
}
func (m *MemoryStore) ListClaims(offset, limit int) []Claim {
	m.mu.RLock()
	defer m.mu.RUnlock()
	out := make([]Claim, 0, len(m.claims))
	for _, c := range m.claims {
		out = append(out, c)
	}
	if offset > len(out) {
		offset = len(out)
	}
	end := offset + limit
	if end > len(out) {
		end = len(out)
	}
	return out[offset:end]
}
func (m *MemoryStore) PutBatch(b Batch) { m.mu.Lock(); defer m.mu.Unlock(); m.batches[b.ID] = b }
func (m *MemoryStore) GetBatch(id string) (Batch, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	b, ok := m.batches[id]
	if !ok {
		return Batch{}, ErrNotFound
	}
	return b, nil
}
func (m *MemoryStore) PutCheckpoint(c Checkpoint) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.checkpoints[c.ID] = c
}
func (m *MemoryStore) GetCheckpoint(id string) (Checkpoint, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	c, ok := m.checkpoints[id]
	if !ok {
		return Checkpoint{}, ErrNotFound
	}
	return c, nil
}
func (m *MemoryStore) PutProof(p Proof) { m.mu.Lock(); defer m.mu.Unlock(); m.proofs[p.ID] = p }
func (m *MemoryStore) GetProof(id string) (Proof, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	p, ok := m.proofs[id]
	if !ok {
		return Proof{}, ErrNotFound
	}
	return p, nil
}
func (m *MemoryStore) PutSigner(s Signer) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if s.LastSeen.IsZero() {
		s.LastSeen = time.Now().UTC()
	}
	m.signers[s.ID] = s
}
func (m *MemoryStore) ListSigners() []Signer {
	m.mu.RLock()
	defer m.mu.RUnlock()
	o := make([]Signer, 0, len(m.signers))
	for _, s := range m.signers {
		o = append(o, s)
	}
	return o
}
func (m *MemoryStore) PutKey(k SignerKey) { m.mu.Lock(); defer m.mu.Unlock(); m.keys[k.ID] = k }
func (m *MemoryStore) ListKeys(sid string) []SignerKey {
	m.mu.RLock()
	defer m.mu.RUnlock()
	o := []SignerKey{}
	for _, k := range m.keys {
		if k.SignerID == sid {
			o = append(o, k)
		}
	}
	return o
}
func (m *MemoryStore) AppendAudit(e AuditEvent) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.audit = append(m.audit, e)
}
func (m *MemoryStore) ListAudit() []AuditEvent {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return append([]AuditEvent(nil), m.audit...)
}
