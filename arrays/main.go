package main

import "fmt"

func main() {
	// Array is a collection of data of the same type.
	// Array has fixed length -> when it's full, you can't add more data to it.
	// Index begins with 0
	var arr [3]string
	arr[0] = "one"
	arr[1] = "two"
	arr[2] = "three"
	fmt.Println(arr)

	// Compact initialization
	arr2 := [3]string{"one", "two", "three"}
	fmt.Println(arr2)

	// Array is a fixed-length sequence so, each elements allocate the same amount of memory (i.e. array of type int allocates 8 bytes per element)
	// !! Unable to do pointer arithmetic
	arr3 := [2]int{1, 2}
	fmt.Println(arr3, &arr3[0], &arr3[1])

	// Initialize with dynamic length
	arr4 := [...]bool{true, true, false, true}
	fmt.Println(arr4)

	// Access each elements via iteration
	// For-each
	for k, v := range arr2 {
		fmt.Printf("Index[%v] = %v\n", k, v)
	}
	// For-i
	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}

	// See more in slice section
}
