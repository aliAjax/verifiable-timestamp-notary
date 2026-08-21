package keylifecycle

import "time"
func Promote(k *Key, now time.Time) { k.Status = "active"; k.NotBefore = now }
func Retire(k *Key, now time.Time) { k.Revoke(now) }
func IsCurrent(k Key, now time.Time) bool { return k.Usable(now) }
