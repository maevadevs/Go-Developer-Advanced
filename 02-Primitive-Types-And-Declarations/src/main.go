package main

import (
	"fmt"
	"math/cmplx"
	"strings"
)

// Example of Package-level Constants
// ----------------------------------
const pi float64 = 3.1415

const (
	idKey   = "id"
	nameKey = "name"
)
const total = 20 * 10 // untyped

// This is the main entry of the application.
func main() {
	// Headers
	fmt.Println()
	fmt.Println(strings.Repeat("-", 100))
	fmt.Println()

	// Example of Integer Literals
	// ---------------------------
	fmt.Println("Example of Integer Literals:")

	var age int = 45
	var blue int = 0x0000ff
	var admin int = 0o777
	var billion int = 1_000_000_000
	var red int = 0xff_00_00

	fmt.Println("age =", age)
	fmt.Println("blue =", blue)
	fmt.Println("admin =", admin)
	fmt.Println("billion =", billion)
	fmt.Println("red =", red)

	fmt.Println()

	// Example of Float Literals
	// -------------------------
	fmt.Println("Example of Float Literals:")

	var length float32 = 24.68
	var avogadro float64 = 6.2214e23
	var hexaFloat float64 = 0x12.34p5
	var bankBalance float64 = 10_314.56

	fmt.Println("length =", length)
	fmt.Println("avogadro =", avogadro)
	fmt.Println("hexaFloat =", hexaFloat)
	fmt.Println("bankBalance =", bankBalance)

	fmt.Println()

	// Example of Rune Literals
	// ------------------------
	fmt.Println("Example of Rune Literals:")

	var gender rune = 'M'
	var newline rune = '\n'
	var tab rune = '\t'
	var quote rune = '\''
	// As 8-Bit Octal Number
	var ninetySeven rune = '\141'
	// As 8-Bit Hexadecimal Number
	ninetySeven = '\x61'
	// As 16-Bit Hexadecimal Number
	ninetySeven = '\u0061'
	// 32-Bit Unicode Number
	ninetySeven = '\U00000061'

	fmt.Println("gender =", gender)
	fmt.Println("newline =", newline)
	fmt.Println("tab =", tab)
	fmt.Println("quote =", quote)
	fmt.Println("ninetySeven =", ninetySeven)

	fmt.Println()

	// Example of String Literals
	// --------------------------
	fmt.Println("Example of String Literals:")

	// Regular string
	var greetings string = "Hello World!"
	// Using Escapes
	var greetingsLong string = "Greetings and \n\"Salutations\"!"
	var sysPath string = "C:\\\\Windows\\System32"
	// Using Raw String Literal
	var greetingsRaw string = `Greetings and
    "Salutations"!`
	var goPathRaw string = `https://go.dev`
	var rawStringWithBacktick string = `Backticks ` + "(`) " + `cannot appear in raw string.
    For that, use "" to contain the ` + "(`)" + `, then concatenate.`

	fmt.Println("greetings =", greetings)
	fmt.Println("greetingsLong =", greetingsLong)
	fmt.Println("sysPath =", sysPath)
	fmt.Println("greetingsRaw =", greetingsRaw)
	fmt.Println("goPathRaw =", goPathRaw)
	fmt.Println("rawStringWithBacktick =", rawStringWithBacktick)

	fmt.Println()

	// Example of Boolean
	// ------------------
	fmt.Println("Example of Boolean:")

	// Declaration: Default to false
	var flag bool
	// Declaration and initialization
	isAdult := true

	fmt.Println("flag =", flag)
	fmt.Println("isAdult =", isAdult)

	fmt.Println()

	// Example of Complex Number
	// -------------------------
	fmt.Println("Example of Complex Number:")

	x := complex(2.5, 3.1)
	y := complex(10.2, 2)

	fmt.Println("x =", x)
	fmt.Println("y =", y)
	fmt.Println("x + y =", x+y)
	fmt.Println("x - y =", x-y)
	fmt.Println("x * y =", x*y)
	fmt.Println("x / y =", x/y)
	fmt.Println("real(x) =", real(x))
	fmt.Println("imag(x) =", imag(x))
	fmt.Println("cmplx.Abs(x) =", cmplx.Abs(x))

	fmt.Println()

	// Example of Type Conversion
	// --------------------------
	fmt.Println("Example of Type Conversion:")

	var myInt int = 10
	var myFloat float64 = 30.2

	// Explicit type conversion is required
	var mySumF float64 = float64(myInt) + myFloat
	var mySumI int = myInt + int(myFloat)

	fmt.Printf("%d + %.1f\n", myInt, myFloat)
	fmt.Println("mySumF =", mySumF)
	fmt.Println("mySumI =", mySumI)

	fmt.Println()

	// Example of Variable Declarations
	// --------------------------------
	fmt.Println("Example of Variable Declarations:")

	// Long-format declaration with var
	var someNum int = 100
	// Declaration-Only with var
	var myNum int
	// Assignment after declaration
	myNum = 77
	// Multiple declaration and assignment
	var userAge, favNum int = 26, 100
	// Multiple declaration only
	var dogAge, dogBirthYear int
	dogAge = 7
	dogBirthYear = 2017
	// Declaration as list
	var (
		uFname, uLname     string
		uAge                    = 26
		uIsAdult           bool = true
		uFavNum, uWorstNum      = 7, -100
	)
	// Using := With Multiple Variables
	sName, sAge := "John", 15
	// Using := With Multiple Variables & Reassignment
	sName, sAge, sIsAdult := "Johnny", 26, true

	fmt.Println("someNum =", someNum)
	fmt.Println("myNum =", myNum)
	fmt.Println("userAge =", userAge)
	fmt.Println("favNum =", favNum)
	fmt.Println("dogAge =", dogAge)
	fmt.Println("dogBirthYear =", dogBirthYear)
	fmt.Println("uFname =", uFname)
	fmt.Println("uLname =", uLname)
	fmt.Println("uAge =", uAge)
	fmt.Println("uIsAdult =", uIsAdult)
	fmt.Println("uFavNum =", uFavNum)
	fmt.Println("uWorstNum =", uWorstNum)
	fmt.Println("sName =", sName)
	fmt.Println("sAge =", sAge)
	fmt.Println("sIsAdult =", sIsAdult)

	fmt.Println()

	// Example of Function-level Constants
	// -----------------------------------
	fmt.Println("Example of Function-level Constants:")

	const greetingsConst = "Hello"

	fmt.Println("greetingsConst =", greetingsConst)

	fmt.Println()

	// Example of Package-level Constants
	// ----------------------------------
	fmt.Println("Example of Package-level Constants:")

	fmt.Println("pi =", pi)
	fmt.Println("idKey =", idKey)
	fmt.Println("nameKey =", nameKey)
	fmt.Println("total =", total)

	// This is an error: Constants are immutable
	// pi = pi + 1
	// greetingsConst = "Hi!"

	fmt.Println()

	// Example of Untyped Constant
	// ---------------------------
	fmt.Println("Example of Untyped Constant:")

	const uConst = 100

	// Legal usage
	var i int = uConst
	var f float64 = uConst
	var b byte = uConst

	fmt.Println("uConst =", uConst)
	fmt.Println("i =", i)
	fmt.Println("f =", f)
	fmt.Println("b =", b)

	fmt.Println()

	// Example of Typed Constant
	// -------------------------
	fmt.Println("Example of Typed Constant:")

	const tConst int64 = 100

	// Legal usage:
	// Can only be assigned to type int64
	// Assigning to any other type is a compile error
	var i64 int64 = tConst

	fmt.Println("tConst =", tConst)
	fmt.Println("i64 =", i64)

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
