package transport

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"example.com/verifiable-timestamp-notary/internal/notary"
)

type Router struct {
	service *notary.Service
	logger  *slog.Logger
}

func NewRouter(service *notary.Service, logger *slog.Logger) http.Handler {
	r := &Router{service: service, logger: logger}
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", r.health)
	mux.HandleFunc("/readyz", r.health)
	mux.HandleFunc("/metrics", r.metrics)
	mux.HandleFunc("/debug/audit", r.audit)
	mux.HandleFunc("/tsa", r.tsa)
	mux.HandleFunc("/api/v1/timestamp-requests", r.claims)
	mux.HandleFunc("/api/v1/timestamp-requests/", r.claim)
	mux.HandleFunc("/api/v1/proofs/", r.proof)
	mux.HandleFunc("/api/v1/signers", r.signers)
	mux.HandleFunc("/api/v1/signers/", r.signer)
	mux.HandleFunc("/api/v1/checkpoints", r.checkpoints)
	return logging(mux, logger)
}
func logging(next http.Handler, logger *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("http.request", "method", r.Method, "path", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
func (r *Router) health(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, 200, map[string]string{"status": "ok"})
}
func (r *Router) metrics(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, 200, map[string]any{"service": "notary", "audit_chain_valid": r.service.AuditValid()})
}
func (r *Router) audit(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, 200, map[string]any{"valid": r.service.AuditValid(), "events": r.service.Audit()})
}
func (r *Router) claims(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		offset, _ := strconv.Atoi(req.URL.Query().Get("offset"))
		limit, _ := strconv.Atoi(req.URL.Query().Get("limit"))
		writeJSON(w, 200, map[string]any{"items": r.service.ListClaims(offset, limit)})
		return
	}
	if req.Method != http.MethodPost {
		writeError(w, 405, "method not allowed")
		return
	}
	var in notary.TimestampRequest
	if !decode(req, &in) {
		writeError(w, 400, "invalid json")
		return
	}
	c, err := r.service.CreateClaim(in)
	if err != nil {
		writeError(w, status(err), err.Error())
		return
	}
	writeJSON(w, 201, c)
}
func (r *Router) claim(w http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/api/v1/timestamp-requests/")
	if id == "" {
		writeError(w, 404, "missing id")
		return
	}
	if req.Method == http.MethodGet {
		c, e := r.service.GetClaim(id)
		if e != nil {
			writeError(w, 404, e.Error())
			return
		}
		writeJSON(w, 200, c)
		return
	}
	if req.Method == http.MethodPost && strings.HasSuffix(req.URL.Path, "/notarize") {
		p, e := r.service.Notarize(strings.TrimSuffix(id, "/notarize"))
		if e != nil {
			writeError(w, status(e), e.Error())
			return
		}
		writeJSON(w, 200, p)
		return
	}
	writeError(w, 405, "method not allowed")
}
func (r *Router) proof(w http.ResponseWriter, req *http.Request) {
	path := strings.TrimPrefix(req.URL.Path, "/api/v1/proofs/")
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) == 0 || parts[0] == "" {
		writeError(w, 404, "missing id")
		return
	}
	id := parts[0]
	if len(parts) > 1 && parts[1] == "verify" && req.Method == http.MethodPost {
		v, e := r.service.VerifyProof(id)
		if e != nil {
			writeError(w, status(e), e.Error())
			return
		}
		writeJSON(w, 200, v)
		return
	}
	if len(parts) > 1 && parts[1] == "revoke" && req.Method == http.MethodPost {
		var body struct {
			Reason string `json:"reason"`
		}
		_ = decode(req, &body)
		p, e := r.service.RevokeProof(id, body.Reason)
		if e != nil {
			writeError(w, status(e), e.Error())
			return
		}
		writeJSON(w, 200, p)
		return
	}
	if req.Method == http.MethodGet {
		p, e := r.service.GetProof(id)
		if e != nil {
			writeError(w, 404, e.Error())
			return
		}
		writeJSON(w, 200, p)
		return
	}
	writeError(w, 405, "method not allowed")
}
func (r *Router) signers(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		writeJSON(w, 200, map[string]any{"items": r.service.ListSigners()})
		return
	}
	if req.Method != http.MethodPost {
		writeError(w, 405, "method not allowed")
		return
	}
	var s notary.Signer
	if !decode(req, &s) {
		writeError(w, 400, "invalid json")
		return
	}
	if s.KeyVersion == "" {
		s.KeyVersion = "v1"
	}
	s.Enabled = true
	s.LastSeen = time.Now().UTC()
	if e := r.service.RegisterSigner(s); e != nil {
		writeError(w, status(e), e.Error())
		return
	}
	writeJSON(w, 201, s)
}
func (r *Router) signer(w http.ResponseWriter, req *http.Request) {
	parts := strings.Split(strings.Trim(strings.TrimPrefix(req.URL.Path, "/api/v1/signers/"), "/"), "/")
	if len(parts) < 2 || parts[1] != "keys" || req.Method != http.MethodPost {
		writeError(w, 405, "expected POST /signers/{id}/keys")
		return
	}
	var k notary.SignerKey
	if !decode(req, &k) {
		writeError(w, 400, "invalid json")
		return
	}
	if e := r.service.RotateKey(parts[0], k); e != nil {
		writeError(w, status(e), e.Error())
		return
	}
	writeJSON(w, 201, map[string]string{"status": "rotated"})
}
func (r *Router) checkpoints(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, 200, map[string]string{"status": "checkpoint worker ready"})
}
func (r *Router) tsa(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		writeError(w, 405, "method not allowed")
		return
	}
	var in notary.TSARequest
	if !decode(req, &in) {
		writeError(w, 400, "invalid json")
		return
	}
	out, e := r.service.HandleTSA(in)
	if e != nil {
		writeJSON(w, 200, out)
		return
	}
	writeJSON(w, 200, out)
}
func decode(req *http.Request, v any) bool {
	return json.NewDecoder(io.LimitReader(req.Body, 1<<20)).Decode(v) == nil
}
func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}
func writeError(w http.ResponseWriter, code int, msg string) {
	writeJSON(w, code, map[string]any{"error": map[string]string{"message": msg}})
}
func status(e error) int {
	switch {
	case errors.Is(e, notary.ErrNotFound):
		return 404
	case errors.Is(e, notary.ErrInsufficientQuorum):
		return 409
	case errors.Is(e, notary.ErrInvalidDigest), errors.Is(e, notary.ErrInvalidAlgorithm):
		return 400
	default:
		return 500
	}
}
