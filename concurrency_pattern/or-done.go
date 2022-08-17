package concurrency_pattern

func OrDone(done <-chan interface{}, c <-chan interface{}) <-chan interface{} {
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
