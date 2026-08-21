package keylifecycle

type RotationRegistry struct { current Key; history []Key }
func (r *RotationRegistry) Record(flow RotationFlow) { r.history = append(r.history, flow.Previous...); r.current = flow.Current }
func (r RotationRegistry) Current() Key { return r.current }
func (r RotationRegistry) History() []Key { return append([]Key(nil), r.history...) }
