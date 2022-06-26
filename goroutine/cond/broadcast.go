package cond

import (
	"fmt"
	"sync"
)

func ButtonCLickHandlerDemo() {
	type Button struct {
		Clicked *sync.Cond
	}

	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}
	subscribe := func(c *sync.Cond, f func()) {
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)
		go func() {
			goroutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			f()
		}()
		goroutineRunning.Wait()
	}

	var clickRegister sync.WaitGroup
	clickRegister.Add(3)
	subscribe(button.Clicked, func() {
		defer clickRegister.Done()
		fmt.Println("Maximizing window.")
	})
	subscribe(button.Clicked, func() {
		defer clickRegister.Done()
		fmt.Println("Displaying annoying dialog box!")
	})
	subscribe(button.Clicked, func() {
		defer clickRegister.Done()
		fmt.Println("Mouse clicked.")
	})
	button.Clicked.Broadcast()
	clickRegister.Wait()
}
