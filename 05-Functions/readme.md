# Functions

---

- [Declaring and Calling Function](#declaring-and-calling-function)
- [Simulating Named and Optional Parameters](#simulating-named-and-optional-parameters)
- [Variadic Function and Slices](#variadic-function-and-slices)
- [Multiple Return Values](#multiple-return-values)
  - [Multiple Return Values are Multiple Values](#multiple-return-values-are-multiple-values)

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
- **Go is a typed languages**
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

- Go does not have *Name Parameters* and *Optional Parameters*
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

- We can supply any value or no values at all
- We can also provider a slice, but we have to *spread* it using `...`
  - **Without the *spread operator*, it is compile-time error**

## Multiple Return Values

- Go allows multiple return values
- **The types of the return values are listed in `()`**
  - **All values must be returned**
  - Do not put `()` around the returned values, esle compile-time error
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

### Multiple Return Values are Multiple Values
