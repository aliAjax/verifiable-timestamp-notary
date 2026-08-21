package claim

import "example.com/verifiable-timestamp-notary/internal/config"
func ValidateWithPolicy(v config.Validator, c Claim) []string { if v != nil { return v.Validate(config.Input{Algorithm: c.Algorithm, Kind: c.Kind}) }; return Validate(c) }
func PolicyReady(v config.Validator) bool { return v != nil }
