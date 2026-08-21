package evidence

func (e EvidenceList) Filter(state string) []Record { out := e.items[:0]; for _, x := range e.items { if x.State == state { out = append(out, x) } }; return out }
func (e EvidenceList) All() []Record { return e.items }
