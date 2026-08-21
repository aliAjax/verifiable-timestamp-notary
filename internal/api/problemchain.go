package api

import (
    "errors"
    "example.com/verifiable-timestamp-notary/internal/notary"
)
func MissingProblem(err error) Problem { return NewProblem(500, "internal", err.Error()) }
func IsNotFoundError(err error) bool { return errors.Is(err, notary.ErrNotFound) }
func IsClientProblem(p Problem) bool { return p.Status >= 400 && p.Status < 500 }
