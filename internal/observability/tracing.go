package observability

import "time"

type Span struct {
	TraceID    string
	Name       string
	Start      time.Time
	Finished   time.Time
	Attributes map[string]string
	Error      string
}

func NewSpan(id, name string) Span {
	return Span{TraceID: id, Name: name, Start: time.Now(), Attributes: map[string]string{}}
}
func (s *Span) End()                   { s.Finished = time.Now() }
func (s *Span) Set(k, v string)        { s.Attributes[k] = v }
func (s Span) Duration() time.Duration { return s.Finished.Sub(s.Start) }
