package evidence

type Page[T any] struct {
	Items      []T
	Offset     int
	Limit      int
	Total      int
	NextOffset int
}

func Paginate[T any](items []T, offset, limit int) Page[T] {
	if offset < 0 {
		offset = 0
	}
	if limit < 1 {
		limit = 20
	}
	if offset > len(items) {
		offset = len(items)
	}
	end := offset + limit
	if end > len(items) {
		end = len(items)
	}
	next := 0
	if end < len(items) {
		next = end
	}
	return Page[T]{Items: append([]T(nil), items[offset:end]...), Offset: offset, Limit: limit, Total: len(items), NextOffset: next}
}
func Empty[T any]() Page[T] { return Page[T]{Items: []T{}} }
