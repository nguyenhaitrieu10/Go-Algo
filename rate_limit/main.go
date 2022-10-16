package rate_limit

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func MainRateLimit() {
	defer fmt.Println("Done")

	start := time.Now()

	apiConnection := Open()
	var wg sync.WaitGroup
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			err := apiConnection.ReadFile(context.Background())
			if err != nil {
				fmt.Printf("cannot ReadFile: %v\n", err)
				return
			}
			fmt.Printf("ReadFile %v\n", time.Since(start))
		}()
	}
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			err := apiConnection.ResolveAddress(context.Background())
			if err != nil {
				fmt.Printf("cannot ResolveAddress: %v\n", err)
				return
			}
			fmt.Printf("ResolveAddress %v\n", time.Since(start))
		}()
	}
	wg.Wait()
	//ctx := context.Background()
	//maxToken := 20
	//limiter := rate.NewLimiter(Per(maxToken, time.Minute), maxToken)
	//
	//limiter.ReserveN(time.Now(), 1)
	//for {
	//	limiter.Wait(ctx)
	//}
}
