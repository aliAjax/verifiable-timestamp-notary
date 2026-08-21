package keylifecycle

type RotationRegistry struct { current Key; history []Key }
func (r *RotationRegistry) Record(flow RotationFlow) { r.history = append(r.history, flow.Current); r.current = flow.Previous[len(flow.Previous)-1] }
func (r RotationRegistry) Current() Key { return r.current }
func (r RotationRegistry) History() []Key { return append([]Key(nil), r.history...) }
