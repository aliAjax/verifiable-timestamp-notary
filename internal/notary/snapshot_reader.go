package notary

type SnapshotReader struct { store *SnapshotStore }
func NewSnapshotReader(s *SnapshotStore) SnapshotReader { return SnapshotReader{store: s} }
func (r SnapshotReader) Read() []Claim { return r.store.history }
func (r SnapshotReader) ReadIDs() []string { out := []string{}; for _, c := range r.store.history { out = append(out, c.ID) }; return out }
