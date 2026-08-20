package protocol

import "encoding/json"

type Envelope struct {
	Version int
	Type    string
	Nonce   string
	Payload json.RawMessage
}

func Encode(e Envelope) ([]byte, error) { return json.Marshal(e) }
func Decode(b []byte) (Envelope, error) { var e Envelope; err := json.Unmarshal(b, &e); return e, err }
func (e Envelope) Valid() bool          { return e.Version > 0 && e.Type != "" && len(e.Payload) > 0 }
