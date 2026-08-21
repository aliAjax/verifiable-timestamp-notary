package worker

import (
    "context"
    "errors"
    "testing"
    "time"
)

func TestProducerClosesResultChannelOnErrorW2(t *testing.T) {
    results := make(chan BatchResult, 1); done := make(chan struct{})
    results2 := make(chan BatchResult, 1); done2 := make(chan struct{}); start := make(chan struct{})
    go func() { <-start; produce(context.Background(), []Job{{ID: "bad"}}, func(context.Context, Job) error { return errors.New("boom") }, results, done) }()
    go func() { <-start; produce(context.Background(), []Job{{ID: "ok"}}, func(context.Context, Job) error { return nil }, results2, done2) }()
    close(start)
    select { case <-done: case <-time.After(time.Second): t.Fatal("producer did not close done") }
    select { case <-done2: case <-time.After(time.Second): t.Fatal("second producer did not close done") }
    select { case result, ok := <-results: if !ok || result.Err == nil { t.Fatalf("missing error result: %#v open=%v", result, ok) }; if _, ok = <-results; ok { t.Fatal("result channel stayed open") }; case <-time.After(time.Second): t.Fatal("result channel not closed") }
}

func TestCollectorReturnsFailureResultW2(t *testing.T) {
    type result struct { items []BatchResult; err error }
    start := make(chan struct{}); out := make(chan result, 2)
    go func() { <-start; items, err := Collect(context.Background(), []Job{{ID: "bad"}}, func(context.Context, Job) error { return errors.New("boom") }); out <- result{items, err} }()
    go func() { <-start; _, _ = Collect(context.Background(), []Job{{ID: "ok"}}, func(context.Context, Job) error { return nil }); out <- result{} }()
    close(start); got := <-out; if got.err == nil { got = <-out }
    if len(got.items) != 1 || got.err == nil { t.Fatalf("results=%v err=%v", got.items, got.err) }
}

func TestDispatchReportsCompleteResultSetW2(t *testing.T) {
    type result struct { items []BatchResult; err error }
    start := make(chan struct{}); out := make(chan result, 2)
    go func() { <-start; items, err := Dispatch(context.Background(), []Job{{ID: "ok"}}, func(context.Context, Job) error { return nil }); out <- result{items, err} }()
    go func() { <-start; _, _ = Dispatch(context.Background(), []Job{{ID: "bad"}}, func(context.Context, Job) error { return errors.New("boom") }); out <- result{} }()
    close(start); got := <-out; if len(got.items) == 0 { got = <-out }
    if got.err != nil || len(got.items) != 1 || got.items[0].Err != nil { t.Fatalf("results=%v err=%v", got.items, got.err) }
}
func TestBatchExpectedMatchesInputW2(t *testing.T) { jobs := []Job{{ID: "a"}, {ID: "b"}}; start := make(chan struct{}); out := make(chan int, 2); for i := 0; i < 2; i++ { go func() { <-start; out <- BatchExpected(jobs) }() }; close(start); if <-out != len(jobs) || <-out != len(jobs) { t.Fatalf("expected count mismatch") } }
