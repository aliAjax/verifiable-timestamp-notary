package config_test

import (
    "testing"
    "example.com/verifiable-timestamp-notary/internal/claim"
    "example.com/verifiable-timestamp-notary/internal/config"
)

func TestDefaultRulesInitializeMap(t *testing.T) { rules := config.NewRuleSet(); rules.AddAllowed("algorithm", "SHA-256"); if !rules.Allows("algorithm", "SHA-256") { t.Fatal("default rule missing") } }
func TestDefaultValidatorRejectsEmptyInput(t *testing.T) { if got := config.ValidateInput(config.BuildValidator(false), config.Input{}); len(got) == 0 { t.Fatal("empty input accepted") } }
func TestMergeRulesInitializesDestination(t *testing.T) { rules := config.RuleSet{}; config.MergeRules(&rules, map[string][]string{"kind": {"legal"}}); if config.RuleCount(rules) != 1 { t.Fatalf("rules=%+v", rules) } }
func TestClaimAdapterDoesNotBypassValidation(t *testing.T) { if got := claim.ValidateWithPolicy(config.BuildValidator(false), claim.Claim{}); len(got) == 0 { t.Fatal("claim accepted without policy") } }
