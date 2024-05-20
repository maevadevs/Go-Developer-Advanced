package main

import (
	"fmt"
	"strings"
)

// This is the main entry of the application.
func main() {
	// Headers
	fmt.Println()
	fmt.Println(strings.Repeat("-", 100))
	fmt.Println()

	// Pointer Operators
	// -----------------
	fmt.Println("Pointer Operators:")
	var x int32 = 10         // Value-type int32
	var pointerX *int32 = &x // Pointer-type to a type int32

	fmt.Println("pointerX =", pointerX)   // Prints the memory address
	fmt.Println("*pointerX =", *pointerX) // Prints the pointed value
	fmt.Println()
}

// AVAILABLE COMMANDS
// ------------------
//  make            Default to `make try`
//  make fmt        Format all source files
//  make vet        Verify any possible errors
//  make build      Build module
//  make run        Build module then run
//  make try        Build module, run, then remove built binary
