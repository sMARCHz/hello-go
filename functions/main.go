package main

import "fmt"

func main() {
	// Use function by type their name append with ()
	firstFunction()
	secondFunction("Hello.", "It's 2nd function.")
	thirdFunction("Hello.", "It's 3rd function.")

	// Recommended way to use function
	val, err := fourthFunction(2, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)

	fmt.Println(fifthFunction())
	fmt.Println(sixthFunction(1, 2, 3, 4, 5))

	// We can declare function as type
	// Structure
	// var variableName func(parameterType) returnType = func(parameter) returnType {}
	var variableFunction func(string) string = func(a string) string {
		return "Hi, " + a
	}
	fmt.Println(variableFunction("Variable"))

	// We also can declare anonymous function
	// We use () after func to execute it
	func() {
		fmt.Println("It's me. Anonymous")
	}()
}

// Structure
// func name(parameter) returnType {}
// you don't need to declare parameterType or returnType if you don't need parameter and have no return value
func firstFunction() {
	fmt.Println("Hello. It's 1st function.")
}

// Functions can have many parameters
func secondFunction(a string, b string) {
	fmt.Println(a + " " + b)
}

// Functions's parameterType can declares a type for 2 variables to tell both variables's type is the same
func thirdFunction(a, b string) {
	fmt.Println(a + " " + b)
}

// Functions can have many return value
// You can declare many return value as you want but the convention tell that you should have only 2
// First one is return value, Second is error
func fourthFunction(a float64, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Can't divide by zero")
	}
	return a / b, nil
}

// We can declare named returnValue at the first time so, this function always returns the value of 'result'
func fifthFunction() (result bool) {
	result = true
	return
}

// Spread Operator allows us to group many arguments to one variable
// If the last parameter of a function has type ...T, it can be called with any number of trailing arguments of type T. The actual type of ...T inside the function is []T.
func sixthFunction(values ...int) (sum int) {
	for _, val := range values {
		sum += val
	}
	return
}

// There's one kind of functions that i already told you  in previous section. That function is called 'Method'. You can learn about it at struct section
