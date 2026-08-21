package quorum

import (
    "context"
    "runtime"
    "testing"
    "example.com/verifiable-timestamp-notary/internal/protocol"
)

func canceledCollector() (context.Context, *Collector) { ctx, cancel := context.WithCancel(context.Background()); cancel(); return ctx, NewCollector() }
func TestCanceledWindowRejectsLateCommitmentQ0(t *testing.T) { ctx, c := canceledCollector(); if AcceptLate(ctx, c, Commitment{MemberID: "b", BatchID: "b"}) { t.Fatal("late commitment accepted") } }
func TestCanceledCollectorDoesNotReportCountQ0(t *testing.T) { ctx, c := canceledCollector(); c.Add(Commitment{MemberID: "a", BatchID: "b"}); if got := c.CountContext(ctx); got != 0 { t.Fatalf("count=%d", got) } }
func TestCanceledWindowIsNotReadyQ0(t *testing.T) { ctx, c := canceledCollector(); if WindowReady(ctx, c, 1) { t.Fatal("canceled window ready") } }
func TestProtocolDecodeStopsOnCancellationQ0(t *testing.T) { ctx, _ := canceledCollector(); raw, _ := protocol.Encode(protocol.Envelope{Version: 1, Type: "commit", Payload: []byte("\"x\"")}); if _, err := protocol.DecodeCommitment(ctx, raw); err == nil { t.Fatal("decode ignored canceled context") } }
func TestCollectionWindowClosesOnCancellationQ0(t *testing.T) { ctx, cancel := context.WithCancel(context.Background()); c := NewCollector(); w := NewCollectionWindow(ctx, c); cancel(); <-ctx.Done(); for i := 0; i < 1000 && !w.closed; i++ { runtime.Gosched() }; if w.Add(Commitment{MemberID: "a", BatchID: "b"}) { t.Fatal("closed window accepted commitment") } }

