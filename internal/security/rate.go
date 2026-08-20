package security

import "time"

type Bucket struct {
	Capacity int
	Tokens   float64
	Rate     float64
	At       time.Time
}

func (b *Bucket) Allow(now time.Time, cost float64) bool {
	if b.At.IsZero() {
		b.At = now
		b.Tokens = float64(b.Capacity)
	}
	elapsed := now.Sub(b.At).Seconds()
	b.Tokens += elapsed * b.Rate
	if b.Tokens > float64(b.Capacity) {
		b.Tokens = float64(b.Capacity)
	}
	b.At = now
	if b.Tokens < cost {
		return false
	}
	b.Tokens -= cost
	return true
}
func NewBucket(capacity int, rate float64) Bucket {
	return Bucket{Capacity: capacity, Tokens: float64(capacity), Rate: rate, At: time.Now()}
}
