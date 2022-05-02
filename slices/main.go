package main

import "fmt"

func main() {
	// Slice is a variable-length sequence which stores elements of a similar type.
	// It is similar to array but it is more powerful, flexible, convenient than an array, it is light-weight data structure.
	// A major difference between arrays and slices -> slice is resizable.

	// Zero value = nil

	// Declaration empty slice
	var sl1 []string
	fmt.Println(sl1)

	// Initialization
	var sl2 []string = []string{"one", "two", "three"}
	fmt.Println(sl2)

	// Initialize using make function
	// Length != Capacity
	// Length is the number of elements currently in the slice
	// Capacity is the number of elements the slice can hold before needing to be reallocated (using double capacity method)
	var sl3 []int = make([]int, 5)    // pass type and length
	var sl4 []int = make([]int, 0, 5) // pass type, length and capacity
	printSlice(sl3)
	printSlice(sl4)

	// Add value to slice
	sl4 = append(sl4, 10)
	sl4 = append(sl4, 20)
	sl4 = append(sl4, 30)           // similar to append(sl4, 10, 20, 30)
	sl4 = append([]int{40}, sl4...) // add 40 to front
	fmt.Println(sl4)

	// Remove value from slice
	sl4 = append(sl4[:1], sl4[2:]...)
	fmt.Printf("After removing: %v\n", sl4)

	// Slice is reference type unlike arrays.
	// sl5 and sl6 are pointing to same underlying array. If we change the value in sl6, we get a change in the value in sl5.
	sl5 := []int{1, 2, 3}
	sl6 := sl5
	sl6[1] = 5
	fmt.Println(sl5, sl6)

	// Slicing
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // we can use array or slice for slicing
	b := a[:]                                 // slice of all elements
	c := a[3:]                                // slice from 4th element to end
	d := a[:6]                                // slice first 6 elements
	e := a[3:6]                               // slice the 4st, 5th, 6th elements
	fmt.Println("Slicing")
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	f := [...]string{"one", "two", "three"}
	fmt.Println(f[1:3]) // หลังตัวที่ 1 ถึงตัวที่ 3
}

func printSlice(x []int) {
	fmt.Printf("length=%d capacity=%d %v\n", len(x), cap(x), x)
}
