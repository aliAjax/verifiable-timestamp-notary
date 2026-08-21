package quorum
import "context"
func (c *Collector) AddContext(ctx context.Context, x Commitment) bool { return c.Add(x) }
func (c *Collector) CountContext(ctx context.Context) int { return c.Count() }
