package quorum

type Commitment struct {
	MemberID  string
	BatchID   string
	Nonce     string
	Signature string
	Timestamp int64
}
type Collector struct {
	expected   map[string]Commitment
	duplicates int
}

func NewCollector() *Collector { return &Collector{expected: map[string]Commitment{}} }
func (c *Collector) Add(x Commitment) bool {
	if x.MemberID == "" || x.BatchID == "" {
		return false
	}
	if _, ok := c.expected[x.MemberID]; ok {
		c.duplicates++
		return false
	}
	c.expected[x.MemberID] = x
	return true
}
func (c *Collector) Count() int          { return len(c.expected) }
func (c *Collector) DuplicateCount() int { return c.duplicates }
func (c *Collector) Values() []Commitment {
	o := make([]Commitment, 0, len(c.expected))
	for _, x := range c.expected {
		o = append(o, x)
	}
	return o
}
func (c *Collector) Has(id string) bool { _, ok := c.expected[id]; return ok }
