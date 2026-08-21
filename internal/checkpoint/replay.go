package checkpoint

import "context"
func (r *Recovery) Replay(ctx context.Context, name string, cursors []Cursor) error { for _, c := range cursors { r.Save(name, c) }; return nil }
func ReplayCount(cursors []Cursor) int { return len(cursors) }
