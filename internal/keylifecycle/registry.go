package keylifecycle

import "time"

type Registry struct{ values map[string]Key }

func NewRegistry() *Registry                  { return &Registry{values: map[string]Key{}} }
func (r *Registry) Put(k Key)                 { r.values[k.ID] = k }
func (r *Registry) Get(id string) (Key, bool) { k, ok := r.values[id]; return k, ok }
func (r *Registry) Active() []Key {
	o := []Key{}
	for _, k := range r.values {
		if k.Status == "active" {
			o = append(o, k)
		}
	}
	return o
}
func (r *Registry) Revoke(id string) {
	k, ok := r.values[id]
	if ok {
		k.Revoke(time.Now())
		r.values[id] = k
	}
}
