package notary

import (
	"encoding/json"
	"sync"
	"time"
)

type AuditChain struct {
	mu     sync.RWMutex
	events []AuditEvent
}

func NewAuditChain() *AuditChain { return &AuditChain{} }
func (a *AuditChain) Append(kind, subject string, data map[string]string) AuditEvent {
	a.mu.Lock()
	defer a.mu.Unlock()
	prev := ""
	if n := len(a.events); n > 0 {
		prev = a.events[n-1].Hash
	}
	e := AuditEvent{ID: HashStrings(kind, subject, time.Now().UTC().String()), Type: kind, Subject: subject, PrevHash: prev, At: time.Now().UTC(), Data: data}
	payload, _ := json.Marshal(e)
	e.Hash = HashBytes(payload, "SHA-256")
	a.events = append(a.events, e)
	return e
}
func (a *AuditChain) List() []AuditEvent {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return append([]AuditEvent(nil), a.events...)
}
func (a *AuditChain) Verify() bool {
	a.mu.RLock()
	defer a.mu.RUnlock()
	prev := ""
	for _, e := range a.events {
		if e.PrevHash != prev {
			return false
		}
		cp := e
		cp.Hash = ""
		payload, _ := json.Marshal(cp)
		if HashBytes(payload, "SHA-256") != e.Hash {
			return false
		}
		prev = e.Hash
	}
	return true
}
