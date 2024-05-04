package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// This is the main entry of the application.
func main() {
	fmt.Println()

	// Example of calling a Function
	// -----------------------------
	fmt.Println("Example of calling a Function:")
	result := div(100, 20)
	fmt.Println("div(100, 20) =", result)
	fmt.Println()

	// Example of a Function With Optional and Named Parameters
	// --------------------------------------------------------
	fmt.Println("Example of a Function With Optional and Named Parameters:")

	// Call MyFunc() one way
	MyFunc(FuncParams{
		LastName: "Smith",
		Age:      50,
	})

	// Call MyFunc() another way
	MyFunc(FuncParams{
		FirstName: "Mary",
		LastName:  "Smith",
	})
	fmt.Println()

	// Example of a Variadic Function
	// ------------------------------
	fmt.Println("Example of a Variadic Function:")
	x := addNums(1, 2, 3, 4, 5, 6, 7, 8, 9)
	y := addNums(21, 43, 65)
	z := addNums()
	nums := []int{10, 20, 30, 40, 50}

	fmt.Println("addNums(1, 2, 3, 4, 5, 6, 7, 8, 9) =", x)
	fmt.Println("addNums(21, 43, 65) =", y)
	fmt.Println("addNums() =", z)
	fmt.Println("addNums(nums...) =", addNums(nums...))
	fmt.Println()

	// Example of Function With Multiple Return Values
	// -----------------------------------------------
	fmt.Println("Example of Function With Multiple Return Values:")
	resDiv, resMod, err := divmod(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("divmod(5, 2) =>", "ResDiv =", resDiv, "ResMod =", resMod)
	fmt.Println()

	// Example of Function With Named Return Values
	// --------------------------------------------
	fmt.Println("Example of Function With Named Return Values:")
	resX, modY, errZ := divmodNamed(5, 2)
	// Error-handling
	if errZ != nil {
		fmt.Println(errZ)
		os.Exit(1)
	}
	fmt.Println("divmodNamed(5, 2) =>", "resX =", resX, "modY =", modY)
	fmt.Println()

	// Example of Declaring a Function Variable
	// ----------------------------------------
	fmt.Println("Example of Declaring a Function Variable:")
	// A function variable
	var myFuncVar func(string) int

	// Using the first case
	myFuncVar = f1
	res := myFuncVar("Hello")
	fmt.Println("myFuncVar(\"Hello\") using f1:", res)

	// Using the second case
	myFuncVar = f2
	res = myFuncVar("Hello")
	fmt.Println("myFuncVar(\"Hello\") using f2:", res)
	fmt.Println()

	// Example of a Simple Calculator With Functions
	// ---------------------------------------------
	fmt.Println("Example of a Simple Calculator With Functions:")
	opMap := map[string]func(int, int) int{
		"+": add,
		"-": sub,
		"*": mul,
		"/": divs,
	}
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
	}
	for _, expr := range expressions {
		if len(expr) != 3 {
			fmt.Println("Invalid expression:", expr)
			continue
		}
		p1, err := strconv.Atoi(expr[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := expr[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("Unsupported Operator:", op)
			continue
		}
		p2, err := strconv.Atoi(expr[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result := opFunc(p1, p2)
		fmt.Println(expr, "=", result)
	}
	fmt.Println()
}

// Example of a Function
// ---------------------
func div(num int, den int) int {
	if den == 0 {
		return 0
	}
	return num / den
}

// Example of a Function With Optional and Named Parameters
// --------------------------------------------------------
type FuncParams struct {
	FirstName string
	LastName  string
	Age       int
}

func MyFunc(params FuncParams) error {
	// Do something
	fmt.Println("Passed parameters:", params)
	return nil
}

// Example of a Variadic Function
// ------------------------------
func addNums(nums ...int) int {
	var res int
	for _, n := range nums {
		res += n
	}
	return res
}

// Example of Function With Multiple Return Values
// -----------------------------------------------
func divmod(num, den int) (int, int, error) {
	if den == 0 {
		return 0, 0, errors.New("cannot divide by 0")
	}
	return num / den, num % den, nil
}

// Example of Function With Named Return Values
// --------------------------------------------
func divmodNamed(num, den int) (res int, mod int, err error) {
	if den == 0 {
		err = errors.New("cannot divide by 0")
		// Returning multiple values
		return res, mod, err
	}
	res, mod = num/den, num%den
	return res, mod, nil
}

// Example of Declaring a Function Variable
// ----------------------------------------
func f1(a string) int {
	return len(a)
}

func f2(a string) int {
	sum := 0
	for _, v := range a {
		sum += int(v)
	}
	return sum
}

// Example of a Simple Calculator With Functions
// ---------------------------------------------
func add(i int, j int) int  { return i + j }
func sub(i int, j int) int  { return i - j }
func mul(i int, j int) int  { return i * j }
func divs(i int, j int) int { return i / j }

// FOR WINDOWS:
//  To Run:                 make run-win
//                          go run Functions\src\main.go
//  To Build:               make build-win
//                          go build -o Functions\bin\Functions.exe Functions\src\main.go
//  To Run after Build:     .\bin\Functions.exe
//                          .\Functions\bin\Functions.exe
//  Try Build + Run:        make try-win
//                          go build -o Functions\bin\Functions.exe Functions\src\main.go && .\Functions\bin\Functions.exe && rm .\Functions\bin\Functions.exe

// FOR LINUX:
//  To Run:                 make run
//                          go run Functions/src/main.go
//  To Build:               make build
//                          go build -o Functions/bin/Functions Functions/src/main.go
//  To Run after Build:     ./bin/Functions
//                          ./Functions/bin/Functions
//  Try Build + Run:        make try
//                          go build -o Functions/bin/Functions Functions/src/main.go && ./Functions/bin/Functions && rm ./Functions/bin/Functions
