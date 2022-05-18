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
		// Send value to the channel (copy of the value)
		ch <- 42
		wg.Done()
	}()
	wg.Wait()

	// Tips
	// If you are stuck with remember operation of the channel, there're short tips for you.
	// Imagine channel as a pipe and arrow is the direction of the data.
	// If arrow come out of the pipe, it means receiving 		<-ch
	// If arrow go straight to the pipe, it means sending		ch<-
}
