package evidence

import "time"

type Retention struct {
	MaxAge    time.Duration
	Protected map[string]bool
}

func (r Retention) Keep(x Record, now time.Time) bool {
	if r.Protected[x.ID] {
		return true
	}
	return now.Sub(time.Unix(x.CreatedUnix, 0)) <= r.MaxAge
}
func (r *Retention) Protect(id string) {
	if r.Protected == nil {
		r.Protected = map[string]bool{}
	}
	r.Protected[id] = true
}
func (r *Retention) Unprotect(id string) { delete(r.Protected, id) }
