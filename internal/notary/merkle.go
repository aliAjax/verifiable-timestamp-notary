package notary

import "sort"

type MerkleTree struct {
	Leaves []string
	Levels [][]string
}

func NewMerkleTree(leaves []string) *MerkleTree {
	cp := append([]string(nil), leaves...)
	sort.Strings(cp)
	levels := [][]string{cp}
	current := cp
	for len(current) > 1 {
		next := make([]string, 0, (len(current)+1)/2)
		for i := 0; i < len(current); i += 2 {
			right := i + 1
			if right >= len(current) {
				right = i
			}
			next = append(next, HashStrings(current[i], current[right]))
		}
		levels = append(levels, next)
		current = next
	}
	if len(current) == 0 {
		levels = append(levels, []string{""})
	}
	return &MerkleTree{Leaves: cp, Levels: levels}
}
func (t *MerkleTree) Root() string {
	if len(t.Levels) == 0 {
		return ""
	}
	return t.Levels[len(t.Levels)-1][0]
}
func (t *MerkleTree) Proof(leaf string) ([]PathNode, bool) {
	idx := sort.SearchStrings(t.Leaves, leaf)
	if idx >= len(t.Leaves) || t.Leaves[idx] != leaf {
		return nil, false
	}
	path := make([]PathNode, 0, len(t.Levels)-1)
	for level := 0; level < len(t.Levels)-1; level++ {
		row := t.Levels[level]
		sibling := idx ^ 1
		if sibling >= len(row) {
			sibling = idx
		}
		path = append(path, PathNode{Hash: row[sibling], Left: sibling < idx})
		idx /= 2
	}
	return path, true
}
func VerifyMerkle(leaf, root string, path []PathNode) bool {
	current := leaf
	for _, node := range path {
		if node.Left {
			current = HashStrings(node.Hash, current)
		} else {
			current = HashStrings(current, node.Hash)
		}
	}
	return current == root
}
