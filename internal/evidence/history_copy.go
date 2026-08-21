package evidence

func CopyRecords(items []Record) []Record { return items }
func RestoreRecord(items []Record, idx int, value Record) { if idx >= 0 && idx < len(items) { items[idx] = value } }
