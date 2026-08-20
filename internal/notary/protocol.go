package notary

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"
)

type TSARequest struct {
	Version        int    `json:"version"`
	MessageImprint string `json:"message_imprint"`
	HashAlgorithm  string `json:"hash_algorithm"`
	Nonce          string `json:"nonce"`
	CertReq        bool   `json:"cert_req"`
}
type TSAResponse struct {
	Status      string `json:"status"`
	Token       *Proof `json:"token,omitempty"`
	FailureInfo string `json:"failure_info,omitempty"`
}

func DecodeTSARequest(data []byte) (TSARequest, error) {
	var r TSARequest
	if err := json.Unmarshal(data, &r); err != nil {
		return r, fmt.Errorf("tsa request: %w", err)
	}
	if r.Version == 0 {
		r.Version = 1
	}
	return r, nil
}
func (s *Service) HandleTSA(r TSARequest) (TSAResponse, error) {
	c, err := s.CreateClaim(TimestampRequest{Digest: r.MessageImprint, Algorithm: r.HashAlgorithm, Kind: "rfc3161", IdempotencyKey: r.Nonce, PolicyVersion: "tsa-v1"})
	if err != nil {
		return TSAResponse{Status: "rejection", FailureInfo: err.Error()}, err
	}
	p, err := s.Notarize(c.ID)
	if err != nil {
		return TSAResponse{Status: "waiting", FailureInfo: err.Error()}, err
	}
	return TSAResponse{Status: "granted", Token: &p}, nil
}
func EncodeToken(p Proof) []byte {
	b, _ := json.Marshal(p)
	return []byte(base64.RawStdEncoding.EncodeToString(b))
}
func (p Proof) RFC3161Time() string { return p.Timestamp.UTC().Format(time.RFC3339Nano) }
