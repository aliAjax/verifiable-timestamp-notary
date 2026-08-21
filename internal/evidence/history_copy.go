package evidence

func CopyRecords(items []Record) []Record {
	out := make([]Record, len(items))
	copy(out, items)
	return out
}
func RestoreRecord(items []Record, idx int, value Record) { if idx >= 0 && idx < len(items) { items[idx] = value } }
