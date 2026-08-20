package keylifecycle

import "time"

type Rotation struct {
	Current  Key
	Previous []Key
	Grace    time.Duration
}

func (r *Rotation) Rotate(next Key, now time.Time) {
	if r.Current.ID != "" {
		r.Previous = append(r.Previous, r.Current)
	}
	next.Activate(now)
	r.Current = next
}
func (r Rotation) Accept(version string, now time.Time) bool {
	if r.Current.Version == version && r.Current.Usable(now) {
		return true
	}
	for _, k := range r.Previous {
		if k.Version == version && !k.Expired(now.Add(-r.Grace)) && !k.Revoked(now) {
			return true
		}
	}
	return false
}
func (k Key) Revoked(now time.Time) bool { return k.RevokedAt != nil && k.RevokedAt.Before(now) }
func (r *Rotation) Purge(now time.Time) {
	out := r.Previous[:0]
	for _, k := range r.Previous {
		if !k.Expired(now.Add(-r.Grace)) {
			out = append(out, k)
		}
	}
	r.Previous = out
}
