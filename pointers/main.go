package main

import "fmt"

func main() {
	// Pointer is a variable that is used to store the memory address of another variable
	// The memory address is always found in hexadecimal format
	// Zero value = nil

	// !! Important operator
	// * -> Use for Declaring variable as a pointer && Dereferencing a pointer to get the value which the pointer points to
	// & -> Use for Getting address of the variable

	var a int = 20

	// Declaration
	var b *int

	// Assign pointer b equal to address of a
	b = &a
	fmt.Println(b)  // print address value
	fmt.Println(*b) // print real value

	// Pointer is a reference type so, if b changes, a will change too.
	fmt.Println(a, *b)
	*b = 60 // dereference pointer to reassign the value where 'b' point to
	fmt.Println(a, *b)

	// We can't do pointer arithmetic in Go
	// c := []int {1, 2, 3}
	// d := &c[0]
	// e := &c[1] - 4
	// fmt.Println(d, e)

	// We can create the pointer of structs by using new() function
	var ms *myStruct
	ms = new(myStruct)
	(*ms).foo = 42         // ms.foo = 42 is fine because compiler understands then dereferences it for us
	fmt.Println((*ms).foo) // fmt.Println(ms.foo)
}

type myStruct struct {
	foo int
}
