# Functions

---

- [Declaring and Calling Function](#declaring-and-calling-function)
- [Simulating Named and Optional Parameters](#simulating-named-and-optional-parameters)
- [Variadic Function and Slices](#variadic-function-and-slices)
- [Multiple Return Values](#multiple-return-values)
  - [Multiple Return Values Are Multiple Values](#multiple-return-values-are-multiple-values)
  - [Ignoring Returned Values](#ignoring-returned-values)
  - [Named Return Values](#named-return-values)
    - [Inconveniences of Named Return Values](#inconveniences-of-named-return-values)
  - [Blank Return: Never Use Them](#blank-return-never-use-them)
  - [Functions Are Values](#functions-are-values)
  - [Function Type Declarations](#function-type-declarations)
  - [Anonymous Functions](#anonymous-functions)

---

## Declaring and Calling Function

- Go functions are similar to functions in other languages
- But similar with control structures, Go adds its own features on functions
  - Some are good improvements
  - Some are experimental to avoid
- **Every Go program starts execution from the `main()` function**
  - `main()` does not take any parameters
  - `main()` does not return any values
- `fmt.Println()` allows to print to the console

```go
// Example of a Function
// ---------------------
func div(num int, den int) int {
    if den == 0 {
        return 0
    }
    return num / den
}
```

- **A function declaration has 4 parts**
  - Keyword `func`
  - Function name
  - Input parameters: Param-name, Param-type
  - Return type
- **Go is a typed language**
  - Must specify the types of the parameters
  - The return types is specified between the input parameters and the function body
  - ***NOTE: When two or more consecutive input parameters are of the same type, we can specify the type once for all of them***

```go
func div(num, den int) int {
    ...
}
```

- **Use `return` keyword to return a value from the function**
  - If a function returns a value, `return` is a must
  - If a function does not return a value, `return` can be removed
    - We do not provide the type of the return
    - `return` can still be used for exiting the function early
    - But used by itself, without returning any value
- **The main function has no input parameters or return values**
  - When a function has no input parameters, use empty parentheses `()`

```go
// This is the main entry of the application.
func main() {
    // Example of calling a Function
    // -----------------------------
    result := div(100, 20)
    fmt.Println("div(100, 20) =", result)
}
```

## Simulating Named and Optional Parameters

- Go does not have *Named Parameters* and *Optional Parameters*
- **All function parameters must be supplied with arguments during call**
- We can emulate named and optional parameters using *struct*
  - Define fields that match the desired parameters
  - Pass the struct to the function

```go
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

func main() {
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
}
```

- **In practice, this is not a limitation**
  - A function should not have more than a few parameters
  - Named and optional parameters are mostly useful when function has a lot of parameters
  - If so, the function might be too complicated

## Variadic Function and Slices

- Note that `fmt.Println()` allows any number of parameters
- **Go supports *Variadic Parameters***
  - Must be the last or only parameter in the parameter list
  - Indicate with `...` before the parameter type
  - *The variable becomes a slice of the indicated type*
  - It can be used just like any other slice

```go
// Example of a Variadic Function
// ------------------------------
func addNums(nums ...int) int {
    var res int
    for _, n := range nums {
        res += n
    }
    return res
}

func main() {
    x := addNums(1, 2, 3, 4, 5, 6, 7, 8, 9)
    y := addNums(21, 43, 65)
    z := addNums()
    nums := []int{10, 20, 30, 40, 50}

    fmt.Println("addNums(1, 2, 3, 4, 5, 6, 7, 8, 9) =", x)
    fmt.Println("addNums(21, 43, 65) =", y)
    fmt.Println("addNums() =", z)
    fmt.Println("addNums(nums...) =", addNums(nums...))
}
```

- We can supply any value or no value at all
- We can also provide a slice, but we have to *spread* it using `...`
  - **Without the *spread operator*, it is compile-time error**

## Multiple Return Values

- Go allows multiple return values
- **The types of the return values are listed in `()`**
  - **All values must be returned**
  - Do not put `()` around the returned values, else compile-time error
- **We also return any associated errors**
  - **In Go, errors are values**
  - We use Go's multiple return value support to return an error
  - This is how error-handling is done in Go
  - **If the function completes succesfully, return `nil` as the error**
  - We check error by comparing it as `err != nil`
  - *By convention, the error is always the last value*

```go
import (
    "errors"
    "fmt"
    "os"
)

// Example of Function With Multiple Return Values
// -----------------------------------------------
func divmod(num, den int) (int, int, error) {
    if den == 0 {
        // Returning multiple values
        return 0, 0, errors.New("cannot divide by 0")
    }
    return num / den, num % den, nil
}

func main() {
    // Example of Function With Multiple Return Values
    // -----------------------------------------------
    resDiv, resMod, err := divmod(5, 2)
    // Error-handling
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    fmt.Println("divmod(5, 2) =>", "ResDiv =", resDiv, "ResMod =", resMod)
}
```

- Most of the time, we use `:=` to capture the multiple returned values
  - Allows to easily capture the values into separate variables

### Multiple Return Values Are Multiple Values

- It is not the same as returning a Tuple (like in Python)
  - In Python, we can *collect* the values in a single variable
  - Or we can assign them to multiple variables as well
  - **But not in Go**
- **In Go, we must assign each returned value individually**
  - **Assigning multiple return values to a single variable is a compile-time error**

```go
// Example of Multiple Return Values
// ---------------------------------
// The following is an error
results := divmod(34, 5)

// The following is correct
resDiv, resMod, err := divmod(34, 5)
```

### Ignoring Returned Values

- Go does not allow *unused variables*
- **If we do not want to use some of the return values, assign them to `_`**

```go
// Example of Ignoring Returned Value
// ----------------------------------
_, resMod, err := divmod(34, 5)
```

- *NOTE: It is also possible to ignore all the return values*
  - We simply do not assign to any variable
  - This is similar to calling `fmt.Println()`
  - **`fmt.Println()` returns 2 values but it is idiomatic to ignore them**
  - In all other case, explicitly ignore return values using `_`

```go
// Example of Ignoring All Returned Values
// ---------------------------------------
divmod(34, 5)
fmt.Println("Hi")
```

### Named Return Values

- Go also allows to specify names for the return values
- **If we supply names to the return values, they become predeclared variables**
  - We can use them in the function body to hold the final return values
  - **The returned names must be surrounded by `()`, even if only one**
  - *The names are initialized to their types' zero values*
  - They can be returned before any explicit use or assignment
- We can also name only some of the return values
  - Use `_` as the name of any that should be nameless
- **The names for the returned values are local to the function**
  - Does not enforce any names outside the function
  - We can assign them to any names outside the function

```go
// Example of Function With Named Return Values
// --------------------------------------------
func divmodNamed(num, den int) (res int, mod int, err error) {
    if den == 0 {
        err = errors.New("cannot divide by 0")
        // Returning multiple values
        return res, mod, err
    }
    res, mod = num / den, num % den
    return res, mod, nil
}

func main() {
    // Example of Function With Named Return Values
    // --------------------------------------------
    resX, modY, errZ := divmodNamed(5, 2)
    // Error-handling
    if errZ != nil {
        fmt.Println(errZ)
        os.Exit(1)
    }
    fmt.Println("divmodNamed(5, 2) =>", "resX =", resX, "modY =", modY)
}
```

#### Inconveniences of Named Return Values

- Shadowing
  - It is possible to shadow with the names
  - Ensure that outside variables are not shadowed by the local variables
- They are not required to be returned
  - We could return different values instead
  - The compiler will not complain
  - **The values of the `return` statement will always be returned**
  - This can create confusion
  - The Go compiler insert code that assigns any return value to the names
  - The named return parameters declares the *intent*, but are *not required* to use them

```go
func divmodNamedUnreturned(num, den int) (res int, mod int, err error) {
    res, mod = 20, 50
    if den == 0 {
        err = errors.New("cannot divide by 0")
        // Actual return values
        return 0, 0, err
    }
    // Actual return values
    return num / den, num % den, nil
}
```

- Named returns could be helpful for documentation
- But they do not add additional values
- They might create more issues instead
- But they are good in one situation

### Blank Return: Never Use Them

- With named return values, it is possible to just write `return` without specifying the values to return
  - Returns the last value assigned to the named return values
  - If returning the zero-values, be sure they make sense
  - `return` keyword is still required, else compile-time error
- Blank return might appear handy
  - But it is a bad idea because it makes is harder to understand the data flow
  - **Never use blank return: Always specify what is being returend**

### Functions Are Values

- Functions in Go are values and can be passed around
  - **Type Signature: `func` + Param Types + Return Type**
  - Function can have the same type signature
- **Since they are values, we can declare a function variable**
  - Can be assigned any function that match the type signature
  - Default Zero value is `nil`
  - Attempting to run a function variable `nil` is a `panic`

```go
// Declaring a Function Variable
// -----------------------------

func main() {
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
}

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
```

- Functions as values are very helpful in different scenarios

```go
// Example of a Simple Calculator With Functions
// ---------------------------------------------
func add(i int, j int) int { return i + j }
func sub(i int, j int) int { return i - j }
func mul(i int, j int) int { return i * j }
func div(i int, j int) int { return i / j }

func main() {
    opMap := map[string]func(int, int) int {
        "+": add,
        "-": sub,
        "*": mul,
        "/": div,
    }
    expressions := [][]string {
        {"2","+","3"},
        {"2","-","3"},
        {"2","*","3"},
        {"2","/","3"},
        {"2","%","3"},
        {"two","+","three"},
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
}
```

- `strconv.Atoi()` allows to convert a string to int
  - Also return an `error` if the conversion fails
- The type of `opFunc` is `func (int, int) int`
- **Calling a function in a variable is the same as calling a regular function**
- **NOTE: Do not write fragile programs**
  - Make sure to check for special-cases and handle appropriately
  - It might be tempting to not validate incoming data
  - Doing so produce unstable and unmaintainable code
  - Good error handling is what separate the pros from the amateurs

### Function Type Declarations

- We can use the `type` keyword to define a function type
- **Any function that match the signature automatically meets the type**

```go
// Declaring a Function Type
type opFuncType func(int, int) int

// Using a Function Type
opMap := map[string]opFuncType {
    // ...
}
```

- Advantages of declaring a function type
  - Documentation
  - Less repetition
  - Easier to maintain

### Anonymous Functions
