package audit

import "errors"
type BatchResource struct { Closed bool; CloseErr error }
func (r *BatchResource) Close() error { r.Closed = true; return r.CloseErr }
func EncodeBatch(events []Event, r *BatchResource) (data []byte, err error) { for _, e := range events { if e.ID == "bad" { return nil, errors.New("encode failed") } }; _ = r.Close(); return Encode(events) }
func ResourceClosed(r *BatchResource) bool { return r.Closed }
