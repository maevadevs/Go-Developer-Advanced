package main

import (
	"fmt"
	"math/cmplx"
)

// This is the main entry of the application.
func main() {
	// Example of Integers
	fmt.Println("Example of Integers:")

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

	// Example of Floats
	fmt.Println("Example of Floats:")

	var length float32 = 24.68
	var avogadro float64 = 6.2214e23
	var hexaFloat float64 = 0x12.34p5
	var bankBalance float64 = 10_314.56

	fmt.Println("length =", length)
	fmt.Println("avogadro =", avogadro)
	fmt.Println("hexaFloat =", hexaFloat)
	fmt.Println("bankBalance =", bankBalance)

	fmt.Println()

	// Example of Runes
	fmt.Println("Example of Runes:")

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
	fmt.Println("Example of String Literals:")

	// Regular string
	var greetings string = "Hello World!"
	// Using Escapes
	var greetingsLong string = "Greetings and \n\"Salutations\"!"
	var goPath string = "https:\\\\go.dev"
	// Using Raw String Literal
	var greetingsRaw string = `Greetings and
    "Salutations"!`
	var goPathRaw string = `https://go.dev`

	fmt.Println("greetings =", greetings)
	fmt.Println("greetingsLong =", greetingsLong)
	fmt.Println("goPath =", goPath)
	fmt.Println("greetingsRaw =", greetingsRaw)
	fmt.Println("goPathRaw =", goPathRaw)

	fmt.Println()

	// Example of Boolean Literals
	fmt.Println("Example of Boolean Literals:")

	// Declaration: Default to false
	var flag bool
	// Declaration and initialization
	isAdult := true

	fmt.Println("flag =", flag)
	fmt.Println("isAdult =", isAdult)

	fmt.Println()

	// Example of Complex Number
	fmt.Println("Example of Boolean Literals:")

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
	fmt.Println("Example of Type Conversion:")

	var myInt int = 10
	var myFloat float64 = 30.2

	// Explicit conversion is required
	var mySumF float64 = float64(myInt) + myFloat
	var mySumI int = myInt + int(myFloat)

	fmt.Println("mySumF =", mySumF)
	fmt.Println("mySumI =", mySumI)
}

// FOR WINDOWS:
//  To Run:                 make run-win
//                          go run 02-Types-And-Declarations\src\main.go
//  To Build:               make build-win
//                          go build -o 02-Types-And-Declarations\bin\Types.exe 02-Types-And-Declarations\src\main.go
//  To Run after Build:     .\bin\Types.exe
//                          .\02-Types-And-Declarations\bin\Types.exe
//  Try Build + Run:        make try-win
//                          go build -o 02-Types-And-Declarations\bin\Types.exe 02-Types-And-Declarations\src\main.go && .\02-Types-And-Declarations\bin\Types.exe && rm .\02-Types-And-Declarations\bin\Types.exe

// FOR LINUX:
//  To Run:                 make run
//                          go run 02-Types-And-Declarations/src/main.go
//  To Build:               make build
//                          go build -o 02-Types-And-Declarations/bin/Types 02-Types-And-Declarations/src/main.go
//  To Run after Build:     ./bin/Types
//                          ./02-Types-And-Declarations/bin/Types
//  Try Build + Run:        make try
//                          go build -o 02-Types-And-Declarations/bin/Types 02-Types-And-Declarations/src/main.go && ./02-Types-And-Declarations/bin/Types && rm ./02-Types-And-Declarations/bin/Types
