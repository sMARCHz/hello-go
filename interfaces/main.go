package main

import "fmt"

func main() {
	// Interface is kind of standards which defines the behavior what they can do (only behaviors)
	// Every implementations must have the behaviors that declare in the interface
	// Users doesn't need to know internal algorithms

	// In Golang, it doesn't have explicit implementation like Java by using implements keyword.
	// Declaring all methods that match the exact same behavior as interface is the way Go implements interfaces

	// Here we can change MyWriter to other implementations easily if we want to.
	var w Writer = MyWriter{}
	w.Write([]byte("Hi"))

	// !!! Warning
	// If we implement interfaces with the concrete value (initiate variable with a normal value not a pointer) but we declare one of the methods with pointer receiver, program's going to panick.
	// It's because when we implement interface with the concrete value, the method need to has a value as the receiver (interfaces don't count that pointer receiver method as the one of implementation method so, it doesn't match the condition to implement interfaces. !! Only for concrete value).
	// var wc WriterCloser = BufferedWriter{}

	// Solution is to implement interface with pointer instead of concrete value. If we implement with pointer, we can use both value receiver and pointer receiver without any problems.
	var wc WriterCloser = &BufferedWriter{}
	wc.Write([]byte("Hello"))

	// Additional information
	// Empty interfaces (mentioned in the struct section)
	// We can use if-else or switch statement to handle the different types of the empty interface
	var myInterface interface{} = 20
	switch myInterface.(type) {
	case int:
		fmt.Println("It's a integer")
	case string:
		fmt.Println("It's a string")
	}
}

// Declaring interface
type Writer interface {
	Write([]byte) (int, error)
}

// Implementing interface
type MyWriter struct{}

func (m MyWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data)) // parse byte to string
	return n, err
}

// We should have one behavior per interface. If we want more than one, we can composit them
type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type BufferedWriter struct{}

func (bw *BufferedWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

func (bw BufferedWriter) Close() error {
	return nil
}
