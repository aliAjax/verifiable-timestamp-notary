package worker

type Job struct {
	ID       string
	Kind     string
	Payload  []byte
	Attempts int
	Priority int
}
type Queue struct{ values []Job }

func (q *Queue) Push(j Job) { q.values = append(q.values, j) }
func (q *Queue) Pop() (Job, bool) {
	if len(q.values) == 0 {
		return Job{}, false
	}
	best := 0
	for n := 1; n < len(q.values); n++ {
		if q.values[n].Priority > q.values[best].Priority {
			best = n
		}
	}
	j := q.values[best]
	q.values = append(q.values[:best], q.values[best+1:]...)
	return j, true
}
func (q *Queue) Len() int { return len(q.values) }
func (q *Queue) Drain() []Job {
	o := append([]Job(nil), q.values...)
	q.values = q.values[:0]
	return o
}
