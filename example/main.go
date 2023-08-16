package main

import (
	"fmt"

	opt "github.com/countersoda/go-opt"
)

func main() {
	// Create an Option with a value
	someOpt := opt.Some(42)

	// Unwrap the value and print it
	fmt.Println(someOpt.Unwrap()) // Output: 42

	// Create an Option with no value
	noneOpt := opt.None()

	// Check if it's None
	if noneOpt.IsNone() {
		fmt.Println("Option is None")
	}
}
