package goroutine

import (
	"fmt"
	"sync"
)

func Example1() {
	var count int
	increment := func() {
		count++
	}

	var once sync.Once
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			once.Do(increment)
		}()
	}
	wg.Wait()
	fmt.Printf("Count is %d\n", count)
}

func Example2() {
	var count int
	increment := func() {
		count++
	}
	decrement := func() {
		count--
	}

	var once sync.Once
	once.Do(increment)
	once.Do(decrement)
	fmt.Printf("Count is %d\n", count)
}
