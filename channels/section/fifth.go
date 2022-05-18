package section

import (
	"fmt"
	"sync"
)

var wg5 sync.WaitGroup = sync.WaitGroup{}

// In this section, we'll talk about looping channel to receive the data
// We can loop the channel to receive the data but we need to tell Go when to stop
// Because if we don't close the channel, deadlock is going to happen.
// It happens because at first, data is sending and receiving normally. But when we receive all of the data and no data is coming, but we still loop for receiving the data. So, deadlock happens.

// Approach 1: for range and close()
func FifthSection1() {
	ch := make(chan int, 50)
	wg5.Add(2)
	go func(ch <-chan int) {
		for i := range ch {
			fmt.Println(i)
		}
		wg5.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		// Use close() to close the channel
		// Careful when closing the channel because if some goroutine send data to the channel, program is going to panic
		close(ch)
		wg5.Done()
	}(ch)

	wg5.Wait()
}

// Approach 2: infinite loop and break (not closing the channel)
func FifthSection2() {
	ch := make(chan int, 50)
	wg5.Add(2)
	go func(ch <-chan int) {
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
			wg5.Done()
		}
	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		wg5.Done()
	}(ch)

	wg5.Wait()
}
