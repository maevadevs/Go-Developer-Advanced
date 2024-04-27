package main

import (
	"fmt"
	"math"
	"math/rand"
)

// This is the main entry of the application.
func main() {
	fmt.Println()
	fmt.Println("-------------------------------------------------------------")
	fmt.Println()

	// Example of Variable Shadowing
	// -----------------------------
	fmt.Println("Example of Variable Shadowing:")

	x := 10
	fmt.Println("Outside the block, x is:", x)

	if x > 5 {
		fmt.Println("\tInside the block before shadowing, x is:", x)
		// This variable is shadowing the outside variable
		x := 5
		fmt.Println("\tInside the block after shadowing, x is:", x)
	}
	// When the scope of the block ends, the shadow also ends
	fmt.Println("Outside the block again, x is back to:", x)
	fmt.Println()

	// Example of Variable Shadowing With :=
	// -------------------------------------
	fmt.Println("Example of Variable Shadowing With :=")

	if x > 5 {
		x, y := 5, 20
		fmt.Println("\tInside the block, x =", x, "and y =", y)
	}

	fmt.Println("Outside the block, x =", x)
	fmt.Println()

	// Example of Shadowing Package Names
	// ----------------------------------
	fmt.Println("Example of Shadowing Package Names:")

	pi := math.Pi

	if float64(x) > pi {
		math := "oops!"

		// This is an error: math.Pi is undefined
		// pi2 := math.Pi

		fmt.Println("\tInside the block, math =", math)
	}

	fmt.Println("Outside the block, math.Pi =", math.Pi)
	fmt.Println()

	// Example of `if`
	// ---------------
	fmt.Println("Example of `if`:")

	n := rand.Intn(10)
	if n == 0 {
		fmt.Println(n, ": That is too low!")
	} else if n > 5 {
		fmt.Println(n, ": That is too big!")
	} else {
		fmt.Println(n, ": That is a good number!")
	}
	fmt.Println()

	// Example of `if` With Scoped Variables
	// -------------------------------------
	fmt.Println("Example of `if` With Scoped Variables:")

	if m := rand.Intn(10); m == 0 {
		fmt.Println(m, ": That is too low!")
	} else if n > 5 {
		fmt.Println(m, ": That is too big!")
	} else {
		fmt.Println(m, ": That is a good number!")
	}
	// This throws an error: m is undefined here
	// fmt.Println(m)
	fmt.Println()

	// Example of C-Style for-loop
	// ---------------------------
	fmt.Println("Example of C-Style for-loop:")

	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	fmt.Println()

	// Example of Condition-Only for-loop (while-Style)
	// ------------------------------------------------
	fmt.Println("Example of Condition-Only for-loop (while-Style):")

	i := 0
	for i < 10 {
		fmt.Printf("%d ", i)
		i++
	}
	fmt.Println()
	fmt.Println()

	// Simulating a Do-While Loop in Go
	// --------------------------------
	fmt.Println("Simulating a Do-While Loop in Go:")

	j := 0
	for {
		fmt.Println("\tThis runs at least once")
		j++
		if j == 1 {
			break
		}
	}
	fmt.Println()

	// Fizzbuzz: Without Using `continue`
	// ----------------------------------
	fmt.Println("Fizzbuzz: Without Using `continue`:")

	for i := 1; i < 25; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				fmt.Print("FizzBuzz ")
			} else {
				fmt.Print("Fizz ")
			}
		} else if i%5 == 0 {
			fmt.Print("Buzz ")
		} else {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()

	// Fizzbuzz: With Using `continue`
	// ----------------------------------
	fmt.Println("Fizzbuzz: With Using `continue`:")

	for i := 1; i < 25; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("FizzBuzz ")
			continue
		}
		if i%3 == 0 {
			fmt.Print("Fizz ")
			continue
		}
		if i%5 == 0 {
			fmt.Print("Buzz ")
			continue
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	fmt.Println()

	// Using for-range With Slices
	// ---------------------------
	fmt.Println("Using for-range With Slices:")
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Printf("%d/%d ", i, v)
	}
	fmt.Println()
	fmt.Println()

	// Using for-range With Slices Without `i`
	// ---------------------------------------
	fmt.Println("Using for-range With Slices Without `i`:")
	evenVals2 := []int{2, 4, 6, 8, 10, 12}
	for _, v := range evenVals2 {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
	fmt.Println()

	// Using for-range With Slices Without `v`
	// ---------------------------------------
	fmt.Println("Using for-range With Slices Without `v`:")
	uniqueNames := map[string]bool{
		"Fred": true,
		"Raul": true,
		"John": true,
	}
	for k := range uniqueNames {
		fmt.Printf("%s ", k)
	}
	fmt.Println()
	fmt.Println()

	// Map-Iteration-Order Varies
	// --------------------------
	fmt.Println("Map-Iteration-Order Varies:")
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	// Looping 10 times over the map
	for i := 0; i < 10; i++ {
		fmt.Printf("Loop iteration %d: ", i)
		for k, v := range m {
			fmt.Printf("%s:%d ", k, v)
		}
		fmt.Println()
	}
	fmt.Println()

	// Using for-range With Strings
	// ----------------------------
	fmt.Println("Using for-range With Strings:")
	greetings := []string{"Hello!", "Hi π!"}

	// Iterating over the slice
	fmt.Println("index\tbytes\tstring(rn)")
	for _, greeting := range greetings {
		// Iterating over the string
		for i, rn := range greeting {
			fmt.Println(i, "\t", rn, "\t", string(rn))
		}
		fmt.Println()
	}

	// Modifying for-range Loop Variables: No effect on Compound
	// ---------------------------------------------------------
	fmt.Println("Modifying for-range Loop Variables: No effect on Compound:")
	evenInts := []int{2, 4, 6, 8, 10}
	fmt.Println("Before the loop, evenInts =", evenInts)
	for i, v := range evenInts {
		fmt.Print(i, "-", v, " ")
		// Modifying i and v here has no effect on evenInts
		v = 1000
		i = 2000
	}
	fmt.Println()
	fmt.Println("At the end of the loop, evenInts =", evenInts)
	fmt.Println()
}

// FOR WINDOWS:
//  To Run:                 make run-win
//                          go run Blocks-Shadows-Control-Structures\src\main.go
//  To Build:               make build-win
//                          go build -o Blocks-Shadows-Control-Structures\bin\Bscs.exe Blocks-Shadows-Control-Structures\src\main.go
//  To Run after Build:     .\bin\Bscs.exe
//                          .\Blocks-Shadows-Control-Structures\bin\Bscs.exe
//  Try Build + Run:        make try-win
//                          go build -o Blocks-Shadows-Control-Structures\bin\Bscs.exe Blocks-Shadows-Control-Structures\src\main.go && .\Blocks-Shadows-Control-Structures\bin\Bscs.exe && rm .\Blocks-Shadows-Control-Structures\bin\Bscs.exe

// FOR LINUX:
//  To Run:                 make run
//                          go run Blocks-Shadows-Control-Structures/src/main.go
//  To Build:               make build
//                          go build -o Blocks-Shadows-Control-Structures/bin/Bscs Blocks-Shadows-Control-Structures/src/main.go
//  To Run after Build:     ./bin/Bscs
//                          ./Blocks-Shadows-Control-Structures/bin/Bscs
//  Try Build + Run:        make try
//                          go build -o Blocks-Shadows-Control-Structures/bin/Bscs Blocks-Shadows-Control-Structures/src/main.go && ./Blocks-Shadows-Control-Structures/bin/Bscs && rm ./Blocks-Shadows-Control-Structures/bin/Bscs
