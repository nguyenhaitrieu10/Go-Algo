package concurrency_pattern

import (
	"fmt"
	"math/rand"
	"time"
)

func DoLongWorkWithHeartBeat(done <-chan interface{}, pulseInterval time.Duration) (<-chan interface{}, <-chan time.Time) {
	heartBeat := make(chan interface{})
	result := make(chan time.Time)
	go func() {
		defer close(heartBeat)
		defer close(result)

		pulse := time.Tick(pulseInterval)
		genWork := time.Tick(pulseInterval * 2)

		sendHeartBeat := func() {
			select {
			case heartBeat <- struct{}{}:
			default:
			}
		}
		sendResult := func(work time.Time) {
			for {
				select {
				case <-done:
					return
				case <-pulse:
					sendHeartBeat()
				case result <- work:
					return
				}
			}
		}

		for {
			select {
			case <-done:
				return
			case <-pulse:
				sendHeartBeat()
			case work := <-genWork:
				sendResult(work)
			}
		}
	}()
	return heartBeat, result
}

func RunExampleLongWork() {
	done := make(chan interface{})
	time.AfterFunc(10*time.Second, func() { defer close(done) })

	const timeout = 2 * time.Second
	heartBeat, result := DoLongWorkWithHeartBeat(done, timeout/2)

	for {
		select {
		case _, ok := <-heartBeat:
			if !ok {
				return
			}
			fmt.Println("pulse...")
		case r, ok := <-result:
			if !ok {
				return
			}
			fmt.Printf("results: %v\n", r.Second())
		case <-time.After(timeout):
			return
		}
	}
}

func HeartBeatTellProgress(done <-chan interface{}, pulseInterval time.Duration) (<-chan interface{}, <-chan int) {
	heartBeat := make(chan interface{})
	result := make(chan int)
	go func() {
		defer close(heartBeat)
		defer close(result)

		for i := 0; i < 10; i++ {
			select {
			case heartBeat <- struct{}{}:
			default:
			}

			select {
			case <-done:
				return
			case result <- rand.Intn(10):
			}
		}
	}()
	return heartBeat, result
}
