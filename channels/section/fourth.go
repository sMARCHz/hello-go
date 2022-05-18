package section

import (
	"fmt"
	"sync"
)

var wg4 sync.WaitGroup = sync.WaitGroup{}

func FourthSection() {

	// We can declare the function as read only or write only by how we declare the function's parameter
	// Program'll give you error if we violate the law
	// Read only  : variable <-chan type
	// Write only : variable chan<- type

	ch := make(chan int)
	wg4.Add(2)

	// Read only
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		wg4.Done()
	}(ch)

	// Write only
	go func(ch chan<- int) {
		ch <- 42
		wg4.Done()
	}(ch)

	wg4.Wait()
}
