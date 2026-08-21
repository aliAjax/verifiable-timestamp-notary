package worker

import "context"
func produce(ctx context.Context, jobs []Job, fn WorkFunc, results chan<- BatchResult, done chan<- struct{}) { defer close(results); defer close(done); for _, job := range jobs { if err := fn(ctx, job); err != nil { results <- BatchResult{Job: job, Err: err}; return }; results <- BatchResult{Job: job} } }
