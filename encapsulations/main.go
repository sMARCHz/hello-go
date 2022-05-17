package main

import (
	"fmt"

	"github.com/sMARCHz/hello-go/access" // import package by using go module(github.com/sMARCHz/hello-go/access) instead of the relative path(./access)
)

func main() {
	// Golang specifies accessability of things in the package(function, variable, etc.) by first letter
	// Uppercase (Pascal case) = public
	fmt.Println(access.PublicValue)

	// Lowercase (Camel case) = private
	// fmt.Println(access.privateValue) -> error

	// Public means other package can access to it
	// Private means other package can't access to it (Unable to use it)
}
