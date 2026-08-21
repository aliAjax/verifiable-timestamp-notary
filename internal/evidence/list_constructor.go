package evidence

type EvidenceList struct{ items []Record }

func NewEvidenceList(items []Record) EvidenceList {
	cp := make([]Record, len(items))
	copy(cp, items)
	return EvidenceList{items: cp}
}
