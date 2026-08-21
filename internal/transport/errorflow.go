package transport

import (
    "errors"

    "example.com/verifiable-timestamp-notary/internal/api"
    "example.com/verifiable-timestamp-notary/internal/notary"
)
func MissingClaimStatus(id string) int {
    if errors.Is(notary.WrapMissing(id), notary.ErrNotFound) {
        return 404
    }
    return 500
}
func MissingClaimProblem(id string) api.Problem { return api.MissingProblem(notary.WrapMissing(id)) }
