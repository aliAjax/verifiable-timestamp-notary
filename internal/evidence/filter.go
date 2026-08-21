package evidence

func (e EvidenceList) Filter(state string) []Record {
	out := make([]Record, 0, len(e.items))
	for _, x := range e.items { if x.State == state { out = append(out, x) } }
	return out
}
func (e EvidenceList) All() []Record { return e.items }
