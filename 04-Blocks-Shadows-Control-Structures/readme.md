# Blocks, Shadows, and Control Structures

---

- [Blocks](#blocks)
- [Shadowing Variables](#shadowing-variables)
- [The Universe Block](#the-universe-block)
- [`if`](#if)
- [`for`](#for)
  - [C-Style `for`](#c-style-for)
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
  - [Choosing Between `if-else` and `switch`](#choosing-between-if-else-and-switch)
- [`goto`](#goto)
  - [Use of `goto` in Go](#use-of-goto-in-go)

---

- About programming logic and organization
- Blocks can control when an identifier is available (scope)
- Control structures: `if`, `for`, `switch`
- Go also has `goto` for some situations

## Blocks

- We can declare variables in lots of places
  - Outside of functions
  - Parameters of functions
  - Local variables of functions
- **Block**
  - Each place where a declaration occurs
  - **Package Block** - Anything declared outside of a function
  - **File Block** - Names that can be used with `import`
- **All variables defined at the top-level of a function are in a block**
  - This includes *Function Parameters*
- **Any set of `{}` within a function defines another block**
  - **Control structures also define block of their own**
- **We can access identifiers defined in an outer block from within an inner block**
  - With same identifier names but different block, the name is *shadowed*

## Shadowing Variables

- **Shadowing Variable**
  - A variable with the same name as another variable in the containing block
  - **As long as the shadowing variable is in-scope, the shadowed variable is inaccessible**
- **Declaring a variable of the same name as the one in the containing block creates a shadowing variable**
  - The shadowing variable is accessed with the name within the block it was declared
  - A closing `}` ends the scope of the shadowing variable

```go
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
```

- **Accessing the same variable after the end of the shadowing scope returns the previously-shadowed variable**
  - This variable did not disappear or got reassigned
  - It was simply inacessible as long as the shadowing variable was in-scope
- **WARNING: It is very easy to accidentaly shadow a variable when using `:=`**
  - We can use `:=` to create multiple variables at once
  - *Not all variables need to be new when using `:=`, just one is enough*
  - **`:=` reuses only variables that are declared in the current block**

```go
// Example of Variable Shadowing With :=
// -------------------------------------
fmt.Println("Example of Variable Shadowing With :=")
fmt.Println("-------------------------------------")

fmt.Println("Outside the block, x =", x)
if x > 5 {
    x, y := 5, 20
    fmt.Println("\tInside the block, x =", x, "and y =", y)
}
fmt.Println("Outside the block, x =", x)
```

- **When using `:=`, make sure outside variables with the same name do not exist**
  - **Unless it is an intentional shadowing**
- **WARNING: Be careful to not shadow package imports**
  - E.g. `math`
  - Once shadowed, package imports are inaccessible until the end of the block

```go
import (
    "fmt"
    "math"
)

// Example of Shadowing Package Names
// ----------------------------------
fmt.Println("Example of Shadowing Package Names:")
fmt.Println("-----------------------------------")

pi := math.Pi

fmt.Println("Outside the block, math.Pi =", math.Pi)
if float64(x) > pi {
    math := "oops!"
    // This is an error: math.Pi is undefined because math == "oops!"
    // pi2 := math.Pi
    fmt.Println("\tInside the block, math =", math)
}
fmt.Println("Outside the block, math.Pi =", math.Pi)
```

## The Universe Block

- Go is a small language
  - Only 25 keywords
  - **The built-in types, constants, `nil`, and functions are not included in that list**
  - Go considers them *Predeclared Identifiers*
  - ***Predeclared Identifiers* are defined in the *Universe Block***
    - The root block that contains all other blocks
    - *These names can be shadowed in any other blocks*
- **WARNING: Never redefine any identifiers in the *Universe* Block**
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
} else if m > 5 {
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
  - But don't do that: Leave it for only declaring variables to be scoped for the `if-else` blocks
  - Anything else would be confusing
  - **Variable declared for the `if-else` block will shadow any existing outside variables**

## `for`

- **`for` is the only loop construct available in Go**
- **But there are 4 ways of using `for` loops**
  - *C-Style `for`*
  - *Condition-Only `for` (`while`-Style)*
  - *Infinite `for`*
  - *`for-range`*

### C-Style `for`

- Similar format to other C-based languages

```go
for init; condition; increment {
    body
}
```

- Similar to `if`, there are no `()` around `init; condition; increment`
- **`init`**
  - *Must use `:=` to initialize the variable*
  - `var` is not legal here
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
fmt.Println("Example of C-Style for-loop:")
fmt.Println("----------------------------")

for i := 0; i < 10; i++ {
    fmt.Printf("%d ", i)
}
fmt.Println()
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
```

### Infinite `for`

- If we remove all `init`, `condition`, and `increment`, we get an *infinite `for`*
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
  - **With *Infinite Loop* + `if` + `break`, we can similate a `do-while` loop**

```go
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
```

- **`continue` allows to skip an iteration**
  - Proceeds directly to the next iteration
  - `continue` can make codes much easier to reason and understand
  - It is also more idiomatic to use `continue`
  - Go encourages short `if` bodies
  - It can make the code easier to reason and understand

```go
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
```

### `for-range` Statement

- Allows to iterate over elements contained in some container types
- Similar to Python's `for x in range(n)`
- Can be used with *strings, arrays, slices, maps, channels*
  - **Can only be used on *built-in compound types***
  - Or user-defined types based on them

```go
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
```

- **With `for-range`, we get back 2 loop variables**
  - The position of the data in the container
  - The value at that position
- **The idiomatic names of these variables depend on what is being iterated over**
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
fmt.Println("Using for-range With Slices Without `i`:")
fmt.Println("----------------------------------------")

evenVals2 := []int{2, 4, 6, 8, 10, 12}
for _, v := range evenVals2 {
    fmt.Printf("%d ", v)
}
fmt.Println()
fmt.Println()
```

- **If we just want the key, we can ignore `v` completely**
  - Mostly used approach when a map is being used as a set
  - Can also be used for iteraring over arrays and slices but rare

```go
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
```

#### Iterating Over Maps

- **When iterating over Maps, the order is not guaranteed**
  - *The order of keys and values will vary for each run*
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

- We can also use `for-range` on strings
- `for-range` accesses the **runes** in a string in order
  - `i` - Number of bytes
  - `v` - The rune

```go
// Using for-range With Strings
// ----------------------------
greetings := []string{"Hello!", "Hi π!"}

// Iterating over the slice
fmt.Println("index\tchar\tstring(char)")
for _, greeting := range greetings {
    // Iterating over the string => runes
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
```

- **NOTE**
  - **Before Go 1.22, the variables are created once and reused**
  - **Since Go 1.22, the variables are re-created for each iteration**
  - This prevents some common bugs with goroutines
  - This is a backward-breaking change
  - Can enable behavior by specifying Go version in `go.mod`

### Labeling `for` Statements

- By default, `break` and `continue` applies to the closest `for`
- **For nested loops, we can manipulate which loop they would break or continue from using *labels***
- The label is always indented by `go fmt` to same-level as surrounding braces for the current block
  - Easier to notice

```go
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
  - Best to use with strings: *Loop over runes/chars instead of bytes*
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
  - **A standard `for`-loop *loops over bytes instead of runes/chars***
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
- *There is no need for `break` statements*
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
```

- **`break` could also be used in tandem with `switch` inside `for`**
  - Can be used to break out of the `for`-loop using label
  - Without a `for`-label, it might try to break out of the `case` instead

```go
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
```

- **NOTE: Go has a `fallthrough` keyword**
  - Allows a `case` to *continue* unto the next `case`
  - But this is not used often
  - If `fallthrough` is needed, it might be good to restructure the code
  - It is only used in very rare cases
- **We can also switch on any *type* that can be compared with `==`**
  - All built-in types except `slice`, `map`, `channel`, `func`, `struct`

### Blank `switch`

- A regular switch only allows a condition to be compared against equality to a value
- **We can use `switch` without specifying the value to compare**
- Allows to use any boolean comparison for each case instead
  - These are logical tests for each case
  - But if the boolean expressions are all `==`, a regalar `switch` is better
- We can also include a short variable declaration in the header of `switch`
  - The variable is scoped to the `switch` body

```go
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
```

### Choosing Between `if-else` and `switch`

- There is not much difference between `if-else` and blank `switch`
- `switch` indicates a relationship between the values in each cases
- This can help write clearer code

```go
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
```

- However, it is not idiomatic to do all sorts of unrelated comparisons in each `case`
  - **Each case should be related when using `switch`**
  - **If not, it is better to use `if-else`**
- **NOTE: Favor blank `switch` when cases are related**
  - Makes comparisons more visible
  - Reinforce the cases to be related set of concerns

## `goto`

- This is Go's 4th control structure
- But it is almost never used: Only in very special cases
- [*`goto` statements are considered harmful*](https://homepages.cwi.nl/~storm/teaching/reader/Dijkstra68.pdf)
  - It has been the *Black Sheep* of the programming community
  - Allows to jump anywhere in a program
  - Makes it very difficult to follow the logic of a program
- **Most modern language do not include `goto` at all**
  - In Go, it has some uses with limitations
- **In Go, `goto` specifies a labeled line of code**
  - Execution jumps to it
  - *However, we cannot jump just anywhere*
    - *Cannot skip over variable declarations*
    - *Cannot jump into an inner or parallel block*

```go
// Limitations of goto in Go
// -------------------------
func main() {
    a := 10
    goto skip // Cannot skip over variable declarations
    b := 20
skip:
    c := 30
    fmt.Println(a, b, c)
    if c > a {
        goto inner // Cannot jump into an inner or parallel block
    }
    if a < b {
    inner:
        fmt.Println("a is less than b")
    }
}
```

### Use of `goto` in Go

- Mostly, you should not use `goto`
- Labeled `break` and `continue` should be enough for loops
- Here is an example of a valid use of `goto`

```go
    // A Valid Reason To Use goto
    // --------------------------
    a := rand.Intn(10)
    for a < 100 {
        if a%5 == 0 {
            goto done
        }
        a = a*2 + 1
    }
    fmt.Println("Do something when the loop completes normally")

done:
    fmt.Println("Do something when the loop completes any differently than normal")
    fmt.Println(a)
```

- Allows to handle some logic that we do not want to run in the middle of a function
  - But we want to run at *some* ends of the function
- **Though we could also handle this without using `goto`**
  - Set up a boolean flag and use `if` to check the flag
  - Duplicate the code that runs after the loop
- **However**
  - Litering code with boolean flags is about the same as using `goto`
  - It is more verbose and can become confusing as well
  - Duplicating codes make it harder to maintain
  - In this case, using `goto` can improve understanding
  - *Example of real-world case:*
    - *`floatBits()` method in `strconv.atof.go` in standard library*
  - A better way would be to simply call a function
- **In general, try very hard to avoid using `goto`**
  - **In the rare situations where it makes code more readable, it is an option**
