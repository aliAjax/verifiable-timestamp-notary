package checkpoint

import "context"

func (c *Chain) AppendWithContext(ctx context.Context, r Record) (Record, error) {
	if err := ctx.Err(); err != nil {
		return Record{}, err
	}
	return c.Append(r), nil
}
func (c *Chain) AppendMany(ctx context.Context, records []Record) error {
	for _, r := range records {
		if _, err := c.AppendWithContext(ctx, r); err != nil {
			return err
		}
	}
	return nil
}
