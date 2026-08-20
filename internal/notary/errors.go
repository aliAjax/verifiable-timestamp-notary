package notary

import "errors"

var (
	ErrNotFound           = errors.New("not found")
	ErrInvalidDigest      = errors.New("invalid digest")
	ErrInvalidAlgorithm   = errors.New("unsupported digest algorithm")
	ErrDuplicate          = errors.New("duplicate request")
	ErrInsufficientQuorum = errors.New("insufficient signer quorum")
	ErrRevoked            = errors.New("proof revoked")
	ErrExpired            = errors.New("proof expired")
	ErrClockSkew          = errors.New("signer clock skew")
	ErrConflict           = errors.New("optimistic version conflict")
	ErrInvalidProof       = errors.New("invalid proof")
)
