package rate_limit

import (
	"context"
	"golang.org/x/time/rate"
	"time"
)

type APIConnection struct {
	rateLimiter RateLimiter
}

func Open() *APIConnection {
	return &APIConnection{
		rateLimiter: NewMultiLimiter(
			rate.NewLimiter(Per(2, time.Second), 1),
			rate.NewLimiter(Per(10, time.Minute), 10),
		),
	}
}

func (a *APIConnection) ReadFile(ctx context.Context) error {
	if err := a.rateLimiter.Wait(ctx); err != nil {
		return err
	}
	return nil
}

func (a *APIConnection) ResolveAddress(ctx context.Context) error {
	if err := a.rateLimiter.Wait(ctx); err != nil {
		return err
	}
	return nil
}
