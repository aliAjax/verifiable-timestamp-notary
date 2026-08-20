package quorum

import "time"

type Member struct {
	ID         string
	Weight     int
	Enabled    bool
	LastSeen   time.Time
	KeyVersion string
	Region     string
}
type Set struct {
	Members   map[string]Member
	Threshold int
}

func NewSet(threshold int) *Set {
	if threshold < 1 {
		threshold = 1
	}
	return &Set{Members: map[string]Member{}, Threshold: threshold}
}
func (s *Set) Add(m Member) error {
	if m.ID == "" {
		return ErrMember
	}
	if m.Weight <= 0 {
		m.Weight = 1
	}
	s.Members[m.ID] = m
	return nil
}
func (s *Set) Remove(id string) { delete(s.Members, id) }
func (s *Set) Active(now time.Time, maxSkew time.Duration) []Member {
	o := []Member{}
	for _, m := range s.Members {
		if m.Enabled && !m.LastSeen.IsZero() && abs(now.Sub(m.LastSeen)) <= maxSkew {
			o = append(o, m)
		}
	}
	return o
}
func (s *Set) Quorum(count int) bool { return count >= s.Threshold }
func (s *Set) Weight(members []Member) int {
	n := 0
	for _, m := range members {
		n += m.Weight
	}
	return n
}
func abs(d time.Duration) time.Duration {
	if d < 0 {
		return -d
	}
	return d
}

var ErrMember = textError("member id required")

type textError string

func (e textError) Error() string { return string(e) }
