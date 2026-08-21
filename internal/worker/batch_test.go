package worker

import (
    "context"
    "errors"
    "testing"
    "time"
)

func TestWorkerCollectionClosesAfterFailure(t *testing.T) {
    jobs := []Job{{ID: "ok-1"}, {ID: "bad"}, {ID: "ok-2"}}
    done := make(chan struct{})
    go func() { results, err := Collect(context.Background(), jobs, func(_ context.Context, j Job) error { if j.ID == "bad" { return errors.New("boom") }; return nil }); if err == nil || len(results) != 2 { t.Errorf("results=%d err=%v", len(results), err) }; close(done) }()
    select { case <-done: case <-time.After(time.Second): t.Fatal("worker collection hung") }
}
