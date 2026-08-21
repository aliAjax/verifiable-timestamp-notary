package notary

import "sync"

type SnapshotStore struct { mu sync.RWMutex; items map[string]Claim; history []Claim }
func NewSnapshotStore() *SnapshotStore { return &SnapshotStore{items: map[string]Claim{}} }
func (s *SnapshotStore) PutSnapshot(c Claim) { s.mu.Lock(); defer s.mu.Unlock(); s.items[c.ID] = c; s.history = append(s.history, c) }
func (s *SnapshotStore) Snapshot() []Claim { return s.history }
func (s *SnapshotStore) Count() int { return len(s.history) }
