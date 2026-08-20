package store

type Versioned[T any] struct {
	Value   T
	Version int64
}

func (v *Versioned[T]) Update(next T, expected int64) bool {
	if expected != v.Version {
		return false
	}
	v.Value = next
	v.Version++
	return true
}
func (v Versioned[T]) Read() (T, int64) { return v.Value, v.Version }
