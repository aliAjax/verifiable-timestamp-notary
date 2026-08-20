package store

type Tx struct{ changes []func() }

func (t *Tx) Add(commit func()) { t.changes = append(t.changes, commit) }
func (t *Tx) Commit() {
	for _, f := range t.changes {
		f()
	}
	t.changes = nil
}
func (t *Tx) Rollback() { t.changes = nil }
func (t *Tx) Len() int  { return len(t.changes) }
