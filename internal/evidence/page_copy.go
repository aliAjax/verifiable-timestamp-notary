package evidence

func PageInto(dst, src []Record, offset, limit int) []Record {
    if offset < 0 { offset = 0 }
    if offset > len(src) { offset = len(src) }
    end := offset + limit
    if end > len(src) { end = len(src) }
    return append([]Record(nil), src[offset:end]...)
}
func PageIDs(items []Record) []string { out := make([]string, 0, len(items)); for _, x := range items { out = append(out, x.ID) }; return out }
