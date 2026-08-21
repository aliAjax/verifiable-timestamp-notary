package protocol
import "context"
func DecodeCommitment(ctx context.Context, b []byte) (Envelope, error) { return Decode(b) }
func DecodeForWindow(ctx context.Context, b []byte) (Envelope, error) { return DecodeCommitment(ctx, b) }
