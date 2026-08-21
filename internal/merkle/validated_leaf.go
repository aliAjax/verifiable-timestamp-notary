package merkle

import "fmt"
func ValidateAndNormalize(c Canonicalizer, leaves []Leaf) ([]Leaf, error) { normalized := c.Normalize(leaves); if err := c.Validate(normalized); err != nil { return nil, fmt.Errorf("leaf validation: %v", err) }; return normalized, nil }
func LeafCount(leaves []Leaf) int { return len(leaves) }
