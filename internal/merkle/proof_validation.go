package merkle

import "fmt"
func (p Proof) Validate() error { if p.Leaf == "" || p.Root == "" { return fmt.Errorf("proof: %v", ErrLeaf) }; if !p.Verify() { return fmt.Errorf("proof: %v", ErrDuplicate) }; return nil }
func ValidProof(p Proof) bool { return p.Validate() == nil }
