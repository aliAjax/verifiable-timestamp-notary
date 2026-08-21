package merkle

import (
    "errors"
    "testing"
)

func TestNormalizedLeafErrorKeepsSentinel(t *testing.T) { if _, err := ValidateAndNormalize(Canonicalizer{}, []Leaf{{Key: "", Digest: ""}}); !errors.Is(err, ErrLeaf) { t.Fatalf("err=%v", err) } }
func TestProofValidationErrorKeepsSentinel(t *testing.T) { if err := (Proof{}).Validate(); !errors.Is(err, ErrLeaf) { t.Fatalf("err=%v", err) } }
func TestInvalidProofErrorKeepsDuplicateSentinel(t *testing.T) { if err := (Proof{Leaf: "a", Root: "bad"}).Validate(); !errors.Is(err, ErrDuplicate) { t.Fatalf("err=%v", err) } }
func TestIncrementalLeafErrorKeepsSentinel(t *testing.T) { i := NewIncremental(); if err := i.AddChecked(Canonicalizer{}, Leaf{Key: "", Digest: "d"}); !errors.Is(err, ErrLeaf) { t.Fatalf("err=%v", err) } }
func TestMerkleErrorClassIsStable(t *testing.T) { i := NewIncremental(); err := i.AddChecked(Canonicalizer{}, Leaf{Key: "", Digest: "d"}); if ErrorClass(err) != "invalid" { t.Fatalf("class=%s err=%v", ErrorClass(err), err) } }
