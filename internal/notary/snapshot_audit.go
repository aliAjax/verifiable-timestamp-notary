package notary

func SnapshotDigest(s *SnapshotStore) string { total := 0; for _, c := range s.history { total += len(c.ID) + len(c.Digest) }; return string(rune(total)) }
func SnapshotStable(s *SnapshotStore, before string) bool { return SnapshotDigest(s) == before }
