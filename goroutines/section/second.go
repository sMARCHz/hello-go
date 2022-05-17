package section

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup = sync.WaitGroup{}
var counter int = 0

// WaitGroup is used to synchronize many goroutines.
// We can tell the main goroutine to wait for all goroutines to finish the executions

func SecondSection() {
	// wg.Add() uses to tell this goroutine(main) how many goroutines to wait
	for i := 0; i < 5; i++ {
		wg.Add(2)
		go sayHello2()
		go increment2()
	}
	// wg.Wait() uses to tell this goroutine(main) to wait for all goroutines to finish before we continue
	wg.Wait()

	// Problem that leads to the next section(3rd)
	// Result:
	// Hello #2
	// Hello #0
	// Hello #4
	// Hello #4
	// Hello #5
	//
	// We can see that the counter is totally mess because of race condition. So, in the next section we'll talk about mutex lock.
}

func sayHello2() {
	// wg.Done() uses to tell Go that this routine finishes execution
	// We use defer because we tell go that we're done after we finished everythings
	defer wg.Done()
	fmt.Printf("Hello #%v\n", counter)
}

func increment2() {
	defer wg.Done()
	counter++
}
