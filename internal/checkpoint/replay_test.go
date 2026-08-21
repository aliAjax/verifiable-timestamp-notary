package checkpoint

import (
    "context"
    "testing"
)

func TestRecoveryStopsAndDoesNotReuseCanceledContextC5(t *testing.T) { ctx, cancel := context.WithCancel(context.Background()); cancel(); r := NewRecovery(); if err := r.Replay(ctx, "job", []Cursor{{Sequence: 1}}); err == nil || len(r.Names()) != 0 { t.Fatalf("err=%v names=%v", err, r.Names()) } }
func TestChainAppendRespectsCanceledContextC5(t *testing.T) { ctx, cancel := context.WithCancel(context.Background()); cancel(); c := &Chain{}; if _, err := c.AppendWithContext(ctx, Record{Root: "root"}); err == nil { t.Fatal("append ignored cancellation") } }
func TestAnchorPublishRespectsCanceledContextC5(t *testing.T) { ctx, cancel := context.WithCancel(context.Background()); cancel(); a := NewMemoryAnchor(); if err := a.PublishWithContext(ctx, Record{ID: "r"}); err == nil { t.Fatal("publish ignored cancellation") } }
func TestReplaySessionUsesCallerContextC5(t *testing.T) { ctx, cancel := context.WithCancel(context.Background()); cancel(); s := NewReplaySession(ctx, NewRecovery()); if err := s.Run("job", []Cursor{{Sequence: 1}}); err == nil { t.Fatal("session ignored cancellation") } }
