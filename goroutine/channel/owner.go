package channel

import "fmt"

func Owner() {
	chanOwner := func() <-chan interface{} {
		c := make(chan interface{}, 5)
		go func() {
			defer close(c)
			for i := 0; i <= 5; i++ {
				c <- i
			}
		}()
		return c
	}

	consumer := func(c <-chan interface{}) {
		for v := range c {
			fmt.Println("Value: ", v)
		}
		fmt.Println("Done")
	}

	c := chanOwner()
	consumer(c)
}
