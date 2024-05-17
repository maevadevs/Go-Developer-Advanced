package main

import "fmt"

// This is the main entry of the application.
func main() {
	fmt.Println()
	fmt.Println("A Simple Hello-World Program")
	fmt.Println("----------------------------")
	fmt.Println("Hello world!")
	fmt.Printf("Hello again %s!\n", "world")
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
