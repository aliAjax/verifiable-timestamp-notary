package notary

import (
    "sync"
    "testing"
)

func TestClaimSnapshotIsolationUnderConcurrentWrites(t *testing.T) {
    s := NewSnapshotStore()
    w := NewSnapshotWriter(s)
    s.PutSnapshot(Claim{ID: "a", Digest: "d1"})
    initial := NewSnapshotReader(s).Read()
    digest := SnapshotDigest(s)
    start := make(chan struct{})
    var wg sync.WaitGroup
    for i := 0; i < 8; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            <-start
            w.Replace(Claim{ID: "x", Digest: "d" + string(rune('a'+i))})
            _ = NewSnapshotReader(s).ReadIDs()
            _ = SnapshotStable(s, digest)
        }(i)
    }
    close(start)
    wg.Wait()
    if len(initial) != 1 || initial[0].ID != "a" { t.Fatalf("history snapshot changed: %#v", initial) }
}
