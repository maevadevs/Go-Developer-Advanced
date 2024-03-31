# Types And Declarations

---

- [Predeclared Types](#predeclared-types)
- [Zero Values](#zero-values)
- [Literals](#literals)
  - [Integer Literal](#integer-literal)
  - [Float Literal](#float-literal)
  - [Rune Literal](#rune-literal)
  - [String Literal](#string-literal)
- [Boolean Types](#boolean-types)
- [Numeric Types](#numeric-types)
  - [Integer Types](#integer-types)
    - [Special Integer Types](#special-integer-types)
    - [Which Integer Type To Use](#which-integer-type-to-use)
    - [Integer Operators](#integer-operators)
  - [Floating-Point Types](#floating-point-types)
    - [Which Floating-Point Type To Use](#which-floating-point-type-to-use)
    - [Float Operators](#float-operators)
  - [Complex Type](#complex-type)
    - [Complex Operators](#complex-operators)
    - [Go For Numerical Computing](#go-for-numerical-computing)
- [Strings and Runes](#strings-and-runes)
  - [String Operators](#string-operators)
  - [Runes](#runes)
- [Explicit Type Conversion](#explicit-type-conversion)
- [Literals Are Untyped](#literals-are-untyped)

---

- Write programs in a way that makes intentions clear
- Some particular approaches produce clearer codes
- Built-in types in Go and how to declare them
- Go does things differently
- Subtle differences with other languages
- **NOTE: All declared variables in Go must be used**

## Predeclared Types

- Go has many types built into the language
- Similar to types found in other languages
  - Booleans
  - Integers
  - Floats
  - Strings

## Zero Values

- Go assigns a default value to declared variables with no initial values
  - Explicit *Zero Value*
  - Makes code clearer
  - Removes sources of bugs
  - [More details on the docs page](https://go.dev/ref/spec#The_zero_value)

## Literals

- Explicitly specified number, character, or string
- 4 common kinds of literals
  - Integer literal
  - Float Literal
  - Rune Literal
  - String Literal
- **Literals are considered *untyped***
  - Situations where the type is not explicitly declared
  - In those cases, Go uses the *default type* for a literal
  - *If a type cannot be determined from an expression (e.g. Should `3` be used as a Float or an Integer?), the literal defaults to a default type*

### Integer Literal

- Sequence of numbers
- Base 10 by default
- **Different prefixes can be used to specify other bases**
  - `0b` => Binary
  - `0x` => Hexadecimal
  - `0o` => Octal: `0` can also be used for *Octal* but it is confusing so do not use it
  - Prefix can be uppercase or lowercase
- **Optional `_` can be used as separators**
  - But cannot be at the beginning or end
  - Cannot be next to each other
  - Can be anywhere else in the number
  - It is better to only use them to improve readability for 1000s-boundaries
  - Or at byte-boundaries for binary, octal, or hexadecimal
- **In practice, use Base-10 to represent integers**
  - Octal is mostly only used for POSIX permission flag
  - Hexadecimal and Binary are sometimes used for bit filters or networking and infrastructure applications

```go
package main

import "fmt"

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
}
```

### Float Literal

- Sequence of numbers
- Has a decimal point to indicate the fractional portion
- Can also have an exponent `e` or `E` with positive of negative number
- Can also be written in hexadecimal
  - E.g. `0x12.34p5`
  - `p` indicates the exponent
  - This is equivalent to `582.5` in decimal
- Optional `_` can be used as separators
  - Same rules as with integers
- **In practice, use Base-10 to represent floats**

```go
package main

import "fmt"

// This is the main entry of the application.
func main() {
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
}
```

### Rune Literal

- Represent a character
- Surrounded by single-quotes `''`
  - **In Go, `'` and `"` are not interchangeable**
  - `'` is for Runes
  - `"` is for Strings
- Rune literals can be written in many ways

Rune Literal|Example
:-|-
Single Unicode Character|`'a'`
8-Bit Octal Number|`'\141'`
8-Bit Hexadecimal Number|`'\x61'`
16-Bit Hexadecimal Number|`'u0061'`
32-Bit Unicode Number|`'\U00000061'`

- There are several backslash-escaped runes as well
  - Newline `\n`
  - Tab `\t`
  - Single-quote `\'`
  - Backslash `\\`
- **Avoid using any of the numeric escapes for rune literals**
  - Unless the context makes the code clearer

```go
package main

import "fmt"

// This is the main entry of the application.
func main() {
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
}
```

### String Literal

- 2 ways: *Interpreted String Literal `""`* and *Raw String Literal <code>``</code>*
  - ***Interpreted String Literal* is surrounded by double-quotes `""`**
    - Contains zero or more rune literals
    - They interpret rune literals into single characters
    - The following cannot appear in *Interpreted String Literal*
      - Unescaped backslash `\`
      - Unescaped newlines `n`
      - Unescaped double-quote `"`
  - **In these cases, using *Raw String Literal* is more appropriate**
    - Delimited with backticks <code>``</code>
    - **Can contain any character except backticks**
    - No escape character can be applied
    - All characters are included as-is

```go
package main

import "fmt"

// This is the main entry of the application.
func main() {
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
}
```

- **NOTE: `\'` is not legal in string literals**
  - It gives an error `error: unknown escape sequence: '`
  - It needs to be escaped twice instead `\\'`

## Boolean Types

- `bool` represents boolean types
- Either `true` or `false`
- Zero-Value is `false`

```go
package main

import "fmt"

// This is the main entry of the application.
func main() {
    // Example of Boolean Literals
    fmt.Println("Example of Boolean Literals:")

    // Declaration: Default to false
    var flag bool
    // Declaration and initialization
    isAdult := true

    fmt.Println("flag =", flag)
    fmt.Println("isAdult =", isAdult)
}
```

## Numeric Types

- 12 Numeric Types grouped into 3 categories
  - 8 Integers
    - `int8`
    - `int16`
    - `int32` / `int`
    - `int64` / `int`
    - `uint8` / `byte`
    - `uint16`
    - `uint32` / `uint`
    - `uint64` / `uint`
  - 2 Special Integers
    - `rune`
    - `uintptr`
  - 2 Floating-Points
    - `float32`
    - `float64`
- Also a Complex Type: `complex`
- Some are used more often than others

### Integer Types

- Both signed and unsigned integer types
- From 1 byte to 8 bytes in length
- **Zero-Value = `0`**
- **Literal Default Type: `int`**

Integer Type|Range
:-:|:-:
`int8`|$-2^7$ ... $2^7-1$
`int16`|$-2^{15}$ ... $2^{15}-1$
`int32`|$-2^{31}$ ... $2^{31}-1$
`int64`|$-2^{63}$ ... $2^{63}-1$
`uint8`|$0$ ... $2^8-1$
`uint16`|$0$ ... $2^{16}-1$
`uint32`|$0$ ... $2^{32}-1$
`uint64`|$0$ ... $2^{64}-1$

#### Special Integer Types

- `byte` is alias for `uint8`
- `int` is alias to:
  - `int32` on 32-bit system
  - `int64` on 64-bit system
- `uint` is alias to:
  - `uint32` on 32-bit system
  - `uint64` on 64-bit system
- **NOTE: Some 64-bit architectures use `int32` as `int` and `uint32` as `uint`**
  - `amd64p32`
  - `mips64p32`
  - `mips64p321e`
- **Comparing or operating between values of type `int`/`uint` and `int32`-`int64`/`uint32`-`uint64` is a compile-time error**
  - Because `int` and `uint` are not consistent across platforms
  - **Explicit type-conversion is necessary to do so**
- **NOTE: `rune` and `uintptr` are also *Special Integer Types***

#### Which Integer Type To Use

- Follow 3 simple rules

1. When working on binary file or networking protocols, if integer size is specified, use equivalent size
2. When writing library function that should work on any integer type, use Generics to represent any integer types
3. In all other cases, use `int`

- **NOTE: Some legacy codes might define 2 separate functions using `int64` and `uint64`**
  - Those APIs were created before Go supported Generics
  - E.g. Go Standard Library `strconv.FormatInt()` and `strconv.FormatUint()`

#### Integer Operators

- Support usual arithmetic operators
  - Addition `+` and `+=`
  - Subtraction `-` and `-=`
  - Multiplication `*` and `*=`
  - Division `/` and `/=`
  - Modulo `%` and `%=`
- **Integer division results in an integer**
  - Truncation toward `0`
  - To get Float, use type-conversion
- **Division by `0` causes a `panic`**
- Support usual comparison operators
  - `==` and `!=`
  - `>` and `>=`
  - `<` and `<=`
- Also support bit-manipulation
  - SHIFT-LEFT `<<` and `<<=`
  - SHIFT-RIGHT `>>` and `>>=`
  - BITWISE-AND `&` and `&=`
  - BITWISE-OR `|` and `|=`
  - BITWISE-XOR `^` and `^=`
  - BITWISE-ANDNOT `&^` and `&^=`

### Floating-Point Types

- 2 Floating-Point types
- IEEE 754 standard
- **Zero-Value = `0.0`**
- **Literal Default Type: `float64`**

Float Type|Range|Precision
:-:|:-:|:-
`float32`|$±1.4012984e^{-45}$ ... $±3.4028234e^{38}$|6-7 digits
`float64`|$±1.7976931e^{308}$ ... $±4.9406564^{324}$|15-16 digits

#### Which Floating-Point Type To Use

- Use `float64` unless needing compatibility
  - Helps mitigate accuracy issues with `float32`
  - Memory size should not be an issue unless profiler determines significant issues
- **In many cases, floating-point should not be used at all**
  - Huge range but only nearest approximation
  - Can only be used in situations where inexact values are acceptable: Graphics, Statistics, Scientific Operations
  - The rules of floating-points must be well understood
- **Do not use floating-points for any monetary applications or when needing accuracy**
  - Use third-party modules for handling exact decimal values instead

#### Float Operators

- Support usual arithmetic operators, **except Modulo**
  - Addition `+` and `+=`
  - Subtraction `-` and `-=`
  - Multiplication `*` and `*=`
  - Division `/` and `/=`
- **Float division has some properties**
  - Truncation toward `0`
  - Division of Non-Zero by `0` results in `±Inf`
  - Division of `0` by `0` results in `NaN`
- Support usual comparison operators
  - `==` and `!=` **but do not use them!**
    - Instead, define max allowed variance *epsilon*
    - Then see if the diff is less than that
    - Check the [*Floating-Point Comparison Guide*](https://floating-point-gui.de/errors/comparison/)
  - `>` and `>=`
  - `<` and `<=`
- Also support bit-manipulation
  - SHIFT-LEFT `<<` and `<<=`
  - SHIFT-RIGHT `>>` and `>>=`
  - BITWISE-AND `&` and `&=`
  - BITWISE-OR `|` and `|=`
  - BITWISE-XOR `^` and `^=`
  - BITWISE-ANDNOT `&^` and `&^=`

### Complex Type

- First-class support for Complex Numbers
- Not typically used
- 2 complex number types
  - `complex64` - Use `float32` for `real` and `imag`
  - `complex128` - Use `float64` for `real` and `imag`
- **Zero-Value = `0.0` for both `real` and `imag`**

```go
var complexNum = complex(20.5, 10.6)
```

- **Assigned Default Type of value follows a specific rule:**
  - If params are untyped constants/literals => Untyped complex literal: `complex128`
  - If params are `float32` => `complex64`
  - If one param is `float32` and the other is untyped constant/literal but can fit in `float32` => `complex64`
  - Else => `complex128`

#### Complex Operators

- Support usual arithmetic operators for floats
- Support usual comparison operators
  - `==` and `!=` **but do not use them!**
    - Instead, define max allowed variance *epsilon*
    - Then see if the diff is less than that
    - Check the [*Floating-Point Comparison Guide*](https://floating-point-gui.de/errors/comparison/)
  - `>` and `>=`
  - `<` and `<=`
- Real portion is available from calling `real(complexNum)`
- Imaginary portion is available from calling `imag(complexNum)`
- `math/cmplx` has additional functions for manipulating Complex Numbers

```go
package main

import "fmt"

// This is the main entry of the application.
func main() {
    // Example of Complex Number
    fmt.Println("Example of Boolean Literals:")

    x := complex(2.5, 3.1)
    y := complex(10.2, 2)

    fmt.Println("x =", x)
    fmt.Println("y =", y)
    fmt.Println("x + y =", x + y)
    fmt.Println("x - y =", x - y)
    fmt.Println("x * y =", x * y)
    fmt.Println("x / y =", x / y)
    fmt.Println("real(x) =", real(x))
    fmt.Println("imag(x) =", imag(x))
    fmt.Println("cmplx.Abs(x) =", cmplx.Abs(x))
}
```

#### Go For Numerical Computing

- **Go is not popular for Numerical Computing**
- Limited adoption because of other features
  - Matrix
  - Vectors
- Libraries have to use inefficient replacements (slices of slices)
  - [Gonum](https://www.gonum.org/)
- Complex Number can be used for Mandelbrot or Quadratic Equation solver

## Strings and Runes

- String is a built-in data type
- Supports Unicode: We can put any unicode characters in strings
- Zero-Value is `""`
- Default Type is `string`
- **Strings are immutable**

### String Operators

- We can use comparison operators: `==`, `!=`
- Ordering operators: `>`, `>=`, `<`, `<=`
- Concatenation: `+`, `+=`

### Runes

- Represent a single code-point
- Used for representing a character
- Alias for `int32` type
  - But use `rune` for character, not `int32`
  - They are same for the compiler, but bad programming practice to mismatch
  - Helps to clarify the intent in the code
- Default type is `rune`

```go
// Good
var fNameInitial rune = 'M'
// Bad - Legal but confusing
var lNameInitial int32 = 'R'
```

## Explicit Type Conversion

- Most languages have *Automatic Type Promotions/Upgrades*
- But the rule to apply this can get complicated and lead to unexpected results
- **Go does not allow *Automatic Type Promotions***
  - **Type conversions must be explicit**
  - Even for integers and floats
- Type conversion is done by calling the type as a function on the value
  - **For any values that are not of the same type: At least one must be converted before they can be used together**
  - The same applies for same type but different sizes (E.g. `int8` and `int32`)

```go
package main

import "fmt"

// This is the main entry of the application.
func main() {
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
```

- Due to this strictness, we cannot treat other types as *booleans*
  - In other languages, some values are trrated as *truthy* or *falsy*
  - **Go does not allow *truthy* and *falsy* either**
  - **No other type can be converted to booleans, either implicitly or explicitly**
  - **Must use comparison operators to get back booleans**
    - E.g. A typical check that is used a lot is `!= nil`

## Literals Are Untyped
