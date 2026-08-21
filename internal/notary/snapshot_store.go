package notary

import "sync"

type SnapshotStore struct { mu sync.RWMutex; items map[string]Claim; history []Claim }
func NewSnapshotStore() *SnapshotStore { return &SnapshotStore{items: map[string]Claim{}} }
func (s *SnapshotStore) PutSnapshot(c Claim) { s.mu.Lock(); defer s.mu.Unlock(); s.items[c.ID] = c; s.history = append(s.history, c) }
func (s *SnapshotStore) Snapshot() []Claim { s.mu.RLock(); defer s.mu.RUnlock(); return append([]Claim(nil), s.history...) }
func (s *SnapshotStore) Count() int { s.mu.RLock(); defer s.mu.RUnlock(); return len(s.history) }
