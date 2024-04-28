# Blocks, Shadows, and Control Structures

---

- [Blocks](#blocks)
- [Shadowing Variables](#shadowing-variables)
- [The Universe Block](#the-universe-block)
- [`if`](#if)
- [`for`](#for)
  - [Complete C-Style `for`](#complete-c-style-for)
  - [Condition-Only `for` (`while`-Style)](#condition-only-for-while-style)
  - [Infinite `for`](#infinite-for)
    - [`break` and `continue`](#break-and-continue)
  - [`for-range` Statement](#for-range-statement)
    - [Iterating Over Maps](#iterating-over-maps)
    - [Iterating Over Strings](#iterating-over-strings)
    - [The `for-range` Value Is A Copy](#the-for-range-value-is-a-copy)
  - [Labeling `for` Statements](#labeling-for-statements)
  - [Choosing The Right `for` Statement](#choosing-the-right-for-statement)
- [`switch`](#switch)
  - [Blank `switch`](#blank-switch)

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

## `if`

- Similar to `if` in other languages
- **But no parenthesis around the conditions**

```go
// Example of `if`
// ---------------
import "math/rand"

n := rand.Intn(10)
if n == 0 {
    fmt.Println(n, ": That is too low!")
} else if n > 5 {
    fmt.Println(n, ": That is too big!")
} else {
    fmt.Println(n, ": That is a good number!")
}
```

- Any variables declared within the `{}` are scoped within that block
- **Any variables declared at the beginning of `if` are scoped to the `if`, `else if`, and `else` blocks**
  - With this approach, the variable related to the condition only lives for the duration of the condition block
  - Allows to create variables only when they are needed
  - *Once the block ends, the variable is undefined*

```go
// Example of `if` With Scoped Variables
// -------------------------------------
import "math/rand"

if m := rand.Intn(10); m == 0 {
    fmt.Println(m, ": That is too low!")
} else if n > 5 {
    fmt.Println(m, ": That is too big!")
} else {
    fmt.Println(m, ": That is a good number!")
}

// This throws an error: m is undefined here
// fmt.Println(m)
```

- **NOTE**
  - We could put any simple statement before the comparison
  - Not just variable declaration
  - But don't do that
  - Leave it for only declaring variables to be scoped for the `if-else` blocks
  - Anything else would be confusing
  - **Variable declared for the `if-else` block will shadow any existing outside variables**

## `for`

- `for` is the only loop construct available in Go
- There are 4 ways of `for` loops
  - Complete *C-Style `for`*
  - *Condition-Only `for` (`while`-Style)*
  - *Infinite `for`*
  - *`for-range`*

### Complete C-Style `for`

- Similar format to other C-based languages

```go
for init; condition; increment {
    body
}
```

- Similar to `if`, there are no `()` around `init; condition; increment`
- **`init`**
  - *Must use `:=` to initialize the variable: `var` is not legal here*
  - *Declared variable will shadow any existing variable*
- **`condition`**
  - Must be an expression that evaluates to a `bool`
  - Checked immediately *before* each iteration of the loop
  - If `true`, the loop-body executes
- **`increment`**
  - Typically something like `i++`
  - Any assignment is also valid
  - Runs immediately *after* each iteration of the loop

```go
// Example of C-Style for-loop
// ---------------------------
for i := 0; i < 10; i++ {
    fmt.Printf("%d ", i)
}
```

- **Any of the 3 header element of the loop can be left-out**
- `init` can be based on a value calculated before the loop

```go
// Example of C-Style for-loop: No init
// ------------------------------------
i := 0
for ; i < 10; i++ {
    fmt.Printf("%d ", i)
}
```

- `increment` can be based on something more complicated in the body

```go
// Example of C-Style for-loop: No increment
// -----------------------------------------
for i := 0; i < 10; {
    fmt.Printf("%d ", i)
    if i % 2 == 0 {
        i++
    } else {
        i += 2
    }
}
```

- `condition` could be handled inside the loop using `if`

```go
// Example of C-Style for-loop: No condition
// -----------------------------------------
for i := 0; ; i++ {
    if i == 10 {
        break
    }
    fmt.Printf("%d ", i)
}
```

### Condition-Only `for` (`while`-Style)

- When both `init` and `increment` are removed, also remove all `;`
- **This leaves `for` to look and act like a `while` statement**
- The downside is that `i` is not scoped to the `for` block

```go
i := 0
for i < 10 {
    fmt.Printf("%d ", i)
    i++
}
```

### Infinite `for`

- If we remove all `init`, `condition`, and `increment`, we get an *infinite for`*
- This is an infinite loop

```go
// Example of an Infinite Loop
// ---------------------------
for {
    fmt.Println("Hello!")
}
```

#### `break` and `continue`

- Infinite loops are not good programming by themselves
- But sometimes, they can be useful when combined with `break` and `continue`
- **We can break from the loop using `break`**
  - Once reached, exits the loop immediately
  - It can also be used with any other format of `for`
- **NOTE**
  - Go does not have a `do-while` loop
  - **With *Infinite Loop*, `if`, and `break`, we can similate a `do-while` loop**

```go
// Simulating a Do-While Loop in Go
// --------------------------------
j := 0
for {
    fmt.Println("\tThis runs at least once")
    j++
    if j == 1 {
        break
    }
}
```

- **`continue` allows to skip an iteration**
  - Proceeds directly to the next iteration
  - `continue` can make codes much easier to reason and understand
  - It is also more idiomatic to use `continue`
  - Go encourage short `if` bodies
  - It can make the code easier to reason and understand

```go
// Fizzbuzz: Without Using `continue`
// ----------------------------------
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
```

### `for-range` Statement

- Allows to iterate over elements contained in some container types
- Similar to Python's `for x in range(n)`
- Can be used with *strings, arrays, slices, maps, channels*
- **Can only be used on *built-in compound types***
  - Also user-defined types based on them

```go
// Using for-range With Slices
// ---------------------------
evenVals := []int{2, 4, 6, 8, 10, 12}
for i, v := range evenVals {
    fmt.Printf("%d/%d ", i, v)
}
```

- **With `for-range`, we get 2 loop variables**
  - The position of the data in the container
  - The value at that position
- *The idiomatic names of these variables depend on what is being iterated over*
  - For Array/Slice/String: Typically `i`
  - For Map/Struct: Typically `k`
  - Second variable is typically `v`
  - For simple loops, single-letter variables work fine
  - For complex loops, be more descriptive
- **If any of the variable is not going to be used, we can ignore it**
  - If so, we assign it to `_` as a *discard variable*
  - Otherwise, Go requires all declared variables to be used

```go
// Using for-range With Slices Without `i`
// ---------------------------------------
evenVals := []int{2, 4, 6, 8, 10, 12}
for _, v := range evenVals {
    fmt.Printf("%d ", v)
}
```

- **If we just want the key, we can ignore `v` completely**
  - Mostly used approach when a map is being used as a set
  - Can also be used for iteraring over arrays and slices but rare

```go
// Using for-range With Slices Without `v`
// ---------------------------------------
uniqueNames := map[string]bool{
    "Fred": true,
    "Raul": true,
    "John": true,
}
for k := range uniqueNames {
    fmt.Printf("%s", k)
}
```

#### Iterating Over Maps

- **When iterating over Maps, the order is not guaranteed**
- The order of keys and values will vary
- This is actually a security feature
  - People could write code that assumed the iteration order is constant
  - This assumption can break codes at weird times
  - Also, it could be attacked with a *Hash DoS* attack
- 2 changes to the `map` implementation in Go
  - Includes a random number generated every time a map variable is created
  - Order of `for-range` over map is made to vary a bit each time
- **Exception: `fmt.Println()` always output maps with keys in ascending order**

```go
// Map-Iteration-Order Varies
// --------------------------
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
```

#### Iterating Over Strings

- We can also use strings with `for-range`
- `for-range` accesses the *runes* in a string in order
  - `i` - Number of bytes
  - `v` - The rune

```go
// Using for-range With Strings
// ----------------------------
greetings := []string{"Hello!", "Hi π!"}

// Iterating over the slice
fmt.Println("index\tchar\tstring(char)")
for _, greeting := range greetings {
    // Iterating over the string
    for i, char := range greeting {
        fmt.Println(i, "\t", char, "\t", string(char))
    }
    fmt.Println()
}
fmt.Println()
```

- **NOTE**
  - The char value for `π` is 960, much larger than a byte
  - Also, index 4 is skipped for `"Hi π!"`
  - **Iterating over a string with `for-range` iterates over *runes*, not *bytes***
  - **With `for-range`, multibytes runes UTF-8 representation are converted to an int32**
    - The offset is incremented by the number of bytes in the rune
    - **If a byte does not represent a valid UTF-8 value, the hex value is returned instead**

#### The `for-range` Value Is A Copy

- **Each time `for-range` iterates, it *copies* the value into the iteration variables**
- ***Modifying the variables will not affect the values in the compound type***

```go
// Modifying for-range Loop Variables: No effect on Compound
// ---------------------------------------------------------
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
```

- **NOTE**
  - **Before Go 1.22, the variables are created once and reused**
  - **Since Go 1.22, the variables are created for each iteration**
  - This prevents some common bugs with goroutines
  - This is a backward-breaking change
  - Can enable behavior by specifying Go version in `go.mod`

### Labeling `for` Statements

- By default, `break` and `continue` applies to the closest `for`
- For nested loops, we can manipulate at which point the loop picks up using *labels*
- The label is always indented by `go fmt` to same-level as surrounding braces for the current block
  - Easier to notice

```go
    // Using for-range With Labels
    // ---------------------------
    greetings2 := []string{"Hello!", "Hi π!"}

    // Iterating over the slice
    fmt.Println("index\tbytes\tstring(rn)")
outerLabel:
    for _, greeting := range greetings2 {
        // Iterating over the string
        for i, rn := range greeting {
            fmt.Println(i, "\t", rn, "\t", string(rn))
            if rn == 'l' {
                // Go to outerLabel
                continue outerLabel
            }
        }
        fmt.Println()
    }
```

- **NOTE: Nested `for`-loops with labels are rare**
  - They are kinda similar to `goto` statement
  - They are moslty used in the following way

```go
outerLabel:
    for _, outerVal := range outereValues {
        for _, innerVal := range outerVal {
            if invalidSitation(innerVal) {
                continue outerLabel
            }
        }
    }
```

### Choosing The Right `for` Statement

- **Most of the time, *`for-range`* is used**
  - Best to use with strings
  - Also works well with slices, maps, and channels
  - Favor when working with built-in compound types
  - Avoids great deal of boilerplate codes with the other `for` styles
- **Use *C-style `for`* when iterating from a given first to last element**
  - Easier to detect the start and end points
- In general, use the styles that is easier to understand

```go
// for-range vs for C-Style
// ------------------------
evenVals := []int{2, 4, 6, 8, 10}

// Using for-range
for i, v := range evenVals {
    if i == 0 or i == 1 {
        continue
    }
    if i == len(evenVals)-1 {
        break
    }
    fmt.Println(i, v)
}

// Using for C-Style
for i := 2; i < len(evenVals)-1; i++ {
    fmt.Println(i, evenVals[i])
}
```

- **NOTE: This pattern does not work for skipping over the beginning of a string**
  - A standard `for`-loop does not work well with multibyte characters
  - **Need to use `for-range` to properly process runes instead of bytes**
- ***Infinite-`for`* and *`while`-style* are used less frequently**
  - *`while`-style* is useful when looping based on a calculated value
  - *Infinite-`for`* can be used to simulate a `do-while` loop
  - *Infinite-`for`* can be used to implement some *iterator* pattern

## `switch`

- Go has a `switch` statement
- At first glance, looks very similar to other languages's `switch`
  - But Go extends the values that can be switched than in other languages
- **Similar to `if`, `()` are not needed around the condition**
- **We can also declare variables to be scoped to the `switch` statement**
- No `{}` around the body of each `case` or `default`
  - It still support multiple lines though
  - Considered part of the same block
  - **New variables declared within a `case` are scoped to the case**
- There is no need for `break` statement
  - Each case is mutually-exclusive
  - There is no *cascading* effect between successive `case`
  - Similar to *Ruby* and *Pascal*
  - **For multiple matches, separate the cases with commas**
  - **If a `case` is empty, it does nothing**
  - **But `break` can still be used to exit early**
  - But using `break` might indicate the need to refactor

```go
// Using switch in Go
// ------------------
words := []string{"a", "cow", "smile", "gopher", "octops", "anthropologist"}
for _, word := range words {
    // switch-scoped variable
    switch size := len(word); size {
    case 1, 2, 3, 4:
        // Multiple matches
        fmt.Println("-", word, "is a short word")
    case 5:
        // Case-scoped variable
        wordLen := len(word)
        fmt.Println("-", word, "is exactly the right length", wordLen)
    case 6, 7, 8, 9:
        // Empty cases do nothing
        // Not cascading into default
    default:
        fmt.Println("-", word, "is too long")
    }
}
```

- **`break` could also be used in tandem with `switch` inside `for`**
  - Can be used to break out of the `for`-loop using label
  - Without a `for`-label, it might try to break out of the `case` instead

```go
    // Using switch Within for-loop
    // ----------------------------
switchLoop:
    for i := 0; i < 10; i++ {
        switch i {
        case 0, 2, 4, 6:
            fmt.Println(i, "is even")
        case 3:
            fmt.Println(i, "is divisible by 3 but not by 2")
        case 7:
            fmt.Println(i, "-> Exiting the loop. Good bye!")
            break switchLoop
        default:
            fmt.Println("-- You have reached the default case --")
        }
    }
```

- **NOTE: Go has a `fallthrough` keyword**
  - Allows a `case` to *continue* unto the next `case`
  - But this is not used often
  - If `fallthrough` is needed, it might be good to restructure the code
  - It is only used in very rare cases
- **We can also switch on any *type* that can be compared with `==`**
  - All built-in types except `slice`, `map`, `channel`, `func`, `struct`

### Blank `switch`
