package checkpoint

type Anchor interface {
	Publish(Record) error
	Resolve(string) (Record, error)
}
type MemoryAnchor struct{ values map[string]Record }

func NewMemoryAnchor() *MemoryAnchor           { return &MemoryAnchor{values: map[string]Record{}} }
func (a *MemoryAnchor) Publish(r Record) error { a.values[r.ID] = r; return nil }
func (a *MemoryAnchor) Resolve(id string) (Record, error) {
	r, ok := a.values[id]
	if !ok {
		return Record{}, ErrMissing
	}
	return r, nil
}

var ErrMissing = anchorError("anchor not found")

type anchorError string

func (e anchorError) Error() string { return string(e) }
