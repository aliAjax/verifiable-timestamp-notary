package merkle

type Incremental struct {
	levels [][]string
	count  uint64
}

func NewIncremental() *Incremental { return &Incremental{levels: make([][]string, 0, 32)} }
func (i *Incremental) Add(value string) {
	carry := value
	level := 0
	i.count++
	for {
		if level >= len(i.levels) {
			i.levels = append(i.levels, []string{carry})
			return
		}
		if len(i.levels[level]) == 0 {
			i.levels[level] = append(i.levels[level], carry)
			return
		}
		carry = combine(i.levels[level][0], carry)
		i.levels[level] = i.levels[level][:0]
		level++
	}
}
func combine(a, b string) string { return a + "|" + b }
func (i *Incremental) Root() string {
	root := ""
	for n := len(i.levels) - 1; n >= 0; n-- {
		if len(i.levels[n]) > 0 {
			if root == "" {
				root = i.levels[n][0]
			} else {
				root = combine(i.levels[n][0], root)
			}
		}
	}
	return root
}
func (i *Incremental) Count() uint64 { return i.count }
func (i *Incremental) Snapshot() []string {
	o := make([]string, len(i.levels))
	for n, x := range i.levels {
		if len(x) > 0 {
			o[n] = x[0]
		}
	}
	return o
}
