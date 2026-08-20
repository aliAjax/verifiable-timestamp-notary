package api

type Cursor struct {
	Offset int
	Limit  int
	Total  int
}

func (c Cursor) Normalize() Cursor {
	if c.Offset < 0 {
		c.Offset = 0
	}
	if c.Limit < 1 {
		c.Limit = 20
	}
	if c.Limit > 100 {
		c.Limit = 100
	}
	return c
}
func (c Cursor) HasNext() bool { return c.Offset+c.Limit < c.Total }
