package goroutine

import (
	"fmt"
	"sync"
	"time"
)

func warmServiceConnCache() *sync.Pool {
	p := &sync.Pool{
		New: connectToService,
	}
	for i := 0; i < 10; i++ {
		p.Put(p.New())
	}
	return p
}

func connectToService() interface{} {
	time.Sleep(1 * time.Second) // simulate connection creation
	return struct{}{}
}

func howToUsePool() {
	pool := sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance.")
			return struct{}{}
		},
	}

	pool.Get()
	instance := pool.Get()
	pool.Put(instance)
	pool.Get()
}
