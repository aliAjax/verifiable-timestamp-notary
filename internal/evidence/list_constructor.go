package evidence

type EvidenceList struct{ items []Record }

func NewEvidenceList(items []Record) EvidenceList { return EvidenceList{items: items} }
