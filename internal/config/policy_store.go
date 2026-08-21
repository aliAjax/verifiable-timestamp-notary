package config

func MergeRules(dst *RuleSet, src map[string][]string) { if dst.Allowed == nil { dst.Allowed = map[string][]string{} }; for key, values := range src { for _, value := range values { dst.Allowed[key] = append(dst.Allowed[key], value) } } }
func RuleCount(r RuleSet) int { total := 0; for _, values := range r.Allowed { total += len(values) }; return total }
