package section

import (
	"fmt"
	"sync"
)

var wg3 sync.WaitGroup = sync.WaitGroup{}

func ThirdSection() {
	// From the second section, we discover the deadlock problem of channel. We can use buffered channels to handle the deadlock easier.
	// Buffered channels can stores many data in the channel so, the deadlock isn't gonna happen until the buffer is full. (Size of buffer or storage defines when creating channels)
	// If we send the data to the full buffered channel, deadlock is going to happen

	// Buffered channels are useful when you know how many goroutines you have launched, want to limit the number of goroutines you will launch, or want to limit the amount of work that is queued up.

	// Initialize buffered channels with buffer size = 1
	// The buffer size is the number of elements that can be sent to the channel without the send blocking.
	// It means we can store 1 data while waiting to send the data
	// For example:
	// c := make(chan int, 1)
	// c <- 1 // doesn't block
	// c <- 2 // blocks until another goroutine receives previous data (line22) from the channel

	ch := make(chan int, 1)
	wg3.Add(2)

	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		i = <-ch
		fmt.Println(i)
		wg3.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		wg3.Done()
	}(ch)

	wg3.Wait()
}
