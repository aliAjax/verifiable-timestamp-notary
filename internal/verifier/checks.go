package verifier

type Input struct {
	Digest     string
	Algorithm  string
	MerkleRoot string
	ProofRoot  string
	Signatures int
	Checkpoint bool
	Revoked    bool
	AgeSeconds int64
}
type Result struct {
	Valid  bool
	Checks []Check
}
type Check struct {
	Name   string
	Passed bool
	Detail string
}

func Evaluate(p Policy, in Input) Result {
	checks := []Check{{"algorithm", p.AllowsAlgorithm(in.Algorithm), in.Algorithm}, {"merkle", in.MerkleRoot != "" && in.MerkleRoot == in.ProofRoot, "root comparison"}, {"signatures", p.ValidSignatures(in.Signatures), "threshold"}, {"checkpoint", !p.RequireCheckpoint || in.Checkpoint, "anchor required"}, {"lifecycle", !in.Revoked && in.AgeSeconds <= p.MaxAgeSeconds, "retention"}}
	ok := true
	for _, c := range checks {
		ok = ok && c.Passed
	}
	return Result{Valid: ok, Checks: checks}
}
func Summary(r Result) string {
	if r.Valid {
		return "valid"
	}
	return "invalid"
}
