package checkpoint

import "context"

func (a *MemoryAnchor) PublishWithContext(ctx context.Context, r Record) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	return a.Publish(r)
}
func (a *MemoryAnchor) PublishMany(ctx context.Context, records []Record) error {
	for _, r := range records {
		if err := a.PublishWithContext(ctx, r); err != nil {
			return err
		}
	}
	return nil
}
