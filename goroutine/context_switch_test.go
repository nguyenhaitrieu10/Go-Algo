package goroutine

import (
	"sync"
	"testing"
)

// Run: go test -bench=. -cpu=1 context_switch_test.go
func BenchMarkContextSwitch(b *testing.B) {
	// Here we wait until we’re told to begin.
	// We don’t want the cost of setting up and starting each goroutine to factor into the measurement of context switching.
	begin := make(chan struct{})

	c := make(chan struct{}) // struct{}{} is called an empty struct and takes up no memory
	var wg sync.WaitGroup

	token := struct{}{}
	sender := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i += 1 {
			c <- token
		}
	}
	receiver := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i += 1 {
			<-c
		}
	}

	wg.Add(2)
	go sender()
	go receiver()
	b.StartTimer()
	close(begin) // Here we tell the two goroutines to begin
	wg.Wait()
}
