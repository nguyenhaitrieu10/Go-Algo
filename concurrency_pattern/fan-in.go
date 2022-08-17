package concurrency_pattern

import "sync"

func FanIn(done <-chan interface{}, channels ...<-chan interface{}) <-chan interface{} {
	multiplexedStream := make(chan interface{})
	var wg sync.WaitGroup

	multiplex := func(c <-chan interface{}) {
		defer wg.Done()
		for v := range c {
			select {
			case <-done:
				return
			case multiplexedStream <- v:
			}
		}
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}

	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()
	return multiplexedStream
}
