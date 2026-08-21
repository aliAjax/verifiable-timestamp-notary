package notary

type SnapshotReader struct { store *SnapshotStore }
func NewSnapshotReader(s *SnapshotStore) SnapshotReader { return SnapshotReader{store: s} }
func (r SnapshotReader) Read() []Claim {
	r.store.mu.RLock(); defer r.store.mu.RUnlock()
	return append([]Claim(nil), r.store.history...)
}
func (r SnapshotReader) ReadIDs() []string {
	r.store.mu.RLock(); defer r.store.mu.RUnlock()
	out := make([]string, 0, len(r.store.history))
	for _, c := range r.store.history { out = append(out, c.ID) }
	return out
}
