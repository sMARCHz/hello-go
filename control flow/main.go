package main

import "fmt"

func main() {
	// If statement
	val1 := "abc"
	if val1 == "abc" {
		fmt.Println(val1)
	}

	// If-else statement
	val2 := false
	if val2 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	// If-else if-else statement
	val3Guess := 50
	val3Number := 10
	if val3Guess < val3Number {
		fmt.Println("Too low")
	} else if val3Guess > val3Number {
		fmt.Println("Too high")
	} else if val3Guess == val3Number {
		fmt.Println("Correct")
	} else {
		fmt.Println("No condition matched")
	}

	// If statement with initialization statement
	if val4 := "init"; val4 == "init" {
		fmt.Println(val4)
	}

	// For loop

	// For i
	// Initialization statement, Expression, Post statement
	for i1 := 0; i1 < 5; i1++ {
		fmt.Printf("%v ", i1)
	}
	fmt.Println("full format")

	// Multiple initial value
	for i2, i3 := 0, 10; i2 <= i3; i2, i3 = i2+1, i3/2 { // Can't use ++ or other arithmetic operators in this type of for-loop
		fmt.Printf("[%v %v] ", i2, i3)
	}
	fmt.Println("full format with multiple initial value")

	// Without post statement
	for i4 := 0; i4 < 5; {
		fmt.Printf("%v ", i4)
		i4++
	}
	fmt.Println("without statement")

	// Without initialization statement and post statement (for ; expression ; {} == for expression {})
	// While loop
	i5 := 0
	for i5 < 5 {
		fmt.Printf("%v ", i5)
		i5++
	}
	fmt.Println("without initializing and statement")

	// Infinite loop (while true)
	i6 := 0
	for {
		if i6 == 5 {
			break
		}
		fmt.Printf("%v ", i6)
		i6++
	}
	fmt.Println("while true loop")

	// break -> used for quiting the loop at that point
	// continue -> used for skipping that loop but still in for loop
	// label break -> used for breaking specific loop by label
	i7 := 0
CustomLoop:
	for {
		if i7 == 3 {
			break CustomLoop
		}
		fmt.Printf("%v ", i7)
		i7++
	}
	fmt.Println()

	// For range (string, array, slices, map, ...)
	i8 := "Hello Go"
	for k, v := range i8 {
		fmt.Printf("[%v,%v] ", k, string(v))
	}
	fmt.Println()

	i9 := []int{1, 2, 4, 8, 16, 32, 64} // make([]int, 7)
	for _, v := range i9 {              // !! You can use _ to ignore the value so, Go won't complain you if you don't wanna use that variable.
		fmt.Printf("%v ", v)
	}
	fmt.Println()

	i10 := map[string]int{
		"America":  1,
		"Thailand": 66,
	}
	for k, v := range i10 {
		fmt.Printf("[%v,%v] ", k, v)
	}
	fmt.Println()

	// Switch case
	// Implicit break
	switch1 := "10"
	switch switch1 {
	case "1":
		fmt.Println(switch1)
	case "5":
		fmt.Println(switch1)
	case "10":
		fmt.Println(switch1)
	default:
		fmt.Println("No condition matched")
	}

	// Switch case multiple condition in one case
	switch2 := 10
	switch switch2 {
	case 1, 3, 5, 7, 9:
		fmt.Println("Odd")
	case 2, 4, 6, 8, 10:
		fmt.Println("Even")
	}

	// Switch case with initialization statement
	switch switch3 := 2 + 3; switch3 {
	case 1, 3, 5, 7, 9:
		fmt.Println("Odd")
	case 2, 4, 6, 8, 10:
		fmt.Println("Even")
	}

	// Switch case with expression
	switch4 := 10
	switch {
	case switch4 <= 10:
		fmt.Println("Less than or equal to 10")
	case switch4 <= 20:
		fmt.Println("Less than or equal to 20")
	default:
		fmt.Println("Greater than 20")
	}

	// Fallthrough is going to fall although value doesn't match the condition
	switch5 := 10
	switch {
	case switch5 <= 10:
		fmt.Println("Fall through") // run
		fallthrough
	case switch5 >= 20000000000:
		fmt.Println("Reach this condition") // certainly run without considering the condition
	case switch5 < 1:
		fmt.Println("Doesn't reach here")
	default:
		fmt.Println("No condition matched")
	}

	var switch6 interface{} = "hi"
	switch switch6.(type) {
	case int:
		fmt.Println("This is integer")
	case float64:
		fmt.Println("This is float64")
	case string:
		fmt.Println("This is string")
	default:
		fmt.Println("Not in above condition")
	}
}
