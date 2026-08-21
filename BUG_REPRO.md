# Bug reproduction

- Bug: empty configuration paths can write to a nil rule map, call a typed-nil validator, and let an empty claim bypass required validation.
- Trigger: construct the default rule set and disabled validator, merge rules, and validate an empty claim.
- Error: `panic: assignment to entry in nil map` and `panic: value method ... requiredValidator.Validate called using nil *requiredValidator pointer`.
