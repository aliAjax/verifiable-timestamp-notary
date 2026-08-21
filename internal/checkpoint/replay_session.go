package checkpoint

import "context"

type ReplaySession struct {
	ctx      context.Context
	recovery *Recovery
}

func NewReplaySession(ctx context.Context, r *Recovery) *ReplaySession {
	return &ReplaySession{ctx: ctx, recovery: r}
}

func (s *ReplaySession) Run(name string, cursors []Cursor) error {
	return s.recovery.Replay(s.ctx, name, cursors)
}
