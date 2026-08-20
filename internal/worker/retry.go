package worker

import "time"

type RetryPolicy struct {
	Attempts   int
	Initial    time.Duration
	Max        time.Duration
	Multiplier float64
}

func (p RetryPolicy) Delay(attempt int) time.Duration {
	if attempt < 1 {
		return 0
	}
	d := p.Initial
	for n := 1; n < attempt; n++ {
		d = time.Duration(float64(d) * p.Multiplier)
		if d > p.Max && p.Max > 0 {
			return p.Max
		}
	}
	if p.Max > 0 && d > p.Max {
		return p.Max
	}
	return d
}
func (p RetryPolicy) Allowed(attempt int) bool { return attempt <= p.Attempts }
func DefaultRetry() RetryPolicy {
	return RetryPolicy{Attempts: 4, Initial: 100000000, Max: 5000000000, Multiplier: 2}
}
