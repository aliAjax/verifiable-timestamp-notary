package audit

import "errors"
func EncodeWithCleanup(events []Event, close func() error) (data []byte, err error) { defer func() { err = close() }(); for _, e := range events { if e.ID == "bad" { return nil, errors.New("encode failed") } }; return Encode(events) }
func ExportValid(events []Event) bool { return Verify(events) }
