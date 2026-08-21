package config

type Input struct { Algorithm, Kind string }
type Validator interface { Validate(Input) []string }
type requiredValidator struct{}
func (requiredValidator) Validate(x Input) []string { out := []string{}; if x.Algorithm == "" { out = append(out, "algorithm is required") }; if x.Kind == "" { out = append(out, "kind is required") }; return out }
func BuildValidator(enabled bool) Validator { if !enabled { return nil }; return requiredValidator{} }
func ValidateInput(v Validator, x Input) []string { if v != nil { return v.Validate(x) }; return []string{"validator disabled"} }
