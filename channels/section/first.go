package section

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup = sync.WaitGroup{}

func FirstSection() {
	// Declare channels
	var ch chan int

	// Initialize channels
	ch = make(chan int)

	wg.Add(2)
	go func() {
		// Receive value from the channel
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()
	go func() {
		// Send value to the channel
		ch <- 42
		wg.Done()
	}()
	wg.Wait()
}
