package observability

import "errors"
func (r *Registry) Flush(close func() error) (err error) { if close == nil { return errors.New("missing close") }; _ = close(); return nil }
func FlushCount(r *Registry) int { return len(r.Snapshot()) }
