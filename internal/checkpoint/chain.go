package checkpoint

import "time"

type Record struct {
	ID       string
	Sequence uint64
	Root     string
	Previous string
	Hash     string
	At       time.Time
}
type Chain struct{ records []Record }

func (c *Chain) Append(r Record) Record {
	r.Sequence = uint64(len(c.records) + 1)
	if n := len(c.records); n > 0 {
		r.Previous = c.records[n-1].Hash
	}
	r.Hash = r.Root + ":" + r.Previous + ":" + r.At.UTC().String()
	c.records = append(c.records, r)
	return r
}
func (c *Chain) Last() (Record, bool) {
	if len(c.records) == 0 {
		return Record{}, false
	}
	return c.records[len(c.records)-1], true
}
func (c *Chain) List() []Record { return append([]Record(nil), c.records...) }
func (c *Chain) Verify() bool {
	prev := ""
	for n, r := range c.records {
		if n+1 != int(r.Sequence) || r.Previous != prev {
			return false
		}
		prev = r.Hash
	}
	return true
}
func (c *Chain) Length() int { return len(c.records) }
