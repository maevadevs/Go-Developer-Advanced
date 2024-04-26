package main

import (
	"fmt"
	"math"
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
	fmt.Println("Example of Shadowing Package Names")

	pi := math.Pi

	if float64(x) > pi {
		math := "oops!"

		// This is an error: math.Pi is undefined
		// pi2 := math.Pi

		fmt.Println("\tInside the block, math =", math)
	}

	fmt.Println("Outside the block, math.Pi =", math.Pi)
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
