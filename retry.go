package github.com/itsabgr/retry

import (
	"context"
	"iter"
	"time"
)

type Steps []time.Duration

func (s Steps) Loop() (bool, time.Duration) {
	if len(s)%2 == 1 {
		return true, s[len(s)-1]
	}
	return false, 0
}

func (s Steps) Retry(ctx context.Context) iter.Seq2[int, time.Duration] {
	return Retry(ctx, s)
}

func Retry(ctx context.Context, s Steps) iter.Seq2[int, time.Duration] {
	return func(yield func(int, time.Duration) bool) {
		var err error
		n := 0
		for i := 0; i < (len(s)/2)*2; i += 2 {
			duration := s[i+1]
			for range s[i] {
				if err = Sleep(ctx, duration); err != nil {
					return
				}
				if !yield(n, duration) {
					return
				}
				n++
			}
		}
		if has, duration := s.Loop(); has {
			for {
				if err = Sleep(ctx, duration); err != nil {
					return
				}
				if !yield(n, duration) {
					return
				}
				n++
			}
		}
	}
}
