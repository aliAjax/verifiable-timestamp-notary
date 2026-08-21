package notary

import "fmt"
func WrapMissing(id string) error { return fmt.Errorf("claim %s: %w", id, ErrNotFound) }
func MissingDetail(id string) string { return WrapMissing(id).Error() }
