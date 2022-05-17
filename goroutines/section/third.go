package section

import (
	"fmt"
	"sync"
)

var wg3 sync.WaitGroup = sync.WaitGroup{}
var counter3 int = 0
var m sync.RWMutex = sync.RWMutex{}

// From the second section's problem we find out that the counter isn't correct so, we'll implement Mutual exclusion concept. The convention name for the data structure that provide the concept is mutex.
// Mutex is used to manage how goroutine can access a variable at a time to avoid conflicts

// Goroutines have to wait to get the lock if some goroutines already acquired it

// Lock -> only one go routine read/write at a time by acquiring the lock.
// RLock -> multiple go routine can read(not write) at a time by acquiring the lock.

func ThirdSection() {
	for i := 0; i < 5; i++ {
		wg3.Add(2)
		m.RLock()
		go sayHello3()
		m.Lock()
		go increment3()
	}
	wg3.Wait()

	// In this section, we only introduce mutex to you so, it's not the best way to code.
	// In the above code, We lose the benefits of concurrency. Therefore, we really don't need to use concurrency that much maybe one thread is fine.
}

func sayHello3() {
	fmt.Printf("Hello #%v\n", counter3)
	m.RUnlock()
	wg3.Done()
}

func increment3() {
	counter3++
	m.Unlock()
	wg3.Done()
}

// These function don't work because maybe we execute 'sayHello3' twice so, Go'll print it twice
// func sayHello3() {
// 	m.RLock()
// 	fmt.Printf("Hello #%v\n", counter3)
// 	m.RUnlock()
// 	wg3.Done()
// }

// func increment3() {
// 	m.Lock()
// 	counter3++
// 	m.Unlock()
// 	wg3.Done()
// }
