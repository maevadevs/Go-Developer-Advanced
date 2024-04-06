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
- [`var` vs `:=`](#var-vs-)
  - [Long-Format Declaration With `var`](#long-format-declaration-with-var)
  - [Short-Format Declaration With `var`](#short-format-declaration-with-var)
  - [Declaration Only With `var`](#declaration-only-with-var)
  - [Multiple Declaration](#multiple-declaration)
  - [Walrus Shortcut `:=`](#walrus-shortcut-)
  - [Which Style To Choose](#which-style-to-choose)
- [Using `const`](#using-const)
  - [Typed Const vs Untyped Const](#typed-const-vs-untyped-const)
- [Unused Variables](#unused-variables)
- [Naming Variables And Constants](#naming-variables-and-constants)

---

- Write programs in a way that makes intentions clear
- Some particular approaches produce clearer codes
- Built-in types in Go and how to declare them
- Go does things differently: Subtle differences from other languages
- **NOTE: All declared variables in Go must be used**

## Predeclared Types

- Go has many types built into the language
- Similar to types found in other languages
  - Booleans
  - Integers
  - Floats
  - Strings

## Zero Values

- **Declared variables with no initial values are assigned default *Zero Values***
  - Explicit *Zero Value*
  - Makes code clearer
  - Removes sources of bugs
  - [More details on the docs page](https://go.dev/ref/spec#The_zero_value)

## Literals

- **Explicitly specified number, character, or string**
- 4 common kinds of literals
  - Integer literal
  - Float Literal
  - Rune Literal
  - String Literal
- **Literals are considered *untyped***
  - Situations where the type is not explicitly declared
  - In those cases, Go uses the *default type* for a literal
  - **If a type cannot be determined from an expression (e.g. Should `3` be used as a Float or an Integer?), the literal defaults to a default type**

### Integer Literal

- Sequence of numbers
- Base 10 by default
- **Different prefixes can be used to specify other bases**
  - `0b` => *Binary*
  - `0x` => *Hexadecimal*
  - `0o` => *Octal*: `0` by itself can also be used for *Octal* but it is confusing so do not use it
  - Prefix can be uppercase or lowercase
- **Optional `_` can be used as separators**
  - But cannot be at the beginning or end
  - Cannot be next to each other
  - Can be anywhere else in the number
  - It is better to only use them to improve readability for 1000s-boundaries
  - Or at byte-boundaries for binary, octal, or hexadecimal
- **In practice, use Base-10 to represent integers**
  - Octal is mostly only used for POSIX permission flags
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

- Sequence of numbers with decimal point to indicate the fractional portion
- Can also have an exponent `e` or `E` with positive or negative number
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

- 2 ways:
  - *Interpreted String Literal `""`*
  - *Raw String Literal <code>``</code>*
- ***Interpreted String Literal* is surrounded by double-quotes `""`**
  - Contains zero or more rune literals
  - They *interpret* rune literals into single characters
  - The following cannot appear in *Interpreted String Literal*
    - Unescaped backslash `\`
    - Unescaped newlines `n`
    - Unescaped double-quote `"`
- **In these cases, using *Raw String Literal* is more appropriate**
  - Delimited with backticks <code>``</code>
  - **Can contain any character except backticks**
  - If <code>\`</code> is required, try using <code>\`raw string\` + "`"</code> instead
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
    var sysPath string = "C:\\\\Windows\\System32"
    // Using Raw String Literal
    var greetingsRaw string = `Greetings and
    "Salutations"!`
    var goPathRaw string = `https://go.dev`

    fmt.Println("greetings =", greetings)
    fmt.Println("greetingsLong =", greetingsLong)
    fmt.Println("sysPath =", sysPath)
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
- **Integer Division by `0` causes a `panic`**
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
  - `float32`
  - `float64`
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
    fmt.Println("Example of Complex Literals:")

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
- Supports Unicode
  - We can put any unicode characters in strings
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
  - **But use `rune` for character, not `int32`**
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
  - The same applies for values of same type but different sizes (E.g. `int8` and `int32`)

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
  - In other languages, some values are treated as *truthy* or *falsy*
  - **Go does not allow *truthy* and *falsy* either**
  - **No other types can be converted to booleans, either implicitly or explicitly**
  - **Must use comparison operators to get back booleans**
    - E.g. A typical check that is used a lot is `x != nil`

## Literals Are Untyped

- We cannot add 2 integers of different types
- But we can use integer literals as floating-point

```go
var x float64 = 100
var y float64 = 200.50 * 50
```

- **Literals are untyped**
  - Their types are not enforced until needed
  - This allows literals to be used with any compatible variables
- **However, this only works on compatible/similar types**
  - We cannot assign strings to numeric variables
  - Size limitations also exist
    - Compile-time error if the value overflows the variable type

## `var` vs `:=`

- Go has different ways for declaring variables
- Each style communicates about how the variable is used
- *Variables can be declared at the package-level or function-level*
  - However, it is preferable to do so at the function-level as much as possible

### Long-Format Declaration With `var`

- This is the most verbose format
- Most general case

```go
// Long-format declaration with var
var x int = 100
```

### Short-Format Declaration With `var`

- If the type of the value can be deduced, we can skip the type
- *The type of the variable is deduced from the assigned value*

```go
// Short-format declaration with var
var x = 100
```

### Declaration Only With `var`

- We can declare only, without assigning an initial value
- The type is required
- Assigns the *Zero-Value* of the type as the default value

```go
// Declaration-Only with var
var x int
```

### Multiple Declaration

- This is similar to Python's *tuple assignment*
- The variables can be of the same type

```go
// Multiple declaration and assignment
var age, favNum int = 26, 100
```

- We can also declare-only and assign Zero-Value

```go
// Multiple declaration only
var age, favNum int
```

- If assigning initial value, the variables can be of different types
- The type of the variables are deduced from the assigned values

```go
// Declaring and assigning multiples
var name, age, isAdult = "John", 26, true
```

- We can also declare a list of variables at once
- This approach allows some to have initial values and some not
- We can combine different previous approach

```go
var (
    uFname, uLname string
    uAge = 26
    uIsAdult bool = true
    uFavNum, uWorstNum = 7, -100
)
```

### Walrus Shortcut `:=`

- `:=` can replace `var` declaration with type inference
- `:=` is preferred over `var`

```go
// Equivalent statements
var num = 100
num := 100
```

- We can also use it with multiple variables at once

```go
// Using := With Multiple Variables
name, age, isAdult := "John", 26, true
```

- **NOTE: `:=` can only be used for declaration, not for re-assignments**
  - However, with multiple declarations, it can be used for re-assignment **if at least one variable is new**

```go
// Using := With Multiple Variables & Reassignment
name, age := "John", 15
name, age, isAdult := "Johnny", 26, true
```

- **Limitation of `:=`**
  - It cannot be used to declare variables at the *package-level*
  - `var` must be used for *package-level* variables
  - `:=` can only be used inside functions

### Which Style To Choose

- **Choose what makes the intent clear**
  - `var` is more explicit but can be verbose
  - `:=` is easy to understand but can be confusing if the value is not explicit
  - `var` is the only option for package-level variables
  - `:=` is the most common inside functions
  - Use declaration lists when declaring multiple variables
- Avoid `:=` when:
  - Declaring Package-level variables
  - Initializing to zero-value
  - Assigning an untyped constant
  - Variable type cannot/should not be deduced from the assigned value
    - Prefer `var x byte = 20` over `x := byte(20)`
  - **Sometimes, `:=` creates a new variable than using an existing one (*Shadowing*)**
    - In those cases, it is better to use `var`
- **NOTE: Only use the *Multiple declaration and assignment* style when assigning multiple values from a function return**
- **NOTE: Rarely declare variables outside funnctions**
  - Package-level variables is not a good idea
  - Can be difficult to track the changes made
  - Make it hard to understand the flow of the data in the program
  - Can lead to subtle bugs
  - **Only declare them as *immutable***

## Using `const`

- Allows to declare values as *immutable*
- Usage is very similar to `var`
- But re-assignments result in error

```go
package main

import "fmt"

// Package-level Constants
const pi float64 = 3.1415
const (
    idKey   = "id"
    nameKey = "name"
)
const total = 20 * 10

// This is the main entry of the application.
func main() {
    // Example of Function-level Constants
    fmt.Println("Example of Function-level Constants:")

    const greetings = "Hello"

    fmt.Println("greetings = ", greetings)

    // Example of Package-level Constants
    fmt.Println("Example of Package-level Constants:")

    fmt.Println("pi = ", pi)
    fmt.Println("idKey = ", idKey)
    fmt.Println("nameKey = ", nameKey)
    fmt.Println("total = ", v)

    // This is an error: Constants are immutable
    // pi = pi + 1
    // greeting = "Hi!"
}
```

- **NOTE: `const` in Go is very limited**
  - Can only hold values that the compiler can figure out at compile-time
  - Can be assigned:
    - Numeric Literals
    - `true` / `false`
    - Strings
    - Runes
    - Values for calling `complex()`, `real()`, `image()`, `len()`, `cap()`
    - Expression with operator and preceding value
    - `iota`
  - **We cannot specify that a value calculated at runtime is immutable**
    - `const` is only a way to give names to literals

```go
// This is an error
x := 5
y := 6
const total = x + y
```

- **NOTE: There are no *immutable arrays, slices, maps, structs***
  - We cannot specify that a field in a struct is immutable
  - But within a function, it is clear if a variable is being modified
  - So immutability is less important
  - *Go is call-by-value*: Prevents modification to passed-parameters

### Typed Const vs Untyped Const

- Constants can be typed or untyped
- **Untyped Constant** is the same as literal
  - Has no type on its own
  - Has default type when type cannot be inferred
- **Typed Constant** can be assigned directly only a *"variable"* of the same type
- **Usage depends on the intent**
  - Constants to be used with multiple numeric types => *Keep untyped*
  - Untyped Constants give more flexibility
  - But in some situations, an enforced type is preferrable: E.g. Enumeration with `iota`

```go
// Example of Untyped Constant
const uconst = 100

// Legal usage
var i int = uconst
var f float64 = uconst
var b byte = uconst

// Example of Typed Constant
const tconst int64 = 100

// Legal usage:
// Can only be assigned to int64
// Asigning to any other type is a compile error
var i64 int64 = tconst
```

## Unused Variables

- **Every *locally* declared variables/constants must be used/read at least once**
- It is a compile-time error if a declared variable is not used
- As long as the variable is read once
  - Go will not catch the unused `x = 30`
  - 3rd-party tools can be used to cleanup these

```go
// This code will compile
package main

func main() {
    x := 10
    x = 20
    fmt.Println("x =", x)
    x = 30
}
```

- **NOTE: This rule does not apply to *constants* and *package-level variables***
  - It is better to avoid *package-level variables*
  - Constants in Go are calculated at compile-time (cannot have side-effects)
  - Can be eliminated if unused: They are excluded from the compiled binary

```go
// This code will compile
package main

var note string = "It will not complain if I am unused!"

func main() {
    const PI = 10
}
```

## Naming Variables And Constants

- Identifier names need to start with a letter or `_`
- Name can contain number, letter, `_`
- **Any Unicode character can be used**
- **However, not any unicode character *should* be used**
  - Name should communicate intent
  - Should not be difficult to type on keyboard
  - **Look-alike Unicode code-points can be a nightmare**
    - Even if the look alike, they are completely different variables

```go
// Avoid this!!
func main() {
    ａ = "hello" // Unicode U+FF41: ａ
    a = "bye" // Standard lowercase-A: U+0061
    fmt.Println(ａ) // hello
    fmt.Println(a) // bye
    fmt.Println(ａ == a) // false

}
```

- **NOTE: `_` is rarely used in naming**
  - Idiomatic Go does not use *snake_casing*
  - Only *PascalCasing* and *camelCasing*
  - **`_` is mostly only used as a *discard* variable**
- **NOTE: Do not use *all-caps* when naming constants**
  - Go use the case of the first letter of a name to determine its accessibility (*private* vs *public*)
  - **Keep `const` naming the same as `var`**
- **Within functions, favor short names**
  - Smaller scope = Shorter name
  - Single-letters are common for `for` loops
  - Sometimes, the first letter of the type might be used
  - These short names serve 2 purposes
    - Eliminate repetitive typing
    - Serve as a check on how complicated the code is: *If you cannot keep track, your block is doing too much*
- **Within package block, favor descriptive names**
  - Clarify what the value represent
- For more details on naming, [check the docs](https://google.github.io/styleguide/go/decisions#naming)
