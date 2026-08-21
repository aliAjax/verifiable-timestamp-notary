package keylifecycle

import (
    "testing"
    "time"
)

func baseKeys() (time.Time, Key, Key) { now := time.Now(); old := Key{ID: "old", Version: "v1", Status: "active", NotBefore: now.Add(-time.Minute), NotAfter: now.Add(time.Hour)}; next := Key{ID: "next", Version: "v2", NotAfter: now.Add(time.Hour)}; return now, old, next }
func TestRotationTransitionReachesActive(t *testing.T) { now, old, next := baseKeys(); flow := NewRotationFlow(old); if !flow.Transition(next, now) || !flow.Ready() { t.Fatalf("flow=%+v", flow) } }
func TestRotationRegistryTracksCurrent(t *testing.T) { now, old, next := baseKeys(); flow := NewRotationFlow(old); _ = flow.Transition(next, now); reg := &RotationRegistry{}; reg.Record(flow); if reg.Current().Version != "v2" { t.Fatalf("current=%+v", reg.Current()) } }
func TestRotationTableAllowsPromotion(t *testing.T) { if !CanTransition("rotating", "active") { t.Fatal("transition rejected") } }
func TestPromotedKeyIsUsable(t *testing.T) { now, _, next := baseKeys(); Promote(&next, now); if !IsCurrent(next, now.Add(time.Second)) { t.Fatalf("key=%+v", next) } }
