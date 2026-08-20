package quorum

import "time"

type Window struct {
	Start   time.Time
	End     time.Time
	MaxSkew time.Duration
}

func (w Window) Contains(t time.Time) bool {
	return !t.Before(w.Start.Add(-w.MaxSkew)) && !t.After(w.End.Add(w.MaxSkew))
}
func (w Window) Duration() time.Duration { return w.End.Sub(w.Start) }
func NewWindow(center time.Time, span, skew time.Duration) Window {
	return Window{Start: center.Add(-span), End: center.Add(span), MaxSkew: skew}
}
func (w Window) Valid() bool { return !w.Start.IsZero() && !w.End.Before(w.Start) }
