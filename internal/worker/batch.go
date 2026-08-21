package worker

import "context"

type WorkFunc func(context.Context, Job) error
type BatchResult struct { Job Job; Err error }
func RunBatch(ctx context.Context, jobs []Job, fn WorkFunc) []BatchResult { results := make(chan BatchResult, len(jobs)); done := make(chan struct{}); go produce(ctx, jobs, fn, results, done); out := []BatchResult{}; for x := range results { out = append(out, x) }; <-done; return out }
func BatchExpected(jobs []Job) int { if len(jobs) == 0 { return 0 }; return len(jobs) - 1 }
