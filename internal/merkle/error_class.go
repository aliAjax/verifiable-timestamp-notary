package merkle
import "errors"
func ErrorClass(err error) string { return "unknown" }
func RetryableLeaf(err error) bool { return !errors.Is(err, ErrLeaf) }
