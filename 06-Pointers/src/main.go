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
	fmt.Println("------------------")

	var x int32 = 10  // Value-type int32
	var y bool = true // Value-type bool

	var ptrX *int32 = &x // Pointer-type to a value of type int32: Referencing
	ptrY := &y           // Pointer-type to a value of type bool

	fmt.Println("ptrX =", ptrX)   // Prints the memory address
	fmt.Println("*ptrX =", *ptrX) // Dereferencing: Prints the pointed value: Same as x
	fmt.Println("x =", x)
	fmt.Println("ptrY =", ptrY)   // Prints the memory address
	fmt.Println("*ptrY =", *ptrY) // Dereferencing: Prints the pointed value: Same as y
	fmt.Println("y =", y)
	fmt.Println()

	// Example of nil Pointer
	// ----------------------
	fmt.Println("Example of nil Pointer:")
	fmt.Println("-----------------------")

	var nilPtr *string // Pointer-type to a value of type string but default to nil

	fmt.Println("nilPtr =", nilPtr)              // Prints nil
	fmt.Println("nilPtr == nil:", nilPtr == nil) // Prints true
	// fmt.Println("*nilPtr =", *nilPtr) // panic: runtime error: invalid memory address or nil pointer dereference
	fmt.Println()

	// Example of Pointer Type
	// -----------------------
	fmt.Println("Example of Pointer Type:")
	fmt.Println("------------------------")

	intVal := 10
	var ptrIntVal *int
	ptrIntVal = &intVal

	fmt.Println("intVal =", intVal)
	fmt.Println("ptrIntVal =", ptrIntVal)
	fmt.Println()

	// Example of Using new()
	// ----------------------
	fmt.Println("Example of Using new():")
	fmt.Println("-----------------------")

	ptrNewVar := new(int)                              // Returns a pointer to 0
	fmt.Println("ptrNewVar == nil:", ptrNewVar == nil) // false
	fmt.Println("*ptrNewVar =", *ptrNewVar)            // 0
	fmt.Println()

	// Generic Pointer Helper For Constants
	// ------------------------------------
	fmt.Println("Generic Pointer Helper For Constants:")
	fmt.Println("-------------------------------------")

	type Person struct {
		FirstName  string
		MiddleName *string
		LastName   string
	}

	p := Person{
		FirstName:  "John",
		MiddleName: makeConstPtr("Edler"), // This works!
		LastName:   "Smith",
	}

	fmt.Println("p =", p)
	fmt.Println()

	// Example of Dereferencing a Pointer to Update Pointed Value
	// ----------------------------------------------------------
	fmt.Println("Example of Dereferencing a Pointer to Update Pointed Value:")
	fmt.Println("-----------------------------------------------------------")

	someInt := 100
	fmt.Println("someInt =", someInt)
	failsToUpdate(&someInt)
	fmt.Println("After failsToUpdate(&someInt), someInt =", someInt)
	succeedToUpdate(&someInt)
	fmt.Println("After succeedToUpdate(&someInt), someInt =", someInt)
	fmt.Println()

	// Footers
	fmt.Println()
	fmt.Println(strings.Repeat("-", 100))
	fmt.Println()
}

// Generic Pointer Helper For Constants
// ------------------------------------

// Generic helper function for getting constant's pointer.
func makeConstPtr[T any](t T) *T {
	return &t
}

// Example of Dereferencing a Pointer to Update Pointed Value
// ----------------------------------------------------------

// A function that does not dereference the parameter fails to update.
func failsToUpdate(ptrX *int) {
	// Attempting to change value at an address by re-assiging a new address to a pointer
	// Address to address
	newValue := 20
	ptrX = &newValue
}

// A function that dereference the parameter succeed to update.
func succeedToUpdate(ptrX *int) {
	// Change a pointed value by dereferencing the pointer, then re-assign a value
	// Value to value
	newValue := 20
	*ptrX = newValue
}

// AVAILABLE COMMANDS
// ------------------
//  make            Default to `make try`
//  make fmt        Format all source files
//  make vet        Verify any possible errors
//  make build      Build module
//  make run        Build module then run
//  make try        Build module, run, then remove built binary
