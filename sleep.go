package retry

import (
	"context"
	"time"
)

func Sleep(ctx context.Context, duration time.Duration) error {
	if ctx == nil {
		time.Sleep(duration)
		return nil
	}
	select {
		case <-ctx.Done():
		case <-time.After(duration):
	}
	return ctx.Err()
}
