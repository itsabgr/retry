package retry

import (
	"context"
	"runtime"
	"time"
)

func Sleep(ctx context.Context, duration time.Duration) error {
	if duration <= time.Nanosecond {
		runtime.Gosched()
	} else if ctx == nil {
		time.Sleep(duration)
	} else {
		select {
		case <-ctx.Done():
		case <-time.After(duration):
		}
	}
	if ctx == nil {
		return nil
	}
	return ctx.Err()
}
