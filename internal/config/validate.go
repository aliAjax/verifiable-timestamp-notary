package config

type Rules struct {
	Required []string
	Allowed  map[string][]string
}

func (r Rules) Validate(values map[string]string) []string {
	errs := []string{}
	for _, k := range r.Required {
		if values[k] == "" {
			errs = append(errs, k+" is required")
		}
	}
	for k, allowed := range r.Allowed {
		v := values[k]
		if v == "" {
			continue
		}
		ok := false
		for _, x := range allowed {
			if x == v {
				ok = true
			}
		}
		if !ok {
			errs = append(errs, k+" has invalid value")
		}
	}
	return errs
}
