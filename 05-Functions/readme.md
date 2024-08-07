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
- [Closures](#closures)
  - [Benefits of Closures](#benefits-of-closures)
- [Passing Functions As Arguments](#passing-functions-as-arguments)
- [Returning Functions From Functions](#returning-functions-from-functions)
- [`defer`](#defer)
- [Go Is Call By Value](#go-is-call-by-value)

---

## Declaring and Calling Function

- Go functions are similar to functions in other languages
- But Go adds its own features on functions
  - Some are good improvements
  - Some are experimental to avoid
- **Every Go program starts execution from the `main()` function**
  - `main()` does not take any parameters
  - `main()` does not return any values
- `fmt.Println()` allows to print to the console

```go
// Example of a Function
// ---------------------

// Divide an integer by another integer.
func div(num int, den int) int {
    if den == 0 {
        return 0
    }
    return num / den
}
```

- **A function declaration has 5 parts**
  - Keyword `func`
  - Function name
  - Input parameters: Param-name + Param-type
  - Return type
  - Body

```go
func div(num int, den int) int {
    // Define the body here
}
```

- **NOTE: Go is a typed language**
  - Must specify the types of the parameters
  - The return type is specified between the input parameters and the function body
  - ***NOTE: When two or more consecutive input parameters are of the same type, we can specify the type once for all of them***

```go
func div(num, den int) int {
    // Define the body here
}
```

- **Use `return` keyword to return a value from the function**
  - If a function returns a value, `return` is a must
  - If a function does not return a value, `return` can be removed
    - We do not provide the type of the return
    - `return` can still be used for exiting the function early
    - But must be used by itself, without returning any value
- **The main function has no input parameters or return values**
  - When a function has no input parameters, use empty parentheses `()`

```go
// This is the main entry of the application.
func main() {
    // Example of calling a Function
    // -----------------------------
    fmt.Println("Example of calling a Function:")
    fmt.Println("------------------------------")

    result := div(100, 20)

    fmt.Println("div(100, 20) =", result)
}
```

## Simulating Named and Optional Parameters

- Go does not have *Named Parameters* and *Optional Parameters*
- **All function parameters must be supplied with arguments during call**
- We can simulate named and optional parameters using *struct*
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

// A test function for Optional and Named Parameters.
func MyFunc(params FuncParams) error {
    // Do something
    fmt.Println("Passed parameters:", params)
    return nil
}

func main() {
    // Example of a Function With Optional and Named Parameters
    // --------------------------------------------------------
    fmt.Println("Example of a Function With Optional and Named Parameters:")
    fmt.Println("---------------------------------------------------------")

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
  - If so, the function might be too complicated and should be simplified

## Variadic Function and Slices

- Note that `fmt.Println()` allows any number of parameters
- **Go supports *Variadic Parameters***
  - Must be the last or only parameter in the parameter list
  - Indicate with `...` before the parameter type
  - *The variable becomes a slice of the indicated type*
  - It can be used just like any other slice
- We can supply any argument value or no value at all
  - We can also provide a slice, but we have to *spread* it using `sliceName...` for the argument
  - **Without the *spread operator*, it is compile-time error**

```go
// Example of a Variadic Function
// ------------------------------

// Add any number of integers and return their sum.
func addNums(nums ...int) int {
    var res int
    for _, n := range nums {
        res += n
    }
    return res
}

func main() {
    // Example of a Variadic Function
    // ------------------------------
    fmt.Println("Example of a Variadic Function:")
    fmt.Println("-------------------------------")

    x := addNums(1, 2, 3, 4, 5, 6, 7, 8, 9)
    y := addNums(21, 43, 65)
    z := addNums()
    nums := []int{10, 20, 30, 40, 50}

    fmt.Println("addNums(1, 2, 3, 4, 5, 6, 7, 8, 9) =", x)
    fmt.Println("addNums(21, 43, 65) =", y)
    fmt.Println("addNums() =", z)
    fmt.Println("nums =", nums)
    fmt.Println("addNums(nums...) =", addNums(nums...))
}
```

## Multiple Return Values

- Go allows multiple return values
- **The types of the return values are listed in `()`**
  - **All listed values must be returned**
  - Do not put `()` around the actual returned values, else compile-time error
- **We also return any associated errors**
  - **In Go, errors are values**
  - We use Go's multiple-return-value-support to return an error
    - This is how error-handling is done in Go
    - **If the function completes succesfully, return `nil` as the error**
    - We check error by comparing it as `err != nil`
    - *By convention, the error is always the last value*
- Most of the time, we use `:=` to capture the multiple returned values
  - Allows to easily capture the values into separate variables

```go
import (
    "errors"
    "fmt"
    "os"
)

// Example of Function With Multiple Return Values
// -----------------------------------------------

// Divide an integer by anoter integer and return the result and the mod.
func divmod(num, den int) (int, int, error) {
    if den == 0 {
        return 0, 0, errors.New("cannot divide by 0")
    }
    return num / den, num % den, nil
}

func main() {
    // Example of Function With Multiple Return Values
    // -----------------------------------------------
    fmt.Println("Example of Function With Multiple Return Values:")
    fmt.Println("------------------------------------------------")

    resDiv, resMod, err := divmod(5, 2)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    fmt.Println("divmod(5, 2) =>", "ResDiv =", resDiv, "ResMod =", resMod)
}
```

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

- **NOTE: It is also possible to ignore all the return values**
  - We simply do not assign to any variable
  - This is similar to calling `fmt.Println()`
    - **`fmt.Println()` normally returns 2 values**
    - **But it is idiomatic to ignore them**
  - Otherwise, explicitly ignore return values using `_`

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
  - **The declared returned names must be surrounded by `()`, even if only one**
  - *The names are initialized to their types' zero values*
  - They can be returned before any explicit use or assignment
- **We can also name only some of the return values**
  - Use `_` as the name of any that should be nameless
- **The names for the returned values are local to the function**
  - Does not enforce any names outside the function
  - We can assign them to any names outside the function

```go
import (
    "errors"
    "fmt"
    "os"
)

// Example of Function With Named Return Values
// --------------------------------------------

// Divide an integer by another integer and return the result and the mod.
func divmodNamed(num, den int) (res int, mod int, err error) {
    if den == 0 {
        err = errors.New("cannot divide by 0")
        // Returning multiple values: Default to zero-values
        return res, mod, err
    }
    res, mod = num/den, num%den
    return res, mod, nil
}

func main() {
    // Example of Function With Named Return Values
    // --------------------------------------------
    fmt.Println("Example of Function With Named Return Values:")
    fmt.Println("---------------------------------------------")

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
  - **The values in the `return` statement will always be returned**
  - This can create confusion
  - The Go compiler insert code that assigns any return value to the names
  - **The named return parameters declares the *intent*, but are *not required* to use them**

```go
// This version will work fine with no compiler errors
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

- **With named return values, it is possible to just write `return` without specifying the values to return**
  - Returns whichever last value assigned to the named return values
  - If returning the zero-values, be sure they make sense
  - **`return` keyword is still required, else compile-time error**
- Blank return might appear handy
  - But it is a bad idea
  - It makes it harder to understand the data flow
  - **Never use blank return**
  - **Always specify what is being returned**

## Functions Are Values

- Functions in Go are values and can be passed around
  - **Type Signature: `func(ParamTypes) ReturnType`**
  - Functions can have the same type signature
- **Since they are values, we can declare a function variable**
  - Can be assigned any function that match the type signature
  - Default Zero value is `nil`
  - Attempting to run a function variable `nil` is a `panic`

```go
// Example of Declaring a Function Variable
// ----------------------------------------

// Get the length of a string.
func f1(s string) int {
    return len(s)
}

// Get the sum of runes in a string.
func f2(s string) int {
    sum := 0
    for _, c := range s {
        sum += int(c)
    }
    return sum
}

func main() {
    // Example of Declaring a Function Variable
    // ----------------------------------------
    fmt.Println("Example of Declaring a Function Variable:")
    fmt.Println("-----------------------------------------")

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
```

- Functions as values are very helpful in different scenarios
  - Here, the type of `opFunc` is `func(int, int) int`

```go
// Example of a Simple Calculator With Functions
// ---------------------------------------------

func add(i int, j int) int { return i + j }
func sub(i int, j int) int { return i - j }
func mul(i int, j int) int { return i * j }
func div(i int, j int) int { return i / j }

func main() {
    // Example of a Simple Calculator With Functions
    // ---------------------------------------------
    fmt.Println("Example of a Simple Calculator With Functions:")
    fmt.Println("----------------------------------------------")

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

- **`strconv.Atoi()` allows to convert a string to int**
  - Returns an `error` if the conversion fails
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
type opFunc func(int, int) int

// Using a Function Type
opMap := map[string]opFunc {
    // ...
}
```

- Advantages of declaring a function type
  - Documentation
  - Less repetition
  - Easier to read
  - Easier to maintain

### Anonymous Functions

- **We can define functions and assign them to variables**
- We declare an anonymous function with `func(ParamTypes) ReturnType`
  - Followed immediately by the input parameters, return values, and body
  - **It is a compile-time error to try to add a name**
- It is called like any other functions

```go
// Example of Anonymous Function
// -----------------------------
fmt.Println("Example of Anonymous Function:")
fmt.Println("------------------------------")

// Anonymous function assigned to a variable
anonFunc := func(j int) {
    fmt.Println("Printing", j, "from inside an anonymous function")
}
for i := range 5 {
    // Calling the function
    anonFunc(i)
}
```

- **We do not even have to assign an anonymous function to a variable**
  - Can be written inline and called immediately
  - But not something that we would normally do
  - If we are doing this approach, might as well just call the code
  - **But this is typically used with `defer` and `go` routines**

```go
// Example of Inline Anonymous Function
// ------------------------------------
fmt.Println("Example of Inline Anonymous Function:")
fmt.Println("-------------------------------------")

// Inline anonymous function
for i := range 5 {
    func(j int) {
        fmt.Println("Printing", j, "from inside an inline anonymous function")
    }(i)
}
```

- **NOTE: We can also declare package-scoped variables that are assigned anonymous functions**

```go
// Examples of Package-Scoped Variables With Anonymous Functions
// -------------------------------------------------------------

var (
    add = func(i int, j int) int { return i + j }
    sub = func(i int, j int) int { return i - j }
    mul = func(i int, j int) int { return i * j }
    div = func(i int, j int) int { return i / j }
)

func main() {
    x := add(2, 3)
    fmt.Println(x)
}
```

- **NOTE: We can also re-assign the value of a Package-level anonymous function**

```go
// Examples of Re-assigning Package-Scoped Variables With Anonymous Functions
// --------------------------------------------------------------------------

var (
    add = func(i int, j int) int { return i + j }
    sub = func(i int, j int) int { return i - j }
    mul = func(i int, j int) int { return i * j }
    div = func(i int, j int) int { return i / j }
)

func main() {
    x := add(2, 3)
    fmt.Println(x)
    changeAdd()
    y := add(2, 3)
    fmt.Println(y)
}

// Reassigns add to a different function than declared earlier.
func changeAdd() {
    // `add` here refers to the package-level variable
    add = func(i int, j int) int { return i + j + j }
}
```

- **NOTE: Be sure you need this capability before using Package-level anonymous function**
  - Package-level states should be left immutable to make data-flow easier to understand
  - If a function changes while running, the data-flow can become very confusing

## Closures

- Functions declared inside functions
- They are able to access and modify variables declared in the outer function
- **We can read/write from/to the outside variables even though they are not passed explicitly**

```go
// Example of Closure
// ------------------
fmt.Println("Example of Closure:")
fmt.Println("-------------------")

a := 20
fmt.Println("Outside fa(), a =", a)
/*fa =*/ func() {
    fmt.Println("\tInside fa() before assignment, a =", a)
    a = 30 // The assignment modifies outside a
    fmt.Println("\tInside fa() after assignment, a =", a)
}()
fmt.Println("Outside fa(), a =", a)
```

- We can also shadow the variables if we assign with `:=`
- **Be careful to use the correct assignment operator when working with closures**

```go
// Example of Closure With Shadow
// ------------------------------
fmt.Println("Example of Closure With Shadow:")
fmt.Println("-------------------------------")

b := 20
fmt.Println("Outside fb(), b =", b)
fb := func() {
    fmt.Println("\tInside fb() before assignment, b =", b)
    b := 30 // This assignment shadows instead of modifying outside b
    fmt.Println("\tInside fb() after assignment, b =", b)
}
fb()
fmt.Println("Outside fb(), b =", b)
```

### Benefits of Closures

- **Allows to limit a function's scope**
  - We can *hide* the function if it is only local to the enclosing function
- **Reduces the number of declarations at the Package-level**
  - Can be used to make repeating logic more DRY inside of a function
- **Makes some very useful patterns when passed around**
  - Allows to take variables within a function and use the values outside the function

## Passing Functions As Arguments

- **Functions are values**
  - **We can pass them as arguments to other functions**
  - *Treat the function as data*
  - **We can pass *Closures* around: This is a very useful pattern**
  - An example is *Sorting Slice* using `sort.Slice(slc, func)`
    - **NOTE: This predates generics in Go**

```go
// Example of Sorting Slice
// ------------------------
fmt.Println("Example of Sorting Slice")
fmt.Println("------------------------")

type Person struct {
    FirstName string
    LastName  string
    Age       int
}
people := []Person{
    {"John", "Smith", 37},
    {"Jeremy", "Trye", 18},
    {"Jasmine", "Alter", 20},
}

fmt.Println("Before Sorting:", people)

// Sorting the slice by last name
sort.Slice(people, func(i int, j int) bool {
    return people[i].LastName < people[j].LastName
})
fmt.Println("After Sorting By Last Name:", people)

// Sorting the slice by age
sort.Slice(people, func(i int, j int) bool {
    return people[i].Age < people[j].Age
})
fmt.Println("After Sorting By Age:", people)
```

- A closure anonymous function is passed to `sort.Slice()`
  - `people` exists in the context of `sort.Slice()`
  - The closure also exists in the context of `sort.Slice()`
  - The outside variable `people` is *captured* within the closure
- **Passing functions as arguments to other functions is useful for performing different operations on the same kind of data**

## Returning Functions From Functions

- We can also return a closure from a function

```go
// Example of Function That Returns a Closure
// ------------------------------------------

// Return a function that multiplies a number to a given base.
func makeMult(base int) func(int) int {
    return func(n int) int {
        return n * base
    }
}

func main() {
    // Example of Function That Returns a Closure
    // ------------------------------------------
    fmt.Println("Example of Function That Returns a Closure:")
    fmt.Println("-------------------------------------------")

    base2Mult := makeMult(2)
    base3Mult := makeMult(3)
    base5Mult := makeMult(5)

    fmt.Println("i\tbase2\tbase3\tbase5")
    fmt.Println("-\t-----\t-----\t-----")
    for i := range 10 {
        fmt.Println(i, "\t", base2Mult(i), "\t", base3Mult(i), "\t", base5Mult(i))
    }
}
```

- **Closures are very helpful in Go**
  - Can be used to sort slices
  - Can be used to efficiently search a sorted slice with `sort.Search()`
  - Returning closures is used when bulding middleware for web server
  - Also used to implement resource-cleanups with `defer`
  - Typical pattern for *High-Order Functions* and *Factory Functions*

## `defer`

- **Progams often create temporary resources that need to be cleaned up later**
  - The cleanup has to happen no matter the exit point of the function
  - We can think of this as the `finally` portion of other languages' error handlers
  - But applied specifically to function calls
- **In Go, the cleanup code is attached to the function using `defer`**
  - We use `defer` to release resources
  - Normally, function calls are called right away
  - **`defer` delays the invocation until the enclosing function exits**
  - *Executed when either the surrounding function returns, reach the end, or panic*

```go
func main() {
    // Example of defer With a cat Command
    // -----------------------------------
    fmt.Println("Example of defer With a cat Command:")
    fmt.Println("------------------------------------")

    // Make sure a filename was passed as argument: make try ARGS="<filename>"
    // Args[0] is the name of the program
    if len(os.Args) < 2 {
        log.Fatal("Error: No file was specified")
    }
    // Open the file: Read-only
    fl, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
    // Close the file after using it
    // This must be run no matter any errors in the program
    defer func() {
        fmt.Println("Defer in main() is called here")
        fmt.Println("This simulates closing a file that was opened in main()")
        fmt.Println()
        fl.Close()
    }()
    // Read from the file
    data := make([]byte, 2048)
    for {
        count, err := fl.Read(data)
        os.Stdout.Write(data[:count])
        if err != nil {
            if err != io.EOF {
                log.Fatal(err)
            }
            break
        }
    }
}
```

- **NOTES**
  - We read from a file by passing a slice of bytes to `fl.Read()`
  - It returns the number of bytes that were read into the slice and error
  - Also need to handle `EOF` marker to stop reading the file
- With `defer`
  - We can use a function, method, or closure
  - We can defer multiple functions in a Go function
    - **They run in a *Last-In, First-Out* (LIFO) order (Stack)**
    - The last `defer` registered will run first
  - **Code within `defer` functions runs *after the `return` statement of the enclosing function***
    - We can supply a function with input parameters to `defer`
    - **The input parameters are evaluated immediately at the declaration location**
    - Their values are stored until the function runs

```go
// Example of Using defer
// ----------------------
func main() {
    // Example of Using defer In a Function
    // ------------------------------------
    fmt.Println("Example of Using defer In a Function:")
    fmt.Println("-------------------------------------")

    deferExample()
}

// A test function for using defer.
func deferExample() int {
    a := 10
    defer func(val int) {
        fmt.Println("First value:", val)
    }(a)
    a = 20
    defer func(val int) {
        fmt.Println("Second value:", val)
    }(a)
    a = 30
    fmt.Println("Exiting deferExample:", a)
    return a
}
```

- We could also supply a function that returns values to `defer`
  - But there is no way to read those values

```go
func example() {
    defer func() int {
        return 2 // We cannot access/read this value
    }()
}
```

- **A deferred function can access/modify the return value of its enclosing functions**
  - It is the best reason to use *Named Return Values*
  - Allow code to take action based on an error
  - `defer` can add contextual information to an error returned from a function
  - It is essentially a closure

```go
// Example of handling DB transaction cleanups Using defer
// -------------------------------------------------------
func InsertIntoTable(ctx context.Context, db *sql.DB, val1, val2 string) (err error) {
    tx, err := db.BeginTx(ctx, nil)
    if err != nil {
        return err
    }
    defer func() {
        // We can access outside err from here
        // This is essentially a closure
        if err != nil {
            tx.Rollback()
        } else {
            err = tx.Commit()
        }
    }()
    _, err = tx.ExecContext(ctx, "INSERT INTO table (col) VALUES $1", val1)
    if err != nil {
        return err
    }
    // Use tx to do more DB inserts here
    return nil
}
```

- **NOTE: The standard library includes good support for databases**
  - `database/sql`
- A function that allocates a resource should also return a closure to cleanup that resource

```go
// Example of Resouce Cleanup with defer
// -------------------------------------
func main() {
    if len(os.Args) < 2 {
        log.Fatal("no file specified")
    }
    // Get the file
    f, closer, err := getFile(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
    defer closer()
    data := make([]byte, 2048)
    for {
        count, err := f.Read(data)
        os.Stdout.Write(data[:count])
        if err != nil {
            if err != io.EOF {
                log.Fatal(err)
            }
            break
        }
    }
}

// Open a file and returns a closure
func getFile(name string) (*os.File, func(), error) {
    file, err := os.Open(name)
    if err != nil {
        return nil, nil, err
    }
    return file, func() {
        file.Close()
    }, err
}
```

- **NOTE: Using `defer` can feel strange than using `try-catch-finally`**
  - But `try-catch-finally` creates an additional indentation
  - Makes the code harder to read

## Go Is Call By Value

- **When a function is called with parameters, Go makes a copy of the passed parameters**
- **Modifying parameters inside a function body does not affect the passed arguments**
  - This is not just for primitive types
  - Also applies to `struct`

```go
// Example of Call-By-Value
// ------------------------
type Person2 struct {
    age  int
    name string
}

func main() {
    // Example of Call-By-Value
    // ------------------------
    fmt.Println("Example of Call-By-Value:")
    fmt.Println("-------------------------")

    i := 2
    s := "Hello"
    p := Person2{}

    fmt.Println("Before function call:", i, s, p)

    // Modifying the passed parameters has no effect on the actual arguments
    func(n int, st string, per Person2) {
        // Attempting to modify the passed parameters
        n = n * 2
        st = "Goodbye"
    }(i, s, p)

    fmt.Println("After function call:", i, s, p)
}
```

- **NOTE: This behavior is different for maps and slices**
  - For *Maps*: Any changes made to a map parameter is reflected on the original map argument
  - For *Slice*: We can modify the slice elements **but cannot lengthen the slice**
  - **Applies for both maps and slices arguments or fields of structs**

```go
// Example of Map-Slice Modification Calls
// ---------------------------------------
fmt.Println("Example of Map-Slice Modification Calls:")
fmt.Println("----------------------------------------")

mapMod := map[int]string{
    1: "first",
    2: "second",
}

fmt.Println("Before: mapMod =", mapMod)
func(m map[int]string) {
    m[2] = "hello"
    m[3] = "goodbye"
    delete(m, 1)
}(mapMod)
fmt.Println("After: mapMod =", mapMod)

slcMod := []int{1, 2, 3}

fmt.Println("Before: slcMod =", slcMod)
func(s []int) {
    for k, v := range s {
        s[k] = v * 200
    }
    s = append(s, 10)
}(slcMod)
fmt.Println("After: slcMod =", slcMod)
```

- **Why do Maps and Slices behave differently?**
  - Because they are both implemented with pointers
  - **Every type in Go is always a *Value-Type***
  - **But sometimes, that value is a *Pointer***
- *Pass-By-Value* is why Go's limited support for Constant is not a huge issue
  - **Calling a function does not modify the original variable**
  - **Unless it is a slice or a map**
  - In general, this is a good thing: Easier flow of data
- **However, sometimes we do need to modify the original value**
  - In those cases, we use *Pointers*
