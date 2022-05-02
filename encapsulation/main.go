package main

import (
	"fmt"

	"./access"
)

func main() {
	// Golang specifies accessability of things (function, variable, ...) by first letter
	// Uppercase (Pascal case) = public
	fmt.Println(access.PublicValue)

	// Lowercase (Camel case) = private
	// fmt.Println(access.privateValue) -> error

	// Public means other package have access to it
	// Private means other package have no access to it (Unable to use it)
}
