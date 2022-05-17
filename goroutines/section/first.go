package section

import (
	"fmt"
	"time"
)

func FirstSection() {
	// Goroutine is a lightweight thread managed by the Go runtime.
	// Goroutine supports concurrency and parallelism
	// !! We can't guarantee the order of execution of the goroutines.
	sayHello()

	// We can declare the function to run in seperate goroutine(frame) by using go keyword
	go sayWorld()

	// If we pass value 'A' to this go function and we're trying to shadow it to 'B' after invoke the function.
	// The result is that we randomly get 'A' sometimes and 'B' sometimes depend on the time we execute it. If the function executes when msg already changed to 'B', it's going to print 'B'. (same to 'A' if it's not)
	// var msg string = "A"
	// go func() {
	// 	fmt.Println(msg)
	// }()
	// msg = "B"

	// We can prevent the above problem by passing value to the function (value type)
	msg := "Goroutines"
	go func(msg string) {
		fmt.Println(msg)
	}(msg)
	msg = "Try shadowing"

	// Sometimes main goroutine finishes the process before the seperated goroutine can run. So, we can add time.Sleep() to make main goroutine sleep(wait)
	time.Sleep(100 * time.Millisecond)
}

func sayHello() {
	fmt.Println("Hello")
}

func sayWorld() {
	fmt.Println("World")
}
