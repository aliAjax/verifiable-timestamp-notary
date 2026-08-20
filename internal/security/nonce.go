package security

import "sync"

type NonceSet struct {
	mu     sync.Mutex
	values map[string]struct{}
}

func NewNonceSet() *NonceSet { return &NonceSet{values: map[string]struct{}{}} }
func (n *NonceSet) Use(v string) bool {
	n.mu.Lock()
	defer n.mu.Unlock()
	if v == "" {
		return false
	}
	if _, ok := n.values[v]; ok {
		return false
	}
	n.values[v] = struct{}{}
	return true
}
func (n *NonceSet) Has(v string) bool {
	n.mu.Lock()
	defer n.mu.Unlock()
	_, ok := n.values[v]
	return ok
}
func (n *NonceSet) Size() int { n.mu.Lock(); defer n.mu.Unlock(); return len(n.values) }
