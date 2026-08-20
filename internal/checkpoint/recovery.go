package checkpoint

type Cursor struct {
	Sequence uint64
	Root     string
	Replay   bool
}
type Recovery struct{ cursors map[string]Cursor }

func NewRecovery() *Recovery                        { return &Recovery{cursors: map[string]Cursor{}} }
func (r *Recovery) Save(name string, c Cursor)      { r.cursors[name] = c }
func (r *Recovery) Load(name string) (Cursor, bool) { c, ok := r.cursors[name]; return c, ok }
func (r *Recovery) Reset(name string)               { delete(r.cursors, name) }
func (r *Recovery) Names() []string {
	o := make([]string, 0, len(r.cursors))
	for n := range r.cursors {
		o = append(o, n)
	}
	return o
}
