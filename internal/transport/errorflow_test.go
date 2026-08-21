package transport

import (
    "errors"
    "testing"
    "example.com/verifiable-timestamp-notary/internal/notary"
)

func TestMissingErrorChainRemainsClassifiableT3(t *testing.T) { if !errors.Is(notary.WrapMissing("x"), notary.ErrNotFound) { t.Fatal("missing sentinel lost") } }
func TestMissingClaimKeepsNotFoundStatusT3(t *testing.T) { if got := MissingClaimStatus("missing-1"); got != 404 { t.Fatalf("status=%d", got) } }
func TestMissingProblemIsClientErrorT3(t *testing.T) { if got := MissingClaimProblem("missing-1"); got.Status != 404 { t.Fatalf("problem=%+v", got) } }
func TestNotFoundIsNotRetryableT3(t *testing.T) { if RetryableStatus(404) { t.Fatal("not found should not be retried") } }
