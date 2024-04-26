# Blocks, Shadows, and Control Structures

---

- [Blocks](#blocks)
- [Shadowing Variables](#shadowing-variables)
- [The Universe Block](#the-universe-block)

---

- About programming logic and organization
- Blocks can control when an identifier is available
- Control structures: `if`, `for`, `switch`
- Go also has `goto` for some situations

## Blocks

- We can declare variables in lots of places
  - Outside of functions
  - Parameters of functions
  - Local variables of functions
- **Block** - Each place where a declaration occurs
- **Package Block** - Anything declared outside of a function
- **File Block** - Names that can be used with `import`
- **All variables defined at the top-level of a function are in a block**
  - This includes *Function Parameters*
- **Any set of `{}` within a function defines another block**
  - Control structures define block of their own
- **We can access identifiers defined in an outer block from within an inner block**
  - With same identifier names but different block, the name is *shadowed*

## Shadowing Variables

- **Shadowing Variable**
  - A variable with the same name as another variable in the containing block
  - **As long as the shadowing variable exist, the shadowed variable is inaccessible**
- **The act of declaring a variable of the same name as the containing block creates a shadowing variable**
  - The shadowing variable is accessed with the name within the block it was declared
  - A closing `}` ends the scope of the shadowing variable

```go
// Example of Variable Shadowing
// -----------------------------
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
```

- **Accessing the same variable after the end of the shadowing scope returns the previously-shadowed variable**
  - This variable did not disappear or got reassigned
  - It was simply inacessible as long as the shadowing variable was in scope
- **It is very easy to accidentaly shadow a variable when using `:=`**
  - We can use `:=` to create multiple variables at once
  - *Not all variables needs to be new when using `:=`, just one is enough*
  - **`:=` reuses only variables that are declared in the current block**

```go
// Example of Variable Shadowing With :=
// -------------------------------------
if x > 5 {
    x, y := 5, 20
    fmt.Println("Inside the block, x =", x, "and y =", y)
}

fmt.Println("Outside the block, x =", x)
```

- **When using `:=`, make sure outside variables with the same name does not exist, unless it is an intentional shadowing**
- Also, be careful to not shadow package imports
  - E.g. `math`
  - Once shadowed, package imports are inaccessible until the end of the block

```go
import (
    "fmt"
    "math"
)

// Example of Shadowing Package Names
// ----------------------------------
pi := math.Pi

if float64(x) > pi {
    math := "oops!"

    // This is an error: math.Pi is undefined
    // pi2 := math.Pi

    fmt.Println("\tInside the block, math =", math)
}

fmt.Println("Outside the block, math.Pi =", math.Pi)
```

## The Universe Block

- Go is a small language: Only 25 keywords
- The built-in types, constants, `nil`, and functions are not included in the list
  - Go considers them *Predeclared Identifiers*
  - **They are defined in the *Universe Block*** - The block that contains all other blocks
  - *These names can be shadowed in any other blocks*
- **Never redefine any identifiers in the *Universe* Block**
  - This can create some strange behaviors
  - Could also create some hard-to-find bugs

```go
// Example of Shadowing `true`
// --------------------------
true := 100
```

- **NOTE: `go vet` does not report shadowing as an issue**
- But some 3rd-party tools can detect accidental shadowing
