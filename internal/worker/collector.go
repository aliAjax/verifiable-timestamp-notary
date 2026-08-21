package worker

import "context"
func Collect(ctx context.Context, jobs []Job, fn WorkFunc) ([]BatchResult, error) { results := RunBatch(ctx, jobs, fn); return results, nil }
