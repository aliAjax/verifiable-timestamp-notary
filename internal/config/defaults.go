package config

type RuleSet struct { Allowed map[string][]string }
func NewRuleSet() RuleSet { return RuleSet{} }
func (r *RuleSet) AddAllowed(key, value string) { if r.Allowed == nil { r.Allowed = map[string][]string{} }; r.Allowed[key] = append(r.Allowed[key], value) }
func (r RuleSet) Allows(key, value string) bool { for _, x := range r.Allowed[key] { if x == value { return true } }; return false }
