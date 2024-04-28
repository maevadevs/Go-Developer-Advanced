# Go Developer Advanced

---

These are notes from reading *Learning Go* book.

- [Book Repo](https://github.com/learning-go-book-2e)

1. [Setup and Hello World](./01-Setup-And-Hello-World/)
   - Installation: Linux, Windows, Mac
   - Uninstallation: Linux, Windows, Mac
   - Commonly Used Go Commands
   - Hello World
   - Go Module
   - `main.go`
   - Build and Execution
   - `go fmt`
   - `go vet`
   - Go Dev Tools
   - VS Code Setup
   - `make` and Build Automation
   - Go Compatibility Promise
2. [Types and Declarations](./02-Types-And-Declarations/)
   - Predeclared Types
   - Zero-Values
   - Literals: Integer, Float, Rune, String, Default Type
   - Boolean Types
   - Numeric Types: Integer, Float, Complex
   - Strings and Runes
   - Explicit Type Conversion Required
   - Literals are Untyped
   - `var` vs `:=`
     - Multiple Declarations
     - Walrus Shortcut `:=`
   - Using `const`
     - Typed Const vs Untyped Const
   - Unused Variables
   - Naming Variables And Constants
3. [Composite Types](./03-Composite-Types/)
   - Arrays
     - Array Declarations
     - Comparing Arrays
     - Multidimensional Array
     - Array Access
     - Getting Array Length
     - Arrays Limitations
     - Arrays Purpose
   - Slices
     - Working With Slices
     - Multidimensional Slice
     - Slice Access
     - Differences With Arrays
     - Slices Helper Functions From `slices`
     - Built-In `len()` Function
     - Built-In `append()` Function
     - Length vs Capacity
     - Built-In `make()` Function
     - Built-In `clear()` Function
     - Choosing Slice Declaration
     - Subslicing Slice
     - Built-In `copy()` Function
     - Converting Array To Slice: `copy()`
     - Converting Slice to Array: Type-Conversion
   - Strings, Runes, Bytes
     - UTF-8
   - Maps
     - Comparison With Slices
     - Map vs Hashmap
     - Reading and Writing a Map
     - *Comma-Ok* Idiom
     - Built-In `delete()` Function
     - Map: Built-In `clear()` Function
     - Comparing Maps
     - Using Maps As Sets
   - Structs
     - Struct Definition and Declaration
     - Struct Field Access
     - Anonymous Struct
     - Comparing And Converting Structs
4. [Blocks, Shadows, and Control Structures](./04-Blocks-Shadows-Control-Structures/)
    - Blocks
    - Shadowing
    - Universe Block
    - `if`
    - `for`
      - Complete C-Style `for`
      - Condition-Only `for` (`while`-Style)
      - Infinite `for`
        - `break` and `continue`
      - `for-range` Statement
        - Iterating Over Maps
        - Iterating Over Strings
        - The `for-range` Value Is A Copy
      - Labeling `for` Statements
      - Choosing The Right `for` Statement
    - `switch`
      - Blank `switch`
