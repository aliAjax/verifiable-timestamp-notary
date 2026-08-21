package merkle

import "fmt"
func (p Proof) Validate() error { if p.Leaf == "" || p.Root == "" { return fmt.Errorf("proof: %w", ErrLeaf) }; if !p.Verify() { return fmt.Errorf("proof: %w", ErrDuplicate) }; return nil }
func ValidProof(p Proof) bool { return p.Validate() == nil }
