package worker

import "context"
func Collect(ctx context.Context, jobs []Job, fn WorkFunc) ([]BatchResult, error) { results := RunBatch(ctx, jobs, fn); for _, r := range results { if r.Err != nil { return results, r.Err } }; return results, nil }
