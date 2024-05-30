package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

// This is the main entry of the application.
func main() {
	// Headers
	fmt.Println()
	fmt.Println(strings.Repeat("-", 100))
	fmt.Println()

	// Example of Variable Shadowing
	// -----------------------------
	fmt.Println("Example of Variable Shadowing:")
	fmt.Println("------------------------------")

	x := 100
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
	fmt.Println("-------------------------------------")

	if x > 5 {
		x, y := 5, 20
		fmt.Println("\tInside the block, x =", x, "and y =", y)
	}

	fmt.Println("Outside the block, x =", x)
	fmt.Println()

	// Example of Shadowing Package Names
	// ----------------------------------
	fmt.Println("Example of Shadowing Package Names:")
	fmt.Println("-----------------------------------")

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
	fmt.Println("----------------")

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
	fmt.Println("--------------------------------------")

	if m := rand.Intn(10); m == 0 {
		fmt.Println(m, ": That is too low!")
	} else if m > 5 {
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
	fmt.Println("----------------------------")

	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	fmt.Println()

	// Example of Condition-Only for-loop (while-Style)
	// ------------------------------------------------
	fmt.Println("Example of Condition-Only for-loop (while-Style):")
	fmt.Println("-------------------------------------------------")

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
	fmt.Println("---------------------------------")

	j := 0
	for {
		// do
		fmt.Println("\tThis runs at least once")
		j++
		// while j != 1
		if j == 1 {
			break
		}
	}
	fmt.Println()

	// Fizzbuzz: Without Using `continue`
	// ----------------------------------
	fmt.Println("Fizzbuzz: Without Using `continue`:")
	fmt.Println("-----------------------------------")

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
	fmt.Println()

	// Fizzbuzz: With Using `continue`
	// ----------------------------------
	fmt.Println("Fizzbuzz: With Using `continue`:")
	fmt.Println("--------------------------------")

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
	fmt.Println("----------------------------")

	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Printf("%d:%d ", i, v)
	}
	fmt.Println()
	fmt.Println()

	// Using for-range With Slices Without `i`
	// ---------------------------------------
	fmt.Println("Using for-range With Slices Without `i`:")
	fmt.Println("----------------------------------------")

	evenVals2 := []int{2, 4, 6, 8, 10, 12}
	for _, v := range evenVals2 {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
	fmt.Println()

	// Using for-range With Slices Without `v`
	// ---------------------------------------
	fmt.Println("Using for-range With Slices Without `v`:")
	fmt.Println("----------------------------------------")

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
	fmt.Println("---------------------------")

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
	fmt.Println("-----------------------------")

	greetings := []string{"Hello!", "Hi π!"}

	// Iterating over the slice
	fmt.Println("index\tbytes\tstring(rn)")
	for _, greeting := range greetings {
		// Iterating over the string => runes
		for i, rn := range greeting {
			fmt.Println(i, "\t", rn, "\t", string(rn))
		}
		fmt.Println()
	}

	// Modifying for-range Loop Variables: No effect on Compound
	// ---------------------------------------------------------
	fmt.Println("Modifying for-range Loop Variables: No effect on Compound:")
	fmt.Println("----------------------------------------------------------")

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

	// Using for-range With Labels
	// ---------------------------
	fmt.Println("Using for-range With Labels:")
	fmt.Println("----------------------------")

	greetings2 := []string{"Hello!", "Hi π!"}

	// Iterating over the slice
	fmt.Println("index\tbytes\tstring(rn)")
outerLoop:
	for _, greeting := range greetings2 {
		// Iterating over the string
		for i, rn := range greeting {
			fmt.Println(i, "\t", rn, "\t", string(rn))
			if rn == 'l' {
				// Go to label
				continue outerLoop
			}
		}
		fmt.Println()
	}

	// Using switch in Go
	// ------------------
	fmt.Println("Using switch in Go:")
	fmt.Println("-------------------")

	words := []string{"a", "cow", "smile", "gopher", "octops", "anthropologist"}
	for _, word := range words {
		// switch-scoped variable: size
		switch size := len(word); size {
		case 1, 2, 3, 4:
			// Multiple matches
			fmt.Println("-", word, "is a short word")
		case 5:
			// Case-scoped variable: wordLen
			wordLen := len(word)
			fmt.Println("-", word, "is exactly the right length", wordLen)
		case 6, 7, 8, 9:
			// Empty cases do nothing
			// Not cascading into default
		default:
			fmt.Println("-", word, "is too long")
		}
	}
	fmt.Println()

	// Using switch Within for-loop
	// ----------------------------
	fmt.Println("Using switch Within for-loop:")
	fmt.Println("-----------------------------")

switchLoop:
	for i := range 10 {
		switch i {
		case 0, 2, 4, 6:
			fmt.Println(i, "is even")
		case 3:
			fmt.Println(i, "is divisible by 3 but not by 2")
		case 7:
			fmt.Println(i, "-> Exiting the loop. Good bye!")
			break switchLoop
		default:
			fmt.Println("----- You have reached the default case -----")
		}
	}
	fmt.Println()

	// Example of Blank Switch
	// -----------------------
	fmt.Println("Example of Blank Switch:")
	fmt.Println("------------------------")

	words = []string{"Hi", "Salutation", "Hello"}

	for _, word := range words {
		switch wordLen := len(word); {
		case wordLen < 5:
			fmt.Println("-", word, ": is a short word")
		case wordLen > 10:
			fmt.Println("-", word, ": is too long")
		default:
			fmt.Println("-", word, ": is a good length")
		}
	}
	fmt.Println()

	// Equivalent
	for _, word := range words {
		wordLen := len(word)
		switch {
		case wordLen < 5:
			fmt.Println("-", word, ": is a short word")
		case wordLen > 10:
			fmt.Println("-", word, ": is too long")
		default:
			fmt.Println("-", word, ": is a good length")
		}
	}
	fmt.Println()

	// Fizzbuzz: With Using `switch`
	// -----------------------------
	fmt.Println("Fizzbuzz: With Using `switch`")
	fmt.Println("-----------------------------")

	for i := 1; i < 25; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Print("FizzBuzz ")
		case i%3 == 0:
			fmt.Print("Fizz ")
		case i%5 == 0:
			fmt.Print("Buzz ")
		default:
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
	// Footers
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
