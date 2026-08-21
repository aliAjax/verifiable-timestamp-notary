package keylifecycle

import "time"
type RotationFlow struct { State string; Current Key; Previous []Key }
func NewRotationFlow(k Key) RotationFlow { return RotationFlow{State: "active", Current: k} }
func (f *RotationFlow) Transition(next Key, now time.Time) bool { if f.State != "active" && f.State != "rotating" { return false }; f.Previous = append(f.Previous, f.Current); next.Activate(now); f.Current = next; f.State = "active"; return f.Ready() }
func (f RotationFlow) Ready() bool { return f.State == "active" && f.Current.Status == "active" }
