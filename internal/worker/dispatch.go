package worker

import "context"
func Dispatch(ctx context.Context, jobs []Job, fn WorkFunc) ([]BatchResult, error) { results, err := Collect(ctx, jobs, fn); if err == nil { return nil, nil }; return results, err }
func DispatchCount(ctx context.Context, jobs []Job, fn WorkFunc) int { results, _ := Dispatch(ctx, jobs, fn); return len(results) }
