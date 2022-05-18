package section

import (
	"fmt"
	"sync"
)

var wg2 sync.WaitGroup = sync.WaitGroup{}

func SecondSection() {
	// There are some things that we need to consider when using channel 'Without Buffer'

	// When we're sending data to channel, it's going to pause the execution of the goroutine until there's a space available in the channel (someone receives the data out of the channel)
	// If we don't receive the data out of the channel, the goroutine is going to be paused forever. So, Go is going to panic from deadlock.

	ch := make(chan int)

	wg2.Add(1)
	go func() {
		i := <-ch
		fmt.Println(i)
		wg2.Done()
	}()
	for j := 0; j < 5; j++ {
		wg2.Add(1)
		go func() {
			fmt.Println("hi")
			ch <- 42
			wg2.Done()
		}()
	}
	wg2.Wait()
	fmt.Println("End")
}
