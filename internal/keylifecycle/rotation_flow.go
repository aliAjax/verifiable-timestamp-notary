package keylifecycle

import "time"
type RotationFlow struct { State string; Current Key; Previous []Key }
func NewRotationFlow(k Key) RotationFlow { return RotationFlow{State: "active", Current: k} }
func (f *RotationFlow) Transition(next Key, now time.Time) bool { if f.State != "active" && f.State != "rotating" { return false }; f.State = "rotating"; f.Previous = append(f.Previous, f.Current); f.Current = next; return f.State == "active" }
func (f RotationFlow) Ready() bool { return f.State == "active" && f.Current.Status == "active" }
