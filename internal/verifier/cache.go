package verifier

type Cache struct {
	values       map[string]Result
	hits, misses uint64
}

func NewCache() *Cache { return &Cache{values: map[string]Result{}} }
func (c *Cache) Get(k string) (Result, bool) {
	v, ok := c.values[k]
	if ok {
		c.hits++
	} else {
		c.misses++
	}
	return v, ok
}
func (c *Cache) Put(k string, v Result)  { c.values[k] = v }
func (c *Cache) Stats() (uint64, uint64) { return c.hits, c.misses }
func (c *Cache) Clear()                  { c.values = map[string]Result{} }
