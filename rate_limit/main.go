package rate_limit

import (
	"context"
	"golang.org/x/time/rate"
	"time"
)

func MainRateLimit() {
	ctx := context.Background()
	maxToken := 20
	limiter := rate.NewLimiter(rate.Every(time.Minute/time.Duration(maxToken)), maxToken)

	limiter.ReserveN(time.Now(), 1)
	for {
		limiter.Wait(ctx)
	}
}
