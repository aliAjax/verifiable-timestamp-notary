package verifier

type Policy struct {
	Name              string
	Algorithms        []string
	MinSignatures     int
	MaxAgeSeconds     int64
	RequireCheckpoint bool
}

func (p Policy) AllowsAlgorithm(a string) bool {
	for _, x := range p.Algorithms {
		if x == a {
			return true
		}
	}
	return false
}
func (p Policy) ValidSignatures(n int) bool { return n >= p.MinSignatures }
func (p Policy) Validate() error {
	if p.Name == "" || p.MinSignatures < 1 {
		return ErrPolicy
	}
	if len(p.Algorithms) == 0 {
		return ErrPolicy
	}
	return nil
}

var ErrPolicy = verifyError("invalid policy")

type verifyError string

func (e verifyError) Error() string { return string(e) }
