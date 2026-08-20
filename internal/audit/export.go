package audit

import "encoding/json"

type Event struct {
	ID      string
	Type    string
	Subject string
	Hash    string
}

func Encode(events []Event) ([]byte, error) { return json.Marshal(events) }
func Verify(events []Event) bool {
	prev := ""
	for _, e := range events {
		if e.Hash == "" {
			return false
		}
		if prev != "" && e.Subject == "" {
			return false
		}
		prev = e.Hash
	}
	return true
}
