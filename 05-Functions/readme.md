# Functions

---

- [Declaring and Calling Function](#declaring-and-calling-function)
- [Simulating Named and Optional Parameters](#simulating-named-and-optional-parameters)

---

## Declaring and Calling Function

- Go functions are similar to functions in other languages
- Similar with control structures, Go adds its own features on functions
  - Some improvements
  - Some experimental to avoid
- **Every Go program starts execution from the `main()` function**
  - `main()` does not take any parameters
  - `main()` does not return any values
- `fmt.Println()` allows to print to the console

```go
// Example of a Function
// ---------------------
func div(num int, denom int) int {
    if denom == 0 {
        return 0
    }
    return num / denom
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
func div(num, denom int) int {
    ...
}
```

- **Use `return` keyword to return a value from the function**
  - If a function returns a value, `return` is a must
  - If a function does not return a value, `return` can be removed
    - We do not provide the type of the return
    - `return` can still be used for exiting the function early
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
