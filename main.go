package main

import (
	"fmt"
	concurrency_pattern "go-algo/concurrency_pattern"
	"sync"
)

func generator(done <-chan interface{}, integers ...int) <-chan interface{} {
	intStream := make(chan interface{})
	go func() {
		defer close(intStream)
		for _, i := range integers {
			select {
			case <-done:
				return
			case intStream <- i:
			}
		}
	}()
	return intStream
}

func add(done <-chan interface{}, intStream <-chan int, value int) <-chan int {
	result := make(chan int)
	go func() {
		defer close(result)
		for i := range intStream {
			select {
			case <-done:
				return
			case result <- i + value:
			}
		}
	}()
	return result
}

func multiply(done <-chan interface{}, intStream <-chan int, value int) <-chan int {
	result := make(chan int)
	go func() {
		defer close(result)
		for i := range intStream {
			select {
			case <-done:
				return
			case result <- i * value:
			}
		}
	}()
	return result
}

func fanIn(done chan interface{}, channels ...<-chan interface{}) <-chan interface{} {
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

func OrDone(done chan interface{}, c <-chan interface{}) <-chan interface{} {
	result := make(chan interface{})

	go func() {
		defer close(result)
		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if !ok {
					return
				}
				select {
				case result <- v:
				case <-done:
				}
			}
		}
	}()
	return result
}

func main() {
	done := make(chan interface{})
	defer close(done)
	intStream := generator(done, 1, 2, 3, 4)

	c1, c2 := concurrency_pattern.Tee(done, intStream)

	for i := range c1 {
		fmt.Println("%v %v", i, <-c2)
	}
}
