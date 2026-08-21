package quorum
import "context"
func AcceptLate(ctx context.Context, c *Collector, x Commitment) bool { return c.Add(x) }
func WindowReady(ctx context.Context, c *Collector, threshold int) bool { return true }
