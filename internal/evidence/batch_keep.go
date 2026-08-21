package evidence

import "time"

func BatchKeep(items []Record, now time.Time, maxAge time.Duration) []Record {
    out := items[:0]
    for _, x := range items { if now.Sub(time.Unix(x.CreatedUnix, 0)) <= maxAge { out = append(out, x) } }
    return out
}
func CountFresh(items []Record, now time.Time, maxAge time.Duration) int { return len(BatchKeep(items, now, maxAge)) }
