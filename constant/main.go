package main

import (
	"fmt"
)

var (
	a int = 10
	b     = "hi"
)

const (
	c int = 10
	d     = "hello"
)

const (
	e = iota
	f
	g
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
)

func main() {
	// Constant is fixed value which can't be changed after initialize
	// Constant is only assignable at compile time. It means we can't assign them in runtime and can't use arithmetic statement
	const myConst int = 20 // typed
	const myConst2 = 10    // untyped
	// myConst = 10 -> error
	// const errConstant float64 = math.Sin(1.57) -> error
	fmt.Println(myConst)
	fmt.Println(myConst2)

	// iota is counting value which increase by one
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
}
