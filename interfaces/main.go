package main

import "fmt"

func main() {
	// Interface is kind of standards which defines the behavior what they can do (only behaviors)
	// Every implementations must have the behaviors that declare in the interface
	// Users doesn't need to know internal algorithms

	// In Golang, it doesn't have explicit implementation like Java by using implements keyword.
	// Declaring methods that match the exact same behavior as interface in the same file is the way Go implements interfaces
	// In summary, there're 2 condition to implement interface
	// 1. Declare all of the methods that the interface need
	// 2. Methods and type must be in the same file as where the interface is declared

	// Here we can change MyWriter to other implementations easily if we want to.
	var w Writer = MyWriter{}
	w.Write([]byte("Hi"))

	// !!! Warning
	// If we implement interfaces with the concrete value (using it directly not an address) but we declare one of the methods with pointer receiver, program's going to panick.
	// It's because when we implement interface with the concrete value, the method need to has a value as the receiver (due to interface context).
	// var wc WriterCloser = BufferedWriter{}

	// Solution is to implement interface with pointer instead of concrete value. If we implement with pointer, we can use both value receiver and pointer receiver without any problems.
	var wc WriterCloser = &BufferedWriter{}
	wc.Write([]byte("Hello"))
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