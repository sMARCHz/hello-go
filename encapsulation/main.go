package main

import (
	"fmt"

	"./access"
)

func main() {
	// Golang specifies access by first letter
	// Uppercase (Pascal case) = public
	fmt.Println(access.PublicValue)

	// Lowercase (Camel case) = private
	// fmt.Println(access.privateValue) -> error
}
