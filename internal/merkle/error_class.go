package merkle

import "errors"

func ErrorClass(err error) string {
	switch {
	case errors.Is(err, ErrDuplicate):
		return "duplicate"
	case errors.Is(err, ErrLeaf):
		return "invalid"
	default:
		return "unknown"
	}
}
func RetryableLeaf(err error) bool {
	return !errors.Is(err, ErrLeaf) && !errors.Is(err, ErrDuplicate)
}
