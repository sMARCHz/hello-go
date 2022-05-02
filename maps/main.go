package main

import "fmt"

func main() {
	// Maps is a collection of unordered pairs of key-value.
	// It is widely used because it provides fast lookups and values that can retrieve, update or delete with the help of keys.
	// Similar to hashtable
	// Zero value = nil

	// Declare empty map (key: int, value: string)
	var map1 map[int]string
	fmt.Println(map1)

	// Initialization
	map2 := map[int]string{0: "zero", 1: "one"}
	fmt.Println(map2)

	// Initialize using make
	map3 := make(map[string]string)

	// Add new pair
	map3["Thailand"] = "TH"
	map3["United State"] = "US"
	fmt.Println(map3)

	// Retrieve value
	fmt.Println(map3["Thailand"])
	// We can use optional value to check that we can actually get the data
	_, ok := map3["China"]
	fmt.Println(ok)

	// Update value
	map3["United State"] = "USA"
	fmt.Println(map3)

	// Delete pair
	delete(map3, "Thailand")
	fmt.Println(map3)
}
