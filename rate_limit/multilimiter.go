package rate_limit

import (
	"context"
	"golang.org/x/time/rate"
	"sort"
	"time"
)

type RateLimiter interface {
	Wait(ctx context.Context) error
	Limit() rate.Limit
}

func NewMultiLimiter(limiters ...RateLimiter) *mutiLimiter {
	byLimit := func(i, j int) bool {
		return limiters[i].Limit() < limiters[j].Limit()
	}
	sort.Slice(limiters, byLimit)
	return &mutiLimiter{
		limiters: limiters,
	}
}

type mutiLimiter struct {
	limiters []RateLimiter
}

func (l *mutiLimiter) Wait(ctx context.Context) error {
	for _, l := range l.limiters {
		if err := l.Wait(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (l *mutiLimiter) Limit() rate.Limit {
	return l.limiters[0].Limit()
}

func Per(eventCount int, duration time.Duration) rate.Limit {
	return rate.Every(duration / time.Duration(eventCount))
}
