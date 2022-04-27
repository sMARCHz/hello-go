package main

import (
	"fmt"
	"strconv"
)

var global string = "global" // Can't initialize with := operator

func main() {
	// Initialize variable
	// 1. full format
	var var1 int = 1
	// 2. compact format
	var2 := "var2"
	fmt.Println(var1)
	fmt.Println(var2)

	// Global variable
	fmt.Println(global)
	global = "shadowing" // shadowing
	fmt.Println(global)

	// Data type
	// This section contains only basic data types. In fact, Go has much more data types than this section (i.e. array, slices, struct, maps, function, interface).

	// 1. Boolean
	// Zero value = false
	var btrue bool = true
	fmt.Printf("%v %T\n", btrue, btrue)
	var bfalse bool
	fmt.Printf("%v %T\n", bfalse, bfalse)

	// 2. Number

	// 2.1. Integer
	// Zero value = 0

	// 2.1.1. Signed integer
	// int8, int16, int32, int64
	// int is either int32 or int64
	var a int = 1
	fmt.Printf("%v %T\n", a, a)

	// 2.1.2. Unsigned integer (only positive)
	// uint8, uint16, uint32, uint64
	// uint is either uint32 or uint64
	var b uint = 1
	fmt.Printf("%v %T\n", b, b)

	// 2.1.3 Byte
	// Synonym of uint8
	var by byte = 1
	fmt.Printf("%v %T\n", by, by)

	// 2.2. Float
	// float32, float64
	// If initialize floating point number with := operator, it automatically is declared as float64
	// Zero value = 0.0
	var c float32 = 1.0
	fmt.Printf("%v %T\n", c, c)

	// 2.3. Complex
	// Real + Imaginary
	// complex64 (float32 + float32), complex128 (float64 + float64)
	// Zero value = 0 + 0i
	var d complex64 = 1 + 2i
	fmt.Printf("%v %T\n", d, d)
	var e complex64 = complex(1, 2)
	fmt.Printf("%v %v\n", real(e), imag(e))

	// 3. Text

	// 3.1. String (utf-8)
	// Each character is byte (uint8)
	// Zero value = ""
	var f string = "1"
	fmt.Printf("%v %T\n", f, f)
	fmt.Printf("%v %T\n", f[0], f[0])
	fmt.Printf("%v %T\n", string(f[0]), string(f[0]))       // convert byte to string
	fmt.Printf("%v %T\n", strconv.Itoa(a), strconv.Itoa(a)) // convert int to string

	// 3.2. Rune (utf-32)
	// Synonym of int32
	// Zero value = 0
	g := 'g'
	fmt.Printf("%v %T\n", g, g)
}
