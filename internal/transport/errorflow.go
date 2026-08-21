package transport

import (
    "example.com/verifiable-timestamp-notary/internal/api"
    "example.com/verifiable-timestamp-notary/internal/notary"
)
func MissingClaimStatus(id string) int { return 500 }
func MissingClaimProblem(id string) api.Problem { return api.MissingProblem(notary.WrapMissing(id)) }
