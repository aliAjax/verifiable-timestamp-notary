package evidence

type Record struct {
	ID          string
	ClaimID     string
	Digest      string
	State       string
	CreatedUnix int64
	Tags        map[string]string
}
type Index struct {
	values   map[string]Record
	byDigest map[string]string
}

func NewIndex() *Index                        { return &Index{values: map[string]Record{}, byDigest: map[string]string{}} }
func (i *Index) Put(r Record)                 { i.values[r.ID] = r; i.byDigest[r.Digest] = r.ID }
func (i *Index) Get(id string) (Record, bool) { r, ok := i.values[id]; return r, ok }
func (i *Index) FindDigest(d string) (Record, bool) {
	id, ok := i.byDigest[d]
	if !ok {
		return Record{}, false
	}
	return i.Get(id)
}
func (i *Index) Mark(id, state string) bool {
	r, ok := i.values[id]
	if !ok {
		return false
	}
	r.State = state
	i.values[id] = r
	return true
}
