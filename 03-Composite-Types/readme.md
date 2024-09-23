# Composite Types

---

- [Arrays](#arrays)
  - [Array Declarations](#array-declarations)
  - [Comparing Arrays](#comparing-arrays)
  - [Multidimensional Array](#multidimensional-array)
  - [Array Access](#array-access)
  - [Getting Array Length](#getting-array-length)
  - [Limitations: Why Arrays Are Rarely Used](#limitations-why-arrays-are-rarely-used)
  - [Purpose: Why Arrays Exist If They Are So Limited](#purpose-why-arrays-exist-if-they-are-so-limited)
- [Slices](#slices)
  - [Working With Slices](#working-with-slices)
  - [Multidimensional Slice](#multidimensional-slice)
  - [Slice Access](#slice-access)
  - [Differences With Arrays](#differences-with-arrays)
    - [The Zero-Value for slice is `nil`](#the-zero-value-for-slice-is-nil)
    - [Slices Are Not Comparable With `==`](#slices-are-not-comparable-with-)
  - [Slices Helper Functions From `slices`](#slices-helper-functions-from-slices)
  - [Built-In `len()` Function](#built-in-len-function)
  - [Built-In `append()` Function](#built-in-append-function)
  - [Length vs Capacity](#length-vs-capacity)
  - [NOTE: About Go Runtime](#note-about-go-runtime)
  - [Built-In `make()` Function](#built-in-make-function)
  - [Built-In `clear()` Function](#built-in-clear-function)
  - [Which Slice Declarations To Use](#which-slice-declarations-to-use)
  - [Subslicing Slices](#subslicing-slices)
  - [Built-In `copy()` Function](#built-in-copy-function)
- [Converting Array To Slice](#converting-array-to-slice)
- [Converting Slice To Array](#converting-slice-to-array)
- [Strings, Runes, Bytes](#strings-runes-bytes)
  - [UTF-8](#utf-8)
- [Maps](#maps)
  - [Comparison With Slices](#comparison-with-slices)
  - [Map vs Hash Map](#map-vs-hash-map)
  - [Reading and Writing a Map](#reading-and-writing-a-map)
  - [*Comma-Ok* Idiom](#comma-ok-idiom)
  - [Built-In `delete()` Function](#built-in-delete-function)
  - [Map: Built-In `clear()` Function](#map-built-in-clear-function)
  - [Comparing Maps](#comparing-maps)
  - [Using Maps As Sets](#using-maps-as-sets)
- [Structs](#structs)
  - [Struct Definition and Declaration](#struct-definition-and-declaration)
  - [Struct Field Access](#struct-field-access)
  - [Anonymous Struct](#anonymous-struct)
    - [When to Use Anonymous Struct](#when-to-use-anonymous-struct)
  - [Comparing And Converting Structs](#comparing-and-converting-structs)

---

## Arrays

- Arrays are fixed-size lists of same-type elements
- Arrays are rarely used in Go
- **All elements in an array must be of the same type of the array**

### Array Declarations

- There are different declaration styles

```go
// Basic array declaration
// -----------------------
var arrA [3]int
```

- **If no value is assigned, all the elements are initialized with the Zero-Value of the type**
- But we can also specify the values with an *Array Literal*

```go
// Array declaration with Array literal
// ------------------------------------
var arrB = [3]int{10, 20, 30}
```

- We can specify a *Sparse Array* as well
- Only specify the Non-Zero elements using indices

```go
// Array declaration with specific non-zero elements
// -------------------------------------------------
var arrC = [15]int{1, 5: 4, 6, 10, 100, 15}
```

- **When using array literal, we can replace the length of the array to be dynamic with `...`**
- However, this can be harder to read

```go
// Array declaration with ... length
// ---------------------------------
var arrD = [...]int{100, 200, 300}
```

- We can combine any of the above with the `:=` operator as well

```go
// Array declaration with :=
// -------------------------
arrE := [...]int{250, 350, 450, 1000}
```

### Comparing Arrays

- We can use `==` and `!=` to compare arrays
- **Arrays are equal if:**
  - **Same length**
  - **Same values**
  - **Same order**

```go
// Comparing Arrays
// ----------------
fmt.Println("Comparing Arrays:")
fmt.Println("-----------------")

arrF := [...]int{1, 2, 3, 4, 5}
arrG := [5]int{1, 2, 3, 4, 5}

fmt.Println("arrF =", arrF)
fmt.Println("arrG =", arrG)
fmt.Println("arrF == arrG?", arrF == arrG)
```

### Multidimensional Array

- **Go does not support multidimensional arrays**
- However, we can simulate them with *Array of Arrays*

```go
// Multidimensional Array
// ----------------------
fmt.Println("Multidimensional Arrays:")
fmt.Println("------------------------")

// Array of array of integers: 5 x 5
var arrH [5][5]int

fmt.Println(arrH)
```

### Array Access

- Arrays in Go are read and written via the `arr[index]` syntax

```go
// Array Access
// ------------
fmt.Println("Array Access:")
fmt.Println("-------------")

arrI := [3]int{10, 20, 30}

fmt.Println("arrI =", arrI)

// Accessing the values
fmt.Println("arrI[0] =", arrI[0])
fmt.Println("arrI[1] =", arrI[1])
fmt.Println("arrI[2] =", arrI[2])

// Changing the array values
arrI[0] = 15
arrI[1] = 25
arrI[2] = 35

fmt.Println("After changing the array values:")
fmt.Println("arrI =", arrI)
```

- We cannot access beyond the length of the array
  - Compile-time error
  - **Out-of-bounds access will *panic* at runtime**
- **We cannot access with negative index**
  - Compile-time error

### Getting Array Length

- Use the built-in function `len()` to  get the length of an array

```go
// Length of an array
// ------------------
fmt.Println("Length of An Array:")
fmt.Println("-------------------")

arrJ := [3]int{10, 20, 30}
arrK := [...]int{15, 25, 35}

fmt.Println("arrJ =", arrJ, "=> length:", len(arrJ))
fmt.Println("arrK =", arrK, "=> length:", len(arrK))
```

### Limitations: Why Arrays Are Rarely Used

- **Go considers the size of the array to be part of the type of the array**
  - E.g. `[3]int` is different from `[4]int`
- **We cannot use variable to specify the size**
  - Types must be resolved at compile-time
  - Because `[n]type` is a type, it must be resolved at compile-time
- **We cannot use type-conversion to convert arrays of different size to the same size**
  - We cannot write functions to work with arrays of any size
  - We cannot assign arrays of any sizes to the same variable
- ***Do not use array unless you know exactly the needed length ahead of time***

### Purpose: Why Arrays Exist If They Are So Limited

- **The main purpose of arrays is to be used behind-the-scenes for *Slices***

## Slices

- Most of the time, *Slices* are what we should use for a *sequence-based* data structure
- **Slice is built on top of an underlying array**
- **Slice length is dynamic: Can grow and shrink as needed**
- **The slice length is NOT part of its type**
  - Removes the limitation from arrays
  - Allows to write functions to process slices of any size

### Working With Slices

- Very similar to working with Arrays
- With some subtle differences
- **Declaration With Literal: Do not specify the length**

```go
// Slice declaration with Slice literal
// ------------------------------------
sliceA := []int{10, 20, 30, 40}
```

- We can specify a *Sparse Slice* as well
  - Only specify the Non-Zero elements using indices

```go
// Slice declaration with specific non-zero elements
// -------------------------------------------------
sliceB := []int{1, 5: 4, 6, 10, 100, 15}
```

- The *Zero-Value* of a slice is a *Nil-Slice* `[]`

```go
// The Zero-Value of a slice is nil
var sliceNil []int
fmt.Println("sliceNil =", sliceNil)
fmt.Println("sliceNil is nil?", sliceNil == nil)
```

- For better definition, we use `make()`
- This initialize the slice with Zero-values instead of `nil`

```go
mySlice = make([]int, 4) // => mySlice = [0, 0, 0, 0]
```

### Multidimensional Slice

- **Go does not support multidimensional slices**
- However, we can simulate them with *Slice of Slices*

```go
// Multidimensional Slice
// ----------------------
fmt.Println("Multidimensional Slices:")
fmt.Println("------------------------")

// Slice of slices of integers
// Default value is a nil slice
var sliceC [][]int

// For better definition, we use make()
// This initialize the slice with Zero-values instead of nil
sliceC = make([][]int, 4)

sliceC[0] = []int{1, 2, 3, 4}
sliceC[1] = []int{5, 6, 7, 8}
sliceC[2] = []int{9, 10, 11, 12}
sliceC[3] = []int{13, 14, 15, 16}

fmt.Println("sliceC =", sliceC)
```

### Slice Access

- Same rules as with *Arrays*
  - Read and write via the `slice[index]` syntax
- **We cannot access beyond the length of the slice**
  - Compile-time error
  - **Out-of-bounds access will *panic* at runtime**
- **We cannot access with negative index**
  - Compile-time error

```go
// Slice Access
// ------------
fmt.Println("Slice Access:")
fmt.Println("-------------")

sliceD := []int{10, 20, 30}

// Accessing the values
fmt.Println("sliceD =", sliceD)
fmt.Println("sliceD[0] =", sliceD[0])
fmt.Println("sliceD[1] =", sliceD[1])
fmt.Println("sliceD[2] =", sliceD[2])

// Changing the values
fmt.Println("Changing the values:")
sliceD[0] = 15
sliceD[1] = 25
sliceD[2] = 35

fmt.Println("sliceD =", sliceD)
```

### Differences With Arrays

#### The Zero-Value for slice is `nil`

- The Zero-Value of array is the Zero-Value of its type
- The Zero-Value of slice is `nil`

```go
// Declaring a slice with Zero-Value: nil
// --------------------------------------
var sliceNil []int
```

- **`nil` is slightly different than `null` in other languages**
  - `nil` is an identifier that represents lack of value for some types
  - **`nil` is untyped**: Can be used for different types
- ***A `nil` slice contains nothing***

#### Slices Are Not Comparable With `==`

- **Using `==` and `!=` on 2 slices result in compile-time error**
- We can only use `slice == nil` or `slice != nil`

```go
// Checking if slice is nil
// ------------------------
var sliceE []int

fmt.Println("sliceE is nil?", sliceE == nil)
```

### Slices Helper Functions From `slices`

- `slices` package includes helper functions for dealing with slices
- Also has functions for comparing slices

Function|Description
:-|:-
`slices.Equal()`|- Takes 2 slices *with comparable elements*<br>- Return `true` if the slices are same length with all equal elements<br>- Else, return `false`<br>- *The slice elements must be comparable*
`slices.EqualFunc()`|- Allows to pass a function argument to customize conditions for equality<br>- *The slice elements do not have to be comparable*

```go
// Using Slices Helper Functions: slices
// -------------------------------------
fmt.Println("Using Slices Helper Functions:")
fmt.Println("------------------------------")

sliceX := []int{1, 2, 3, 4, 5}
sliceY := []int{1, 2, 3, 4, 5}
sliceZ := []int{1, 2, 3, 4, 5, 6}
sliceS := []string{"a", "b", "c"}

fmt.Println("sliceX =", sliceX)
fmt.Println("sliceY =", sliceY)
fmt.Println("sliceZ =", sliceZ)
fmt.Println("sliceS =", sliceS)

fmt.Println("slices.Equal(sliceX, sliceY) ?", slices.Equal(sliceX, sliceY))
fmt.Println("slices.Equal(sliceX, sliceZ) ?", slices.Equal(sliceX, sliceZ))
// This one will not compile: Different types for each elements
// fmt.Println("slices.Equal(sliceX, sliceS) ?", slices.Equal(sliceX, sliceS))
```

- **NOTE: `reflect` package contains a `DeepEqual` function that can compare almost anything**
  - Legacy function, primarily for testing
  - Was often used to compare slices before `slices.Equal` and `slices.EqualFunc`
  - **Do not use `reflect.DeepEqual` in new codes**
  - It is slower and less safe

### Built-In `len()` Function

- **Built-in function for getting length of collections**
  - Can be used on both **Arrays** and **Slices**
  - Also works on **Strings**, **Maps**, **Channels**
  - **However, any other type will result in compile-time error**
- Passing a `nil` slice returns `0`
- Parameter can be any type of array or slice

```go
// Length of a Slice
// -----------------
sliceF := []int{10, 20, 30, 40, 50}
fmt.Println("sliceF =", sliceF, "length =", len(sliceF))
```

### Built-In `append()` Function

- Used to grow a slice by appending new elements to its end
- **Takes 2 parameters**
  - *Slice of any type*
  - *Value of that type to append to the slice*
- Returns a new slice of the same type
  - Assign it to a variable to capture the result
  - We typically re-assign to the original variable
  - **NOTE: Not assigning the returned value to a variable is a compile-time error**
    - *Go is a call-by-value language*: All passed parameters are a copy of the value passed in
    - The function appends unto the copy of the slice
    - The copy of the slice is returned with appended values
    - *The copy needs to be re-assigned to make the change permanent*
- **`append()` can take more than 1 value at once to append**

```go
// Appending to a slice
// --------------------
fmt.Println("Appending to a Slice:")
fmt.Println("---------------------")

// Appending to a nil slice
var sliceG []int
// Appending to an existing slice
sliceH := []int{1, 2, 3, 4, 5}

fmt.Println("Before: sliceG =", sliceG)
fmt.Println("Before: sliceH =", sliceH)

sliceG = append(sliceG, 100)
// We can append more than 1 value at once
sliceH = append(sliceH, 6, 7, 8, 9)
// Appending one slice to another

fmt.Println("After: sliceG =", sliceG)
fmt.Println("After: sliceH =", sliceH)
```

- We can also append one slice unto another
  - Using the `slice...` operator
  - **This has the effect of *extending* the slice**

```go
// Extending a slice with another slice
// ------------------------------------
fmt.Println("Extending a Slice:")
fmt.Println("------------------")

sliceH = append(sliceH, sliceG...)

fmt.Println("Extending sliceH with sliceG =", sliceH)
```

### Length vs Capacity

- **A slice is a sequence of values**
  - *Each element assigned to consecutive memory locations*
  - Closer memory location == Quick read and write
- **Length of slice = Number of consecutive memory locations that *have been assigned a value***
  - Each value added via `append()` increases the length
- **Capacity of slice = Number of consecutive memory locations that *have been reserved***
  - Can be larger than *length*
  - When `length == capacity`, there are no more rooms to append additional values
  - If so:
    - The runtime allocate a new backing array with larger capacity
    - ***The values in the original backing array are copied into the new backing array***
    - The new value is appended at the end of the new array
    - The slice is updated to refer to the new backing array
    - The updated slice is returned
    - **NOTE: Slice is a reference type**
- **For this reason, growing a slice is expensive and can cause performance hit: $O(n)$**
  - Allocate new memory
  - Copy all contents from old memory to new memory
  - Garbage collection
- **Go usually increases capacity by more than one each time it runs out of capacity**
  - Double when the current capacity is less than 256
  - For bigger slice: Increase by `(currCap + 768) / 4`
  - Slowly converge to 25% growth

$$
\begin{aligned}
  currCap <= 256&: newCap = currCap + currCap \\
  currCap > 256&: newCap = currCap + (currCap + 768) / 4
\end{aligned}
$$

- **Built-in `cap()` function returns the capacity of a slice**
  - Less used than `len()`
  - Mostly used to check if a slice is large enough to hold a data
  - If not, a call to `make()` is needed to create a new slice
  - *NOTE: We can also pass an array to `cap()`*
    - But it is always `len(arr) == cap(arr)`
    - Not useful to use in code

```go
// Appending and Slice Capacity
// ----------------------------
fmt.Println("Appending and Slice Capacity:")
fmt.Println("-----------------------------")

var xSlice []int
fmt.Println("Slice\t\t\tLen\tCap")
fmt.Println(xSlice, "\t\t\t", len(xSlice), "\t", cap(xSlice))

xSlice = append(xSlice, 10)
fmt.Println(xSlice, "\t\t\t", len(xSlice), "\t", cap(xSlice))

// We reach the current slice cap here
xSlice = append(xSlice, 20)
fmt.Println(xSlice, "\t\t", len(xSlice), "\t", cap(xSlice))

// This one double the cap
// Copy into a new slice
xSlice = append(xSlice, 30)
fmt.Println(xSlice, "\t\t", len(xSlice), "\t", cap(xSlice))

// We reach the current slice cap here
xSlice = append(xSlice, 40)
fmt.Println(xSlice, "\t\t", len(xSlice), "\t", cap(xSlice))

// This one double the cap
// Copy into a new slice
xSlice = append(xSlice, 50)
fmt.Println(xSlice, "\t", len(xSlice), "\t", cap(xSlice))
```

- **While it is nice that slices can grow dynamically, it is more efficient to size them once**
  - If the size can be determined at the beginning
  - **We can use `make()` to specify the capacity**

### NOTE: About Go Runtime

- **Go Runtime provides different services**
  - Set of libraries that enables programs written in Go to run
  - Memory Allocation
  - Garbage Collection
  - Concurrency Support
  - Networking
  - Built-in types and functions
- **The Runtime is compiled and bundled into every Go programs**
  - No VM involved
  - Allows easier distribution of compiled Go programs
  - Avoids compatibility issues between runtime and program
  - But can slightly increase the size of the compiled program
  - Minimum size 2MB, even for simplest programs

### Built-In `make()` Function

- We currently declare slices without allowing to create an empty slice with specified length and capacity
- To do that, we use `make()` instead
  - Allows to specify the type, length, and capacity
- ***NOTE: `make()` can be used to make a `slice`, `map`, or `chan`***

```go
// Declaring a slice using make()
// Length: 5, Capacity: 5
sliceI := make([]int, 5)
```

- ***When using `make()`, all elements are initialized to `Zero-Value` instead of `nil`***
- **WARNING: Since elements are initialized, we should not use `append()` to add to the slice**
  - I.e. when initializing with specified `length`
  - The slice is not empty
  - Using `append()` will increase the length and double the capacity of the slice
- We can also specify a different initial capacity

```go
// Declaring a slice using make()
// Length: 5, Capacity: 10
sliceJ := make([]int, 5, 10)
```

- We can also make a slice of length `0` but with a specified capacity
  - **Using this approach allows to use `append()` to the slice as usual**

```go
// Declaring a slice using make()
// Length: 0, Capacity: 10
sliceK := make([]int, 0, 10)
sliceK = append(sliceK, 1, 2, 3, 4, 5)
```

- ***WARNING: Never specify a capacity less than the length***
  - If done via constant, it is a compile-time error
  - If done via variables, it is a runtime error

```go
// Declaring a slice using make()
// ------------------------------
fmt.Println("Declaring a slice using make():")
fmt.Println("-------------------------------")

// Length: 5, Capacity: 5
sliceI := make([]int, 5)
// Length: 5, Capacity: 10
sliceJ := make([]int, 5, 10)
// Length: 0, Capacity: 10
sliceK := make([]int, 0, 10)

fmt.Println("Before Append:")
fmt.Println("\tsliceI =", sliceI, "len(sliceI) =", len(sliceI), "cap(sliceI) =", cap(sliceI))
fmt.Println("\tsliceJ =", sliceJ, "len(sliceJ) =", len(sliceJ), "cap(sliceJ) =", cap(sliceJ))
fmt.Println("\tsliceK =", sliceK, "len(sliceK) =", len(sliceK), "cap(sliceK) =", cap(sliceK))

sliceK = append(sliceK, 1, 2, 3, 4, 5)

fmt.Println("After Append to sliceK:")
fmt.Println("\tsliceK =", sliceK, "len(sliceK) =", len(sliceK), "cap(sliceK) =", cap(sliceK))
```

### Built-In `clear()` Function

- `clear()` resets all elements in a slice back to the zero-value
- ***WARNING: The length of the slice remains unchanged***
  - It only resets the elements to their *zero-value*
  - It does not reset the length of the slice

```go
// Resetting slice using clear()
// -----------------------------
fmt.Println("Resetting slice using clear():")
fmt.Println("------------------------------")

sliceL := []string{"first", "second", "third"}
sliceLInt := []int{100, 200, 300}
fmt.Println("Before:")
fmt.Println("\tsliceL =", sliceL, "len =", len(sliceL), "cap =", cap(sliceL))
fmt.Println("\tsliceLInt =", sliceLInt, "len =", len(sliceLInt), "cap =", cap(sliceLInt))
clear(sliceL)
clear(sliceLInt)
fmt.Println("After:")
fmt.Println("\tsliceL =", sliceL, "len =", len(sliceL), "cap =", cap(sliceL))
fmt.Println("\tsliceLInt =", sliceLInt, "len =", len(sliceLInt), "cap =", cap(sliceLInt))
```

### Which Slice Declarations To Use

- **Primary goal is to try to minimize the number of times the slice has to grow**
- If the slice will not need to grow at all, create a `nil` slice

```go
// Slice with nil value
var data []int
```

- ***WARNING: A slice with zero-length and zero-capacity is different from a `nil` slice***
  - Comparing the 2 will return `false`
  - Comparing a `nil` to a `nil` returns `true`
  - *For simplicity, favor `nil` slice*
  - **A zero-length slice is only useful when converting a slice to JSON**

```go
// Slice with zero length and zero capacity
var data = []int{}
// Equivalent
data := make([]int, 0, 0)
```

- With specified starting values or if slice's values will not change, use slice literal

```go
// Slice literal
data := []int{1, 2, 3, 4, 5}
```

- **If you do not know the slice values but have an idea of how large the slice will be, use `make()`**
  - If using slice as a buffer: *Specify a non-zero length*
  - If you know the exact size needed: *Specify the length and index into the slice to set the values*
  - In other situations: *Zero length and specify the capacity*
- **The choice is usually between the 2nd and the 3rd option**
  - Option 2 might be slower but introduces less bugs
  - **However, `append()` always increases the length**
  - ***NOTE: Specifying the length with `make()` must be done carefully***

### Subslicing Slices

- **Slice expression creates a slice from a slice**
- Written inside brackets
- Consists of *starting* and *ending* offsets separated by colon
  - Ending offset is *up-to-but-not-including*
  - Default starting offset is `0`
  - Default ending offset is the end of the slice

```go
// Subslicing slice
// ----------------
fmt.Println("Subslicing slice:")
fmt.Println("-----------------")

sliceM := []string{"a", "b", "c", "d"}
subM1 := sliceM[:2]
subM2 := sliceM[1:]
subM3 := sliceM[1:3]
subM4 := sliceM[:]

fmt.Println("sliceM =", sliceM)
fmt.Println("sliceM[:2] =", subM1)
fmt.Println("sliceM[1:] =", subM2)
fmt.Println("sliceM[1:3] =", subM3)
fmt.Println("sliceM[:] =", subM4)
```

- ***WARNING: Slicing a slice does not make a separate copy***
  - Assigning the slice of the slice to a variable create an alias to the original slice variable
  - **This is because they will share the same underlying array**
  - The original slice and subslices are all linked
  - **Changing an element in one variable affects all other aliases**

```go
// Modifying one slice affects all others
subM1[1] = "y"
subM2[0] = "x"
subM3[1] = "z"

fmt.Println("sliceM =", sliceM)
fmt.Println("subM1 =", subM1)
fmt.Println("subM2 =", subM2)
fmt.Println("subM3 =", subM3)
fmt.Println("subM4 =", subM4)
```

- **This behavior can get extra-confusing when combined with `append()`**
  - `cap(subSlice) = cap(origSlice) - startingOffsetSubSlice`
  - Elements of the original slice beyond the end of the subslice are shared between them
  - ***This can cause slices and subslices to append and overwrite each other's data***

```go
// Slice of slice with append()
// ----------------------------
fmt.Println("Slice of slice with append():")
fmt.Println("-----------------------------")

fmt.Println("BEFORE:")
fmt.Println("\tsliceM", "cap =", cap(sliceM), sliceM)
fmt.Println("\tsubM1", "cap =", cap(subM1), subM1)

subM1 = append(subM1, "zz")

fmt.Println("AFTER:")
fmt.Println("\tsliceM", "cap =", cap(sliceM), sliceM)
fmt.Println("\tsubM1", "cap =", cap(subM1), subM1)
```

- To avoid this issue, either:
  - **Never use `append()` with subslices**
  - **Always use *full slice expression* with `append()`**
    - Includes a 3rd part
    - *Indicates last position in the capacity of the original slice that is available for the subslice*
    - **Subtract the starting offset from this number to get the subslice capacity**
    - $x-SubsliceStartingOffset=SubliceCapacity$ => $x=SubliceCapacity+SubsliceStartingOffset$
    - The $SubliceCapacity$ is determined by the `start:end` part of the slice
- Makes clear how much memory is shared between original slice and subslice

```go
// Full Slice Expressions
// ----------------------
fmt.Println("Full Slice Expressions:")
fmt.Println("-----------------------")

sliceN := make([]string, 0, 5)
sliceN = append(sliceN, "a", "b", "c", "d", "e")
subN1 := sliceN[1:4:4]
subN2 := sliceN[2:4:4]

fmt.Println("sliceN =", sliceN)
fmt.Println("subN1 =", subN1)
fmt.Println("subN2 =", subN2)
fmt.Println("Capacities:", cap(sliceN), cap(subN1), cap(subN2))

subN1 = append(subN1, "i", "j", "k")
sliceN = append(sliceN, "x", "y")
subN2 = append(subN2, "z", "zz")

fmt.Println("Operations:")
fmt.Println("\tappend(subN1, \"i\", \"j\", \"k\")")
fmt.Println("\tappend(sliceN, \"x\", \"y\")")
fmt.Println("\tappend(subN2, \"z\", \"zz\")")

fmt.Println("sliceN =", sliceN)
fmt.Println("subN1 =", subN1)
fmt.Println("subN2 =", subN2)
```

- `subN1` has  capacity of `3` (`4 - 1`)
- `subN2` has a capacity of `2` (`4 - 2`)
- We are basically limiting the capacity of the sublice to its length
  - **By doing so, appending to the subslices would now force to create new slices for them**
  - **This breaks the aliases with the other subslices**
  - **However, they still share underlying array until after a first `append()` is executed**
- ***NOTE:***
  - ***If possible, avoid modifying slices after they have been sliced or produced by slicing***
  - **Else, use the *Full-slice expression* to prevent `append()` from sharing capacity**
  - ***Overall, it is better to create independent copies of subslices using `copy()`***

### Built-In `copy()` Function

- **Allows to create a subslice independant of the original slice**
- Takes 2 parameters:
  - `dest` slice
  - `src` slice
- Copies as many values as possible from `src` to `dest`
  - Limited by whichever slice is smaller
  - **Capacities do not matter: Only the lengths**
  - Specifying capacity with 0-length for the destination will not work
- **Returns the number of elements copied**

```go
// Slicing Copy
// ------------
fmt.Println("Slicing copy:")
fmt.Println("-------------")

sliceO := []int{1, 2, 3, 4, 5}
sliceP := make([]int, 10)

fmt.Println("Before Copying:", "src =", sliceO, "dest =", sliceP)
num := copy(sliceP, sliceO)
fmt.Println("After Copying:", "src =", sliceO, "dest =", sliceP, "copied", num, "elements")
```

- Can also copy a subset of a slice

```go
// Copying a subset
// ----------------
fmt.Println("Copying subset of slice:")
sliceQ := make([]int, 2)
num = copy(sliceQ, sliceO)
fmt.Println("After Copying:", "src =", sliceO, "dest =", sliceQ, "copied", num, "elements")
```

- Can also copy from the middle of the source slice (or any subslice)
- The length of the destination slice can be derived from the subslice
- **NOTE: We do not have to assign the return of `copy()` if we do not need it**

```go
// Copying subset from the middle of slice
// ---------------------------------------
fmt.Println("Copying subset from the middle of slice:")
fmt.Println("----------------------------------------")

// Length of the destination slice can be derived from the subslice
// If the return value is not needed, ignore
sliceR := make([]int, len(sliceO[2:]))
copy(sliceR, sliceO[2:])

fmt.Println("After Copying:", "src =", sliceO, "dest =", sliceR)
```

- We can copy between 2 slices that cover overlapping sections of an underlying slice

```go
// Overlapping copies
// ------------------
fmt.Println("Overlapping copies:")
fmt.Println("-------------------")

num = copy(sliceO[:3], sliceO[1:])
fmt.Println("After Copying:", sliceO, "copied", num, "elements")
```

- **NOTE: We can use `copy()` on arrays as well**
  - Can be source or destination

```go
// Using copy() with arrays
// ------------------------
fmt.Println("Using copy() with arrays:")
fmt.Println("-------------------------")

sliceT := []int{1, 2, 3, 4} // Slice
arrL := [4]int{5, 6, 7, 8}  // Array
sliceU := make([]int, 2)

fmt.Println("Before copying:", sliceT, arrL, sliceU)
copy(sliceU, arrL[:])
fmt.Println("After copying from array:", sliceU, arrL)
copy(arrL[:], sliceT)
fmt.Println("After copying into array:", arrL, sliceT)
```

## Converting Array To Slice

- We can make a slice from an array
- Using slice expression
- Allows to use an array as a slice
- **To convert an entire array into a slice, use `[:]`**
- We can also convert only a subset with `[start:]` or `[:end]`
- ***WARNING: This results in the same memory-sharing issue as a slice of slice***
  - **Use `copy()` to avoid memory-sharing issue**

```go
// Converting Array To Slice
// -------------------------
fmt.Println("Converting array to slice:")
fmt.Println("--------------------------")

bArray := [4]int{100, 200, 300, 400}
fmt.Println("bArray Before =", bArray)

bSlice := make([]int, len(bArray[:]))
copy(bSlice, bArray[:])
bSlice[0] = 10000 // No memory-sharing after using copy()
bArray[3] = 40000 // No memory-sharing after using copy()

fmt.Println("bArray =", bArray)
fmt.Println("bSlice =", bSlice)
```

## Converting Slice To Array

- **For *slice* to *array*, we use a type-conversion**
  - Can convert an entire slice to an array of the same type
  - Can create an array from a subset of a slice
  - **The data is *copied* into a new memory location**
  - **There is no memory-sharing issue**

```go
// Converting Slice to Array
// -------------------------
fmt.Println("Converting slice to array:")
fmt.Println("--------------------------")

cSlice := []int{10, 20, 30, 40}
fmt.Println("cSlice Before =", cSlice)

cArray := [4]int(cSlice)
smallArray := [2]int(cSlice)
cSlice[0] = 10000 // No memory-sharing
cArray[3] = 40000 // No memory-sharing
fmt.Println("cSlice =", cSlice)
fmt.Println("cArray =", cArray)
fmt.Println("smallArray =", smallArray)
```

- **The size of the array must be specified at compile-time**
  - Cannot use `[...]` notation
  - The size of the array must be `<=` the size of the slice
  - But the compiler cannot detect this
  - Code will panic at runtime if this is not the case
- **NOTE: We can use a type-conversion to convert a slice into a pointer to an array**
  - The storage between the 2 will be shared
  - This will have a memory-sharing issue

```go
slice := []int{1,2,3,4}
arrayPtr := (*[4]int)(slice)
```

## Strings, Runes, Bytes

- **Strings are not made of *Runes***
  - **Using a sequence of *bytes* instead**
  - **Using `len()` on string returns the number of bytes in the string**
  - Bytes do not have to be in any particular character encoding
  - But usually assumed to be UTF-8 code points
  - String literals are UTF8 unless using hexadecimal escapes
- *NOTE: Go Source-code is always in UTF-8*
- We can extract values from a string using indexing
  - Zero-based indexing
  - Slice-expression also works with strings

```go
// Example of Strings
// ------------------
fmt.Println("Example of Strings:")
fmt.Println("-------------------")

var strA string = "Hello there!"
var strASub1 byte = strA[6]
var strASub2 string = strA[4:7]
var strASub3 string = strA[:5]
var strASub4 string = strA[6:]

fmt.Println("strA =", strA)
fmt.Println("strASub1 =", strASub1)
fmt.Println("strASub2 =", strASub2)
fmt.Println("strASub3 =", strASub3)
fmt.Println("strASub4 =", strASub4)
```

- **Since strings are immutable, they do not have the modification issue as with slices**
- However, a string is a sequence of bytes
- A UTF8 code-point (Rune) can be 1 to 4 bytes
  - When dealing with non-english languages and emojis, this can be an issue
  - **We can break the code-points of some characters**
  - ***Only use slicing and indexing when string contains only 1-byte characters***
  - ***Better option: extract substrings and code-points using functions from `strings` and `unicode/utf8` packages***

```go
// String Code-point Issue
// -----------------------
fmt.Println("Example of Strings With Code-Point Issues:")
fmt.Println("------------------------------------------")

var strB string = "Hello 😊!" // 😊 takes 4 bytes
var strBSub1 byte = strB[6]
var strBSub2 string = strB[4:7]
var strBSub3 string = strB[:5]
var strBSub4 string = strB[6:]

fmt.Println("strB =", strB, "len =", len(strB)) // len() returns the number of bytes
fmt.Println("strBSub1 =", strBSub1)
fmt.Println("strBSub2 =", strBSub2)
fmt.Println("strBSub3 =", strBSub3)
fmt.Println("strBSub4 =", strBSub4)
```

- We can get the length of a string using the `len()` function
  - **However, it returns the length in *Bytes*, not in *UTF-8 Codepoints***
  - This is also an issue with non-english languages and emojis
  - ***Only use `len()` when string contains only 1-byte characters***
- **In summary**
  - String = Sequence of *Bytes* (8-bits)
  - Rune = UTF8 Codepoint, between 1 to 4 bytes (8-bits to 32-bits)
    - 1 Byte if only ASCII characters
    - Possibly 2 Bytes+ if beyond ASCII
  - `len(str)` return number of *Bytes*

```go
// Length of String Issue
// ----------------------
fmt.Println("Length of String Issue:")
fmt.Println("-----------------------")

var strC string = "Hello 😊!"
fmt.Println("strC =", strC)
// len() returns the number of bytes, not runes
fmt.Println("len(strC) =", len(strC))
```

- **Go has a complicated relationship between strings, runes, and bytes**
  - *A single rune or byte can be converted to a string with `string()`*
  - *A string can be converted back and forth to `[]byte` or `[]rune`*
- **Most data in Go is used as a sequence of bytes**
  - **Most common string-type conversion is to/from `[]byte`**
  - `[]rune` are uncommon

```go
// Strings, runes, bytes type conversion
// -------------------------------------
fmt.Println("Strings, runes, bytes type conversion:")
fmt.Println("--------------------------------------")

var charA rune = 'a'
var charB byte = 'b'
stringA := string(charA)
stringB := string(charB)
stringC := "Hello 😊!"
var bytesStringC []byte = []byte(stringC)
var runesStringC []rune = []rune(stringC)

fmt.Println("charA =", charA)
fmt.Println("charB =", charB)
fmt.Println("stringA =", stringA)
fmt.Println("stringB =", stringB)
fmt.Println("stringC =", stringC)
fmt.Println("bytesStringC =", bytesStringC)
fmt.Println("runesStringC =", runesStringC)
```

- **NOTE: `go vet` blocks `string(someInt)`**
  - Other than from `rune` and `byte`
  - This is because the integer is interpreted as ASCII code
  - E.g. `string(65)` becomes `"A"`, not `"65"`
  - **It is not a good idea to convert an integer into a string with `string()`**
  - For that case, use `strconv.Itoa()` instead

### UTF-8

- The most commonly-used encoding for Unicode
- **Unicode uses a total of 4 bytes/32-bits for each codepoint ("character")**
  - UTF-32 - Strict 4-bytes codepoint, but waste space
  - UTF-16 - One or two 2-bytes codepoint, but also waste space
  - UTF-8
    - 1-byte codepoint for value below 128 (Most English characters in ASCII)
    - Expand to max 4-bytes for larger values
    - Worst-case is the same as UTF-32
    - We cannot randomly access a string encoded in UTF-8
    - Need to start at the beginning of the string and count
    - **Not required in Go but strongly recommended**
- **Instead of using *slice expression* and *index expression*, use `strings` and `unicode/utf8` packages**

## Maps

- Allows to associate one value to another
- Format: `map[KeyType]ValueType`
- Zero-Value: `nil`
- Length of a `nil` map: `0`
- Attempting to read a `nil` map returns the zero-value of the `ValueType`
- ***WARNING: Attempting to write to a `nil` map results in a `panic`***

```go
// Declaring a nil Map
var nilMap map[string]int
```

- We can also create a map using *Map-Literal*
  - **This does not result in a `nil` map**
  - **It has a length of `0`**
  - We can read from and write to this map

```go
// Declaring a Map with Map Literal
// This does not result in a nil map
notNilMap := map[string]int{}
```

- We can also provide key-value pairs to the map literal
- **NOTE: Comma is required for the last entry as well**

```go
// Declaring a Map with map literal and key-value pairs
valuedMap := map[string][]string{
    "Orcas": []string{"Fred", "Ralph", "Mandarin"},
    "Lions": []string{"Sarah", "Peter", "Mark"},
    "Kittens": []string{"Waldo", "Raul", "Alpha"},
}
```

- *The type of value can be anything*
- *Some restrictions apply to the type of keys*
  - Can be any *comparable* types
- **We can use `make()` to build a map**
  - Similar result to using a Map Literal
  - Allows to specify a default size
  - Will have a length of `0`
  - Can grow past the initially-specified size

```go
// Creating a Map with make()
ages := make(map[int][string], 10)
```

### Comparison With Slices

- Maps are similar to Slices
  - Dynamically grow as new entries are added
  - Specific initial size can be provided with `make()`
  - `len(map)` returns the number of `k:v` pairs in the map
  - Zero-Value is `nil`
  - Cannot be compared via `==` or `!=`
  - Can only be compared against `nil`
  - Cannot check if 2 maps are identical
- ***Use Slice for list of data to be processed in sequential order***
- ***Use Maps for organizing values using something other than sequential order***

### Map vs Hash Map

- *Map*
  - Data structure that associates one value to another
  - Can be implemented in several ways
- **In Go, *Map* is implemented as a *Hash Map* / *Hash Table***
- [More details about Map in Go](https://www.youtube.com/watch?v=Tl7mi9QmLns)
- The Go runtime implements its own Hash Algorithms for all types that are allowed to be keys

### Reading and Writing a Map

- Same principles as a `dict` in Python
  - Writing: `map[Key] = Value`
  - Reading: `map[Key]`
- **NOTE: Cannot use `:=` to assign a value to a map key**

```go
// Reading & Writing to Maps
// -------------------------
fmt.Println("Reading and Writing to Maps:")
fmt.Println("----------------------------")

totalWins := map[string]int{}

fmt.Println("Before:")
fmt.Println("totalWins =", totalWins)

// Writing to a Map
totalWins["Orcas"] = 1
totalWins["Lions"] = 2
totalWins["Kittens"]++
totalWins["Lions"] = 3

// Reading a Map
fmt.Println("After:")
fmt.Println("totalWins[\"Orcas\"] =", totalWins["Orcas"])
fmt.Println("totalWins[\"Lions\"] =", totalWins["Lions"])
fmt.Println("totalWins[\"Kittens\"] =", totalWins["Kittens"])
```

- **NOTE: Trying to access a non-existing key returns the *Zero-Value* of the value type**
- We can use `++` to increment numeric values of a map key
  - This works even if a key does not exist
  - Because a non-existing key returns a default *Zero-Value*

### *Comma-Ok* Idiom

- We sometimes need to know if a key is *really* in the map
- ***Comma-Ok* idiom allows to differentiate between**
  - ***a key that is in the map with zero-value***
  - ***a key that does not exist defaulting to a zero-value***
- Assign the result of reading a map to 2 variables: `val, ok`
- **The 2nd variable `ok` allows to check the existence of the key**
  - `true` if the key really exists
  - `false` if the key does not exist (default value)

```go
// Comma-Ok Idiom
// --------------
fmt.Println("Comma-Ok Idiom:")
fmt.Println("---------------")

greetMap := map[string]int{
    "hello": 5,
    "world": 0,
}
fmt.Println(greetMap)

// In the map with value 5
val, ok := greetMap["hello"]
fmt.Println("hello =>", val, ok)

// In the map with value 0
val, ok = greetMap["world"]
fmt.Println("world =>", val, ok)

// Not in the map, default value 0
val, ok = greetMap["hi"]
fmt.Println("hi =>", val, ok)
```

### Built-In `delete()` Function

- Allows to delete an entry from the map
  - Take a map and a key
  - Removes the key-value pair from the map
- **If the key is not found, nothing happens**
- **If the map is `nil`, nothing happens**

```go
// Deleting From Maps
// ------------------
fmt.Println("Deleting From Maps:")
fmt.Println("-------------------")

delMap := map[string]int{
    "hello": 5,
    "world": 10,
}

fmt.Println("Before Delete:", delMap)
delete(delMap, "hello")
delete(delMap, "unknownKey") // Nothing happens
fmt.Println("After Delete:", delMap)
```

### Map: Built-In `clear()` Function

- Allows to empty a map, similar with *slices*
- *Unlike with `slice`, it also resets the map's length to `0`*

```go
// Clearing a Map
// --------------
fmt.Println("Clearing a Map:")
fmt.Println("---------------")

clearMap := map[string]int{
    "hello": 5,
    "world": 10,
}

fmt.Println("Before Clear:", clearMap, "len =", len(clearMap))
clear(clearMap)
fmt.Println("After Clear:", clearMap, "len =", len(clearMap))
fmt.Println()
```

### Comparing Maps

- Use the `maps` library (Since Go 1.21)
  - Contains helper functions for working with maps

Function|Description
:-|:-
`maps.Equal()`|- Takes 2 maps *with comparable elements*<br>- Return `true` if maps are same length, all equal elements (keys and values)<br>- Else, return `false`<br>- *The map elements must be comparable*
`maps.EqualFunc()`|- Allows to pass a function argument to customize conditions for equality<br>- *The map elements do not have to be comparable*

```go
// Comparing 2 maps
// ----------------
fmt.Println("Comparing 2 maps:")
fmt.Println("-----------------")

mp1 := map[string]int{
    "hello": 5,
    "world": 5,
}
mp2 := map[string]int{
    "hello": 5,
    "world": 5,
}
mp3 := map[string]int{
    "hello": 6,
    "world": 5,
}

fmt.Println("mp1 =", mp1)
fmt.Println("mp2 =", mp2)
fmt.Println("mp3 =", mp3)
fmt.Println("maps.Equal(mp1, mp2): ", maps.Equal(mp1, mp2))
fmt.Println("maps.Equal(mp1, mp3): ", maps.Equal(mp1, mp3))
```

### Using Maps As Sets

- *Set*
  - Contains *unique* unordered elements
  - Fast-check vs checking a slice
  - Some languages contains `set` data structure (E.g. Python)
- **Go does not include a `set` data structure**
- **But we can use `map` to simulate it**
  - The keys of a map *must be unique*
  - *We can use `map[KeyType]bool` as a set*
  - *Or use `map[KeyType]string` with `""` as value*

```go
// Using map[KeyType]bool as Set
// -----------------------------
fmt.Println("Using map[KeyType]bool as Set:")
fmt.Println("------------------------------")

intSet := map[int]bool{}
intSet2 := map[int]string{}
vals := []int{5, 10, 2, 1, 5, 8, 7, 3, 9, 8, 1, 2, 10}
for _, v := range vals {
    intSet[v] = true
    intSet2[v] = ""
}

fmt.Println("vals =", vals, "len =", len(vals))
fmt.Println("intSet =", intSet, "len =", len(intSet))
fmt.Println("intSet2 =", intSet2, "len =", len(intSet2))
```

- To check for the existence of the value in the set, we can:
  - Use the value if boolean
  - Use the `Comma-Ok` idiom

```go
// Checking for element in the set
// -------------------------------
fmt.Println("Checking for element in the set:")
fmt.Println("--------------------------------")

// Using Boolean as Value
inIntSet := false
if intSet[5] {
    inIntSet = true
}
fmt.Println("5 in intSet?", inIntSet)

// Using the Comma-Ok Idiom
_, inIntSet2 := intSet2[2000]
fmt.Println("2000 in intSet2?", inIntSet2)
```

- *For advanced set manipulations, it is better to implement a Data Structure yourself or use third-party modules*
  - [golang-collections/collections](https://github.com/golang-collections/collections)
  - [deckarep/golang-set](https://github.com/deckarep/golang-set)
  - [emirpasic/Gods](https://github.com/emirpasic/Gods)
- **NOTE: Using `struct{}` might be a better value instead of `bool`**
  - Empty struct uses 0 byte
  - But can make code clumsier
  - Unless having a large set, the memory usage would not be too significant

## Structs

- Maps have limitations
  - Cannot constrain to only allow certain keys
  - All values must be of the same type
  - Maps are not ideal to pass data between functions
- **Related data should be grouped as `struct` instead**
- **NOTE: `struct` is close to the concept of `class` in OOP**
  - **However, Go is not an OOP language**
  - Go does not support `class`
  - But it support some features of OOP in a different way

### Struct Definition and Declaration

- **Defined with the keyword `type` and the type `struct`**
  - `struct` is a value-type
  - Contain a list of typed fields
  - Not comma-separated, unlike with `map`
- **Can be defined inside or outside functions, or any block-level**
- **NOTE: Capitalized names are public, Lowercase names are private**

```go
// Example of defining a struct
type Person struct {
    fName   string
    lName   string
    dob     string
    favNum  int
    isAdult bool
}
```

- Once defined, we can make use of the struct
- **Zero-Value is a struct with the Zero-Value of each field-type**
  - **Using *Struct-Literal* and *Struct With Zero-Value* are equivalent**
  - Both initialize the struct with fields with zero-values

```go
// Declaring a struct with zero-values
var fred Person

// Declaring a struct with struct literal (zero-values)
bob := Person{}
```

- **We can provide struct-literal with actual values**
  - A value for every field must be specified
  - Values are assigned in-order

```go
// Declaring a struct with struct literal and values
julia := Person{
    "julia",
    "smith",
    "1969-01-01",
    77,
    true,
}
```

- **However, it is preferable to use *named-fields***
  - Not all fields have to be specified
  - Values are assigned by field name
  - Fields can be specified in any order
  - ***Unassigned fields default to their Zero-value***

```go
// Declaring a struct with struct literal and values, with named-fields
john := Person{
    dob:   "2023-01-01",
    lName: "smith",
    fName: "john",
}
```

- ***WARNING: The 2 struct-literal styles cannot be mixed***
  - **Recommended to always use *named-fields***
  - More verbose but clearer
  - More maintainable
  - Has more advantages over *ordered*

### Struct Field Access

- **A struct field is accessed with a `.`**

```go
// Struct field access
// -------------------
fmt.Println("Struct field access:")
fmt.Println("--------------------")

bob.fName = "Bob"
fmt.Println("bob.fName =", bob.fName)
```

### Anonymous Struct

- We do not have to provide a name to a struct
- **We can assign the struct definition directly to a variable**
- **There is no difference of functionality between a *Named Struct* and an *Anonymous Struct***

```go
// Anonymous Struct
// ----------------
fmt.Println("Anonymous Struct:")
fmt.Println("-----------------")

pet := struct {
    // Struct Definition
    name  string
    kind  string
    breed string
}{
    // Struct value
    name:  "fido",
    kind:  "dog",
    breed: "golden retriever",
}

fmt.Println("pet.name:", pet.name)
fmt.Println("pet.kind:", pet.kind)
fmt.Println("pet.breed:", pet.breed)
```

#### When to Use Anonymous Struct

- For **Unmarshalling / Marshalling**
  - Translating external data (*JSON, Protocol Buffer...*) to struct or vice-versa
- For **Testing**
  - Using anonymous structs for writing table-driven tests

### Comparing And Converting Structs

- **Comparability depends on the fields**
  - If the fields are entirely *comparable*, then the struct is *comparable*
  - **`slice` and `map` fields are not *comparable***
- **There is no magic method to redefine `==` and `!=`**
  - We can write our own functions
- **Go does not allow comparison between structs of different types**
  - *Struct type conversion is possible*
  - **But only if the fields of both structs have the same name, order, and types**

```go
type firstPerson struct {
    name string
    age  int
}
type secondPerson struct {
    name string
    age  int
}
type thirdPerson struct {
    age  int
    name string
}
type fourthPerson struct {
    fName string
    age  int
}
type fifthPerson struct {
    name string
    age  int
    dob  string
}
```

- **We can use type-conversion from `firstPerson` instance to `secondPerson`**
- We *cannot* do `firstPerson == secondPerson` because they are different types
- We *cannot* use type-conversion from `firstPerson` instance to `thirdPerson` because the fields have different order
- We *cannot* use type-conversion from `firstPerson` instance to `fourthPerson` because the field names are different
- We *cannot* use type-conversion from `firstPerson` instance to `fifthPerson` because not all columns match
- With *Anonymous Structs*
  - *Anonymous Structs are compatible with any other structs*
  - **But as long as the fields of both structs have the same name, order, and types**
  - Can be compared without type-conversion as well with the same rule

```go
// Comparing Against Anonymous Struct
// ----------------------------------
fmt.Println("Comparing Against Anonymous Struct:")
fmt.Println("-----------------------------------")

type firstPerson struct {
    name string
    age  int
}

f := firstPerson{
    name: "Bob",
    age:  50,
}

g := struct {
    name string
    age  int
}{
    name: "Bob",
    age:  50,
}

// Comparing
fmt.Println("f == g?", f == g)

// Using with each other as equivalent
f = g
fmt.Println("f =", f)
```
