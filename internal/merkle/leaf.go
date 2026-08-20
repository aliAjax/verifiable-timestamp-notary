package merkle

import "sort"

type Leaf struct {
	Key      string
	Digest   string
	Index    int
	Metadata map[string]string
}
type Canonicalizer struct {
	Algorithm  string
	SortLeaves bool
}

func (c Canonicalizer) Normalize(input []Leaf) []Leaf {
	out := append([]Leaf(nil), input...)
	for i := range out {
		out[i].Key = normalizeKey(out[i].Key)
		out[i].Digest = normalizeKey(out[i].Digest)
		out[i].Index = i
	}
	if c.SortLeaves {
		sort.SliceStable(out, func(i, j int) bool { return out[i].Key < out[j].Key })
	}
	return out
}
func normalizeKey(v string) string {
	b := []byte(v)
	for i, x := range b {
		if x >= 'A' && x <= 'Z' {
			b[i] += 32
		}
	}
	return string(b)
}
func (c Canonicalizer) Validate(leaves []Leaf) error {
	seen := map[string]bool{}
	for _, l := range leaves {
		if l.Key == "" || l.Digest == "" {
			return ErrLeaf
		}
		if seen[l.Key] {
			return ErrDuplicate
		}
		seen[l.Key] = true
	}
	return nil
}
func (c Canonicalizer) Digest(leaf Leaf) string { return leaf.Key + ":" + leaf.Digest }

var ErrLeaf = errorValue("leaf is incomplete")
var ErrDuplicate = errorValue("duplicate leaf")

type errorValue string

func (e errorValue) Error() string { return string(e) }
