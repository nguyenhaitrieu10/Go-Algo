package concurrency_pattern

import (
	"math/rand"
	"net/http"
)

type Result struct {
	Err      error
	Response *http.Response
}

func GeneratorRand(done <-chan interface{}) <-chan int {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for {
			randValue := rand.Int()
			select {
			case <-done:
				return
			case intStream <- randValue:
			}
		}
	}()
	return intStream
}

func GeneratorInt(done <-chan interface{}, integers ...int) <-chan int {
	intStream := make(chan int)
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
