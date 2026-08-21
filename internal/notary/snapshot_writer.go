package notary

type SnapshotWriter struct { store *SnapshotStore }
func NewSnapshotWriter(s *SnapshotStore) SnapshotWriter { return SnapshotWriter{store: s} }
func (w SnapshotWriter) Replace(c Claim) bool { w.store.PutSnapshot(c); return true }
func (w SnapshotWriter) Remove(id string) { w.store.mu.Lock(); defer w.store.mu.Unlock(); delete(w.store.items, id) }
