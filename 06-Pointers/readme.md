# Pointers

---

- [Quick Pointer Primer](#quick-pointer-primer)
- [Pointers Behavior Like Classes](#pointers-behavior-like-classes)
- [Pointers Indicate Mutable Parameters](#pointers-indicate-mutable-parameters)
- [Pointer: Last Resort](#pointer-last-resort)
- [Pointer-Passing Performance](#pointer-passing-performance)
- [Zero-Value vs No Value](#zero-value-vs-no-value)
- [Map vs Slice](#map-vs-slice)
  - [Slices As Buffers](#slices-as-buffers)
  - [Reducing the GC's Workload](#reducing-the-gcs-workload)

---

## Quick Pointer Primer

- **Pointer**
  - A variable that holds the memory address of where in the memory an associated value is stored
- **Variables are stored in one or more contiguous memory addresses**
  - Different variable types can take different number of memory addresses
  - Depends on the size of the data type
  - **The smallest amount of addressable memory is a byte (8-bits)**

```go
var x int32 = 10    // 4 bytes
var y bool = true   // 1 byte
```

<img src="./img/2-variables-in-memory.png" width=30%>

- **A pointer is a variable that contains the memory address of another variable**
  - Holds a number that indicates the memory location where the data is stored
  - This number is the *Memory Address*

```go
// Pointer Operators
// -----------------
fmt.Println("Pointer Operators:")
fmt.Println("------------------")

var x int32 = 10  // Value-type int32
var y bool = true // Value-type bool

var ptrX *int32 = &x // Pointer-type to a value of type int32
ptrY := &y           // Pointer-type to a value of type bool

fmt.Println("ptrX =", ptrX)   // Prints the memory address
fmt.Println("*ptrX =", *ptrX) // Prints the pointed value: Same as x
fmt.Println("x =", x)
fmt.Println("ptrY =", ptrY)   // Prints the memory address
fmt.Println("*ptrY =", *ptrY) // Prints the pointed value: Same as y
fmt.Println("y =", y)
```

<img src="./img//pointers-in-memory.png" width=70%>

- **Every pointer is always occupying the same (fixed) number of memory locations**
  - Regardless of the type it is pointing to
  - In the example, we are using 4 bytes memory-address length
  - **In modern computer, it is usually 8 bytes memory-address length**
  - In Go, the size of a pointer variable is:
    - 8-bytes for 64-bit machines
    - 4-bytes for 32-bit machines
- ***Zero-Value of a Pointer: `nil`***
  - `nil` is untyped identifier
  - Represents lack of value
  - Not another name for `0` unlike in C
  - Cannot convert back and forth with a number
  - Defined in the *Universe* block
    - *Can be shadowed*
    - *Never name variables `nil`*
- **WARNING: Before dereferencing a pointer, make sure that it is not `nil`**
  - Attempting to dereference a `nil` pointer results in a panic

```go
// Example of nil Pointer
// ----------------------
fmt.Println("Example of nil Pointer:")
fmt.Println("-----------------------")

var ptrZ *string // Pointer-type to a value of type string but default to nil

fmt.Println("ptrZ =", ptrZ) // Prints nil

// Attempting to dereference a nil pointer results in a panic
// Invalid Memory Address & Segmentation Violation
// fmt.Println("*ptrZ =", *ptrZ) // panic: runtime error: invalid memory address or nil pointer dereference
```

- ***Slice*, *Map*, and *Function* are implemented using Pointers**
  - Which is why their zero-values are the same: `nil`
  - *Channel* and *Interface* are also implemented using Pointers
- **Go's pointer syntax is partially borrowed from C/C++**
  - But without painful memory management
  - Go is a *Garbage-Collected* language
  - Some pointer features in C/C++ are not allowed (E.g. *Pointer Arithmetics*)
- *NOTE: Go has [`unsafe`](https://pkg.go.dev/unsafe) package for low-level operations*
  - But it is exceedingly rare to use `unsafe`
- **`&` - *Address Operator***
  - Precedes a *value-type* variable
  - Returns the memory address where the value of that variable is stored
  - **This is called *Referencing* the variable**
- **`*` - *Indirection Operator***
  - Precedes a *pointer-type* variable
  - Returns the pointed value at that memory address
  - **This is called *Dereferencing* the pointer**
  - **However, when used on a *type* instead of a variable, it denotes a *Pointer-Type* to that type**
- **WARNING: Before dereferencing a pointer, make sure that it is not `nil`**
  - Attempting to dereference a `nil` pointer results in a panic

```go
// Pointer Operators
// -----------------
var x int32 = 10        // Value-type int32
var ptrX *int32 = &x    // Pointer-type to a value of type int32: Referencing

fmt.Println("ptrX =", ptrX)     // Prints the memory address: 0xc000012128
fmt.Println("*ptrX =", *ptrX)   // Dereferencing: Prints the pointed value: 10

var nilPtr *string      // Pointer-type to a value of type string but default to nil

fmt.Println("nilPtr =", nilPtr) // Prints nil
fmt.Println(nilPtr == nil)      // Prints true
// fmt.Println("*nilPtr =", *nilPtr) // panic: runtime error: invalid memory address or nil pointer dereference
```

- **Pointer Type**
  - A type that represents a pointer
  - Written with a `*` before the type name (E.g. `*int`)
  - Can be based on any type

```go
// Example of Pointer Type
// -----------------------
fmt.Println("Example of Pointer Type:")
fmt.Println("------------------------")

intVal := 10
var ptrIntVal *int
ptrIntVal = &intVal

fmt.Println("intVal =", intVal)
fmt.Println("ptrIntVal =", ptrIntVal)
```

- **Built-in function `new()` creates a pointer variable**
  - Returns **a pointer to a zero-value instance of the type**
  - Allows to *not set the pointer to `nil`*
  - But `new()` is rarely used

```go
// Example of Using new()
// ----------------------
fmt.Println("Example of Using new():")
fmt.Println("-----------------------")

ptrNewVar := new(int)                              // Returns a pointer to 0
fmt.Println("ptrNewVar == nil:", ptrNewVar == nil) // false
fmt.Println("*ptrNewVar =", *ptrNewVar)            // 0
```

- For Struct, use `&` before the Struct literal
- **Cannot use `&` on primitive literals or constants**
  - They do not have memory address
  - Exist only at compile time
  - *If pointer is needed for them, declare a variable instead*

```go
x := &Foo{} // struct pointer
var y string
var z int
ptrY := &y  // String pointer
ptrZ := &z  // Integer pointer
```

- **Not being able to get the address of a constant is sometimes inconvenient**
  - Cannot assign literals directly to pointer-type fields

```go
// Unable to Get Address of Constants
// ----------------------------------

type Person struct {
    FirstName   string
    MiddleName  *string
    LastName    string
}

p := Person{
    FirstName: "John",
    MiddleName: "Edler", // cannot use "Edler" (type string) as type *string in field value.
    LastName: "Smith",
}

p := Person{
    FirstName: "John",
    MiddleName: &"Edler", // cannot take the address of "Edler".
    LastName: "Smith",
}
```

- 2 ways to solve this:
  - *1. Introduce a variable to hold the constant value*
  - *2. Write a generic helper function: Takes a param of any type and return a pointer to that type*
- With the generic approach
  - **The constant is copied to the generic function as variable (param)**
  - Variables have memory address

```go
// Generic Pointer Helper For Constants
// ------------------------------------

// Generic helper function for getting constant's pointer.
func makeConstPtr[T any](t T) *T {
    return &t
}

func main() {
    // Generic Pointer Helper For Constants
    // ------------------------------------
    fmt.Println("Generic Pointer Helper For Constants:")
    fmt.Println("-------------------------------------")

    type Person struct {
        FirstName   string
        MiddleName  *string
        LastName    string
    }

    p := Person{
        FirstName: "John",
        MiddleName: makeConstPtr("Edler"), // This works!
        LastName: "Smith",
    }

    fmt.Println("p =", p)
}
```

## Pointers Behavior Like Classes

- Pointers might look intimidating
- **But Pointers are actually the familiar behavior for classes in other languages**
- **In other languages, there is a behavior difference between primitives and classes**
  - When primitives are aliased or passed to functions, changes made to the other variable/parameter are not reflected
  - The aliases (params vs args) do not share the same memory
  - **They are often referred to as *Passed-By-Value* in Java and JavaScript**
  - **Python and Ruby use *Immutable Instances* for the same purpose**

```py
# Python As An Example: Immutable Instance
# ----------------------------------------
x = 10
y = x # Attempt aliasing
y = 20
print(x) # Prints 10


def attempt_change(a):
    a = 1000


attempt_change(x)
print(x) # Prints 10
```

- **This is not the case when an instance of a class is done the same**
  - Change in one variable also affect the other

```py
# Python As An Example: Mutable Instance
# --------------------------------------
class Foo:
    def __init__(self, x):
        self.x = x


def inner1(f):
    f.x = 20


def inner2(f):
    f = Foo(30) # New instance: Local scope


def outer():
    f = Foo(10)
    inner1(f)
    print(f.x) # Prints 20: f.x assigned in inner1()
    inner2(f)
    print(f.x) # Prints 20: f in inner2() is new local-scoped instance. Outside f was shadowed.
    g = None
    inner2(g)
    print(g is None) # Prints True: f in inner2() is new local-scoped instance. Outside g was shadowed.


outer()
# 20
# 20
# True
```

- The following scenario is true in other languages
  - Pass an instance of a class to a function and change the value of a field
    - **The change is reflected in the variable that was passed in**
  - Reassign the parameter in the function
    - **The change is *not* reflected in the variable that was passed in**
  - Pass `nil`/`null`/`None` for a parameter value: Setting the parameter itself to a new value
    - **Does not modify the variable in the calling function**
- **This is often explained that in other languages, class instances are *Passed-By-Reference***
  - **But that is not true**
  - Else, scenario 2 and 3 above would affect the variable
- **They are always Pass-By-Value, just as in Go**
  - **However, every instance of a class in these languages are implemented as Pointer**
  - When class instance passed to a function => **The copied passed value is the *Pointer***
  - Referring to the same memory address => Changes made to one is reflected to the other (E.g. `f` above is a pointer)
  - **Re-assigning a new instance creates a separate instance/local variable (separate memory address)**
- **The same behavior applies when using *Pointer Variables* in Go**
  - But Go gives the choice to use pointers or values for both primitives and Structs
  - Most of the time, use values
    - Make it easier to understand how and when the data is modified
    - Also reduces the work of the Garbage Collector

## Pointers Indicate Mutable Parameters

- Go constants provide names for literal expressions that can be calculated at compile time
- **Go has no mechanism to declare immutability**
- But immutability is a good thing
  - Safer from bugs
  - Easier to understand
  - More ready for change
- **Ability to choose between *Value* and *Pointer* addresses mutability/immutability**
  - **Using Pointer == The variable is mutable**
  - **Not using Pointer == The variable is not mutable**
- **Go is a *Call-By-Value* language**
  - Values passed to functions are copies
  - For non-pointer-types, called functions cannot modify original arguments
  - The original data's immutability is guaranteed
- **When passing pointers to function, the original data can be modified**
  - This is because the pointer itself is *passed-by-value*
- **When passing `nil` pointers, cannot make the value non-nil**
  - Can only reassign if a value was already assigned to the pointer
  - `nil` is a fixed location in-memory
  - Also, we cannot change the pointer as it is *passed-by-value* to the function
- **If we want the value assigned to a pointer parameter to last after exiting the function, dereference the pointer and set the value**
  - Dereferencing allows to access the value pointed by the pointer
  - Also allows to set the value pointed by the pointer
  - *Attempting to change value at an address by re-assiging a new address to a pointer will not work*
    - We cannot change the pointer as it is *Passed-By-Value* to the function
    - **We can only change a pointed value by dereferencing the pointer, then re-assign a value**

```go
// Example of Dereferencing a Pointer to Update Pointed Value
// ----------------------------------------------------------

// A function that does not dereference the parameter fails to update.
func failsToUpdate(ptrX *int) {
    // Attempting to change value at an address by re-assiging a new address to a pointer
    // Address to address
    // Does not work because ptrX was Passed-By-Value
    newValue := 20
    ptrX = &newValue
}

// A function that dereference the parameter succeed to update.
func succeedToUpdate(ptrX *int) {
    // Change a pointed value by dereferencing the pointer, then re-assign a value
    // Value to value
    newValue := 20
    *ptrX = newValue
}

func main() {
    // Example of Dereferencing a Pointer to Update Pointed Value
    // ----------------------------------------------------------
    fmt.Println("Example of Dereferencing a Pointer to Update Pointed Value:")
    fmt.Println("-----------------------------------------------------------")

    someInt := 100
    fmt.Println("someInt =", someInt)
    failsToUpdate(&someInt)
    fmt.Println("After failsToUpdate(&someInt), someInt =", someInt)
    succeedToUpdate(&someInt)
    fmt.Println("After succeedToUpdate(&someInt), someInt =", someInt)
    fmt.Println()
}
```

## Pointer: Last Resort

- Be careful when using pointers in Go
  - Can make it hard to understand data flow
  - Can create extra-work for the Garbage Collector
- E.g. Prefer instantiating a Struct instead of modifying a Struct

```go
// Don't do this
func MakeFoo(f *Foo) error {
    f.Field1 = "val"
    f.Field2 = 20
    return nil
}

// Do this instead
func MakeFoo() (Foo, error) {
  f := Foo{
    Field1: "val",
    Field2: 20,
  }
  return f, nil
}
```

- **When the function expects an interface, use pointer parameters**
- E.g. When working with JSON

```go
someJson := struct {
    Name string `json:"name"`
    Age  int `json:"age"`
}{}
err := json.Unmarshal([]byte(`{"name": "Bob", "age": 30}`), &someJson)
```

- `json.Unmarshal()`
  - Populates a variable from a slice of bytes containing JSON
  - Takes a `[]byte` and an `any` parameters
  - Value passed for `any` must be a pointer, else error
- 2 reasons for using pointer with `json.Unmarshal()`
  - *This function predates generics*
    - Without Generics, we don't know what type of value to create and return
  - *Passing a pointer gives control over memory allocation*
    - `Unmarshall` is optimized for iterative type conversion between `json` and `Struct`
    - This can be more memory-efficient
- JSON integration is very common
  - **But `json.Unmarshal()` should be treated as an exception case**
  - **When returning values from a function, favor value types**
  - **Use pointer type only when there is a state in the data type that needs to be modified**
  - Some data types used with concurrency must always be passed as pointers

## Pointer-Passing Performance

- **If a Struct is large enough, using pointer improves performance**
  - Time to pass a pointer to a function is always constant
  - The size of pointers is always the same for all data types (32-bit or 64-bit)
- Passing values to a function takes longer depending on the size of the value
  - **For data structures less than 10 MB, it is slower to return pointer than the value**
  - **The performance flips with larger data structures**
  - For the most cases, the difference might be trivial
  - For large data, consider using pointers instead of values, even if the data should be immutable

## Zero-Value vs No Value

- Pointers are often used to differentiate between variable/field assigned zero value and unassigned variable/field
  - **Use a `nil` pointer to represent unnassigned variable/field**
  - **NOTE: Be careful as pointer also indicate mutability**
- It is preferrable to use Map's comma-ok idiom instead
  - `nil` pointer as parameter is useless
  - Non-`nil` pointer as parameter means mutability
- JSON-conversion are the exception
  - When converting data back and forth between JSON and Struct, need to differentiate zero-value and no value
  - **Use a pointer value for fields in the Struct that are nullable**
- **When not working with JSON, do not use pointer to indicate no value**
  - If the value will be immutable, use a value-type paired with a boolean

## Map vs Slice

- *Any modifications made to a Map via a function is reflected in the original Map variable*
  - A Map is implemented as a pointer to a Struct
  - Passing a Map to a function is copying a pointer
  - The pointer is *passed-by-value*
- **Carefully consider before using Maps as input parameters or return values**
  - Maps are bad choices for API-design
  - They say nothing about the values contained within
  - Nothing explicitly defines any keys in the Map
  - We can only trace through the code to know what they contain
  - Maps prevent API from being self-documenting
- **Instead of using Maps, it is better to use Structs**
  - Structs also help reduce the Garbage Collector's work
- **NOTE: Maps are correct choice in certain situations**
  - Struct field names are required at compile-time
  - If the keys are not available at compile-time, a Map is ideal
- **Passing a Slice to a function has more complicated behaviors**
  - *Any modifications made to a Slice via a function is reflected in the original Slice variable*
  - However, using `append()` is **not** reflected in the original slice
  - A Slice is implemented as a Struct with 3 fields: *`int` length*, *`int` capacity*, and pointer to an underlying Array
  - *When slice is copied to a function, a copy is made off those 3 fields*
  - **Changing value in the Slice = Changing values in the Array pointed by the pointer**
  - **Changing *length* or *capacity* = Only changing the local copies of the copied `int` values**
  - The *length* and *capacity* in the original Slice remains unchanged even when its value has changed
  - If the *length* of the original Slice is smaller than the new Slice, some values would not be visible in the original Slice
  - If the Slice copy is appended to and there is not enough capacity, it moves to a different underlying array
- **A Slice passed to a function can have its content modified, but it cannot be resized**
  - Slices are passed around a lot in Go
  - By default, assume it is not mutable
- **NOTE: We can pass a Slice of any size to functions**
  - It is really just a Struct of 2 `int` values (*length* and *capacity*) and a pointer to an Array
  - Always the same size of data passed around no matter the size of the Slice
  - That is not the case with Arrays: Arrays are passed by value of their contents

### Slices As Buffers

- Useful approach for reading data from external sources
- Reading data by chunks creates a lot of unnecessary memory allocations
  - Handled automatically by the Garbage Collector
  - But the work still needs to be done when done processing
- **Writing idiomatic Go means avoid unnecessary memory allocations**
  - Create a slice of bytes once
  - Use it as a buffer to read data from source

```go
// Example of Using Slice as Buffer
// --------------------------------

file, err := os.Open(fileName)
if err != nil {
    return err
}
defer file.Close()

data := make([]byte, 100)
for {
    count, err := file.Read(data)
    process(data[:count])
    if err!= nil {
        if errors.Is(err., io.EOF) {
            return nil
        }
        return err
    }
}
```

- Created a buffer of 100 bytes
- Each loop, copy the next block of bytes into the slice
- Pass the populated portion of the buffer to `process()`
- `io.EOF` indicates that there is no more data
  - Any other error is returned

### Reducing the GC's Workload

- Using buffer is one example
- **Garbage**
  - Data that has no more pointers pointing to it
  - The memory taken up by the data can be reused
  - If the memory is not recovered, the program's memory usage would continue to grow
  - Can result to *Stackoverflow* bug
- **Garbage Collector**
  - Automatically detect unused memory and recover it for reuse
  - Simplify memory management
- *Just because we have a Garbage Collector does not mean we should create a lot of garbages*
- **Stack**
  - Consecutive blocks of memory
  - Every function call in a thread of execution shares the same stack
  - Simple and fast memory allocation
  - *Stack Pointer* tracks the allocated location
    - Allocating additional memory is done by changing the value of the Stack Pointer
  - Invoking function
    - New stack frame is created for the function
    - New scope
    - Local variables are stored on the stack
  - Each new variable moves the Stack Pointer by the size of the value
  - Returning functions
    - Return values are returned to the caller via the Stack
    - Stack Pointer is moved back to the beginning of the stack frame
    - Deallocating stack memory used by the function
  - **NOTE: Since Go 1.17, uses combination of *Registers* and Stack to pass values into functions**
    - *Register* - Small set of high-speed memory on the CPU
    - Faster but more complicated
    - The general concepts of Stack-Only functions still apply
  - ***To store on the Stack, we have to know the value's size at compile-time***
    - Case for all value types: Constant size
    - For value types, the size is considered part of the type
    - Because the size is known, they can be allocated on the Stack
    - *The size of Pointers is also known and constant (32-bits or 64-bits)*
- **NOTE: Go can increase the size of the Stack while the program is running**
  - Each goroutine has its own stack
  - Goroutines are managed by the Go Runtime
  - Advantages:
    - Go stacks start small
    - Use less memory
  - Disadvantages:
    - When stack needs to grow, all data on the stack is copied
    - This copy can be slow
    - It is possible to write code that causes the stack to shrink and grow over and over
- **Different rules for the data that the pointer points to**
  - To allocate them on the Stack:
    - Data must be a local variable with known data-size at compile-time
    - Pointer cannot be returned from a function
    - If pointer is passed to a function, these rules must still hold
  - If unable to store on the Stack, the data that the pointer points to *escapes* the Stack
    - It is stored on the **Heap**
- **Heap**
  - Memory that is managed by the *Garbage Collector*
  - In other languages (C/C++), it is managed manually
  - **Any data on the Heap is valid as long as it can be tracked back to pointer on the Stack**
  - If there are no pointers pointing to the data, it becomes *Garbage*
  - The *Garbage Collector* clears it up
