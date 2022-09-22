package concurrency_pattern

import (
	"testing"
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

func TestDoWork_GeneratesAllNumbers(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	inSlice := []int{0, 1, 2, 3, 5}
	timeout := 2 * time.Second
	result, heartbeat := DoWork(done, timeout/2, inSlice...)

	var got int
	var i int
	for _, expected := range inSlice {
		got = <-result
		if got != expected {
			t.Errorf("index %v: expected %v, but received %v,", i, expected, got)
		}
		i++
	}

}
