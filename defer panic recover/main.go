package main

import "fmt"

func main() {
	// Defer tell Go that the function or method doesn't execute until the function that call the defer function returns (nearby function)
	// Often used for closing database connection and etc.
	fmt.Println("===== Defer =====")
	tryDefer()
	tryDefer2()

	// ==========================================================================================================================

	// Panic is an unexpected condition that arises in program before program gonna terminated. It stops the program at the line where panic happen so, everything which is under that line aren't going to execute.
	// Similar to Exception in other language
	// We can manually create panic with panic() function
	// !! Defer function doesn't stop its execution if the program panic (See more in recover)
	// tryPanic()

	// ==========================================================================================================================

	// Recover is a function used to handle panic.
	// We always see recover() inside of defer functions because defer function ain't stop even though program's panicking so, it gonna execute recover() to handle panic for us.
	// We need to declare defer function before panic happens
	fmt.Println("===== Panic & Recover =====")
	fmt.Println("Start")
	tryRecover()
	fmt.Println("End") // program still running
}

func tryDefer() {
	defer fmt.Println("Closing...")
	fmt.Println("Executing...")
}

func tryDefer2() {
	a := "open"
	defer fmt.Println(a) // In this line, we pass a copy of 'a' to the function not the real one. Therefore, 'a' will print "start" even though its value was changed before function returns.
	a = "close"
}

func tryPanic() (ans int) {
	a, b := 1, 0
	ans = a / b // gonna panic because any number can't divide by zero
	return
}

func tryRecover() {
	fmt.Println("About to panic")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error:", err)
		}
	}()
	panic("Something bad happened")
	fmt.Println("Done panicking") // not be executed because panic stops executing function at line 55 (recover doesn't help because it's in defer function which executes after nearby function returns)
}
