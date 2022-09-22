package concurrency_pattern

import (
	"time"
)

func DoWork(done <-chan interface{}, pulseInterval time.Duration, nums ...int) (<-chan int, <-chan interface{}) {
	heartbeat := make(chan interface{})
	intStream := make(chan int)

	go func() {
		defer close(heartbeat)
		defer close(intStream)

		time.Sleep(2 * time.Second) // simulate any long operation

		pulse := time.Tick(pulseInterval)

	numLoop:
		for _, n := range nums {
			for {
				select {
				case <-done:
					return
				case <-pulse:
					select {
					case heartbeat <- struct{}{}:
					default:
					}
				case intStream <- n:
					continue numLoop
				}
			}
		}
	}()

	return intStream, heartbeat
}
