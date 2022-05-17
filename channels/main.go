package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup = sync.WaitGroup{}

func main() {
	// Declare channels
	var ch chan int

	// Initialize channels
	ch = make(chan int)

	wg.Add(2)
	go func() {
		// Retreive value from the channel
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()
	go func() {
		// Push value to the channel
		ch <- 42
		wg.Done()
	}()
	wg.Wait()
}
