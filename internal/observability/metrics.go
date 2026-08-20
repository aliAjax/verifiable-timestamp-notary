package observability

type Counter struct {
	Name   string
	Value  uint64
	Labels map[string]string
}
type Registry struct{ counters map[string]*Counter }

func NewRegistry() *Registry { return &Registry{counters: map[string]*Counter{}} }
func (r *Registry) Inc(name string) {
	c := r.counters[name]
	if c == nil {
		c = &Counter{Name: name, Labels: map[string]string{}}
		r.counters[name] = c
	}
	c.Value++
}
func (r *Registry) Value(name string) uint64 {
	if c := r.counters[name]; c != nil {
		return c.Value
	}
	return 0
}
func (r *Registry) Snapshot() []Counter {
	o := []Counter{}
	for _, c := range r.counters {
		cp := *c
		o = append(o, cp)
	}
	return o
}
