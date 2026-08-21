package merkle

import "fmt"
func (i *Incremental) AddChecked(c Canonicalizer, leaf Leaf) error { if err := c.Validate([]Leaf{leaf}); err != nil { return fmt.Errorf("incremental add: %v", err) }; i.Add(c.Digest(leaf)); return nil }
func AddLeaves(i *Incremental, c Canonicalizer, leaves []Leaf) error { for _, leaf := range leaves { if err := i.AddChecked(c, leaf); err != nil { return err } }; return nil }
