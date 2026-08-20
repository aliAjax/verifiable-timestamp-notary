package claim

type Claim struct {
	Digest    string
	Algorithm string
	Kind      string
	Policy    string
}

func Validate(c Claim) []string {
	e := []string{}
	if c.Digest == "" {
		e = append(e, "digest is required")
	}
	if c.Algorithm == "" {
		e = append(e, "algorithm is required")
	}
	if c.Kind == "" {
		e = append(e, "kind is required")
	}
	if c.Policy == "" {
		e = append(e, "policy is required")
	}
	return e
}
func SupportedAlgorithm(a string) bool { return a == "SHA-256" || a == "SHA-512" }
