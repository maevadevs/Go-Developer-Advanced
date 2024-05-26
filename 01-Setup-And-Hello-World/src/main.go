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
	fmt.Println("A Simple Hello-World Program")
	fmt.Println("----------------------------")
	fmt.Println("Hello world!")
	fmt.Printf("Hello again %s!\n", "world")
	fmt.Println()
	fmt.Println(strings.Repeat("-", 100))
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
