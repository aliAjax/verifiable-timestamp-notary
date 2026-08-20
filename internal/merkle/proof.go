package merkle

type Proof struct {
	Leaf     string
	Root     string
	Siblings []Sibling
	Height   int
}
type Sibling struct {
	Value string
	Left  bool
}

func (p Proof) Verify() bool {
	v := p.Leaf
	for _, s := range p.Siblings {
		if s.Left {
			v = combine(s.Value, v)
		} else {
			v = combine(v, s.Value)
		}
	}
	return v == p.Root
}
func (p Proof) Size() int    { return len(p.Siblings) }
func (p Proof) Clone() Proof { o := p; o.Siblings = append([]Sibling(nil), p.Siblings...); return o }
func BuildProof(root, leaf string, siblings []Sibling) Proof {
	return Proof{Leaf: leaf, Root: root, Siblings: append([]Sibling(nil), siblings...), Height: len(siblings)}
}
