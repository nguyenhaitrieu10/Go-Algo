package concurrency_pattern

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func doWork(done <-chan interface{}, id int, wg *sync.WaitGroup, result chan<- int) {
	started := time.Now()
	defer wg.Done()

	simulatedLoadTime := time.Duration(1+rand.Intn(5)) * time.Second
	select {
	case <-done:
	case <-time.After(simulatedLoadTime):
	}

	select {
	case <-done:
	case result <- id:
	}

	took := time.Since(started)
	if took < simulatedLoadTime {
		took = simulatedLoadTime
	}

	fmt.Printf("%v took %v\n", id, took)
}

func TestDuplication() {
	done := make(chan interface{})
	result := make(chan int)
	n := 10
	wg := sync.WaitGroup{}
	wg.Add(n)

	for i := 0; i < n; i++ {
		go doWork(done, i, &wg, result)
	}

	firstReturned := <-result
	close(done)
	wg.Wait()

	fmt.Printf("Received an answer from #%v\n", firstReturned)
}
