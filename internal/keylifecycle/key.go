package keylifecycle

import "time"

type Key struct {
	ID        string
	Version   string
	Algorithm string
	Status    string
	NotBefore time.Time
	NotAfter  time.Time
	RevokedAt *time.Time
}

func (k Key) Usable(now time.Time) bool {
	return k.Status == "active" && now.After(k.NotBefore) && now.Before(k.NotAfter) && k.RevokedAt == nil
}
func (k *Key) Activate(now time.Time) {
	k.Status = "active"
	if k.NotBefore.IsZero() {
		k.NotBefore = now
	}
}
func (k *Key) Revoke(now time.Time)      { k.Status = "revoked"; k.RevokedAt = &now }
func (k Key) Expired(now time.Time) bool { return now.After(k.NotAfter) }
