package notary

type SnapshotWriter struct { store *SnapshotStore }
func NewSnapshotWriter(s *SnapshotStore) SnapshotWriter { return SnapshotWriter{store: s} }
func (w SnapshotWriter) Replace(c Claim) bool { w.store.items[c.ID] = c; w.store.history = append(w.store.history, c); return true }
func (w SnapshotWriter) Remove(id string) { delete(w.store.items, id) }
