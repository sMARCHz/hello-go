package section

import (
	"fmt"
	"sync"
)

var wg2 sync.WaitGroup = sync.WaitGroup{}

func SecondSection() {
	// There are some things that we need to consider when using channel 'Without Buffer'

	// When we're sending data to channel, it's going to block the execution of the goroutine until there's a space available in the channel (This goroutine will wait)
	// If all goroutines are also waiting for each other and none of them is able to proceed, the deadock happens. So, program's going to panic
	// Same as receiving. When we're receiving data from channel, it's going to block the execution of the goroutine until there's a data in the channel

	// Remark :
	// A deadlock happens when a group of goroutines are waiting for each other and none of them is able to proceed.

	ch := make(chan int)
	wg2.Add(1)

	// Only receive once
	go func() {
		i := <-ch
		fmt.Println(i)
		wg2.Done()
	}()

	// Send multiple times
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
