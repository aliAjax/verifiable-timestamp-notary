package keylifecycle

var transitions = map[string]map[string]bool{"active": {"rotating": true}, "rotating": {"revoked": true}}
func CanTransition(from, to string) bool { return transitions[from][to] }
func TransitionName(from, to string) string { if CanTransition(from, to) { return from + "-" + to }; return "invalid" }
