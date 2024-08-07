# Go Developer Advanced

---

These are in-depth notes from learning Go.

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
   - Staying Up-To-Date
2. [Primitive Types and Declarations](./02-Primitive-Types-And-Declarations/)
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
     - Built-In Functions: `len()`, `append()`
     - Length vs Capacity
     - Built-In Functions: `make()`, `clear()`
     - Choosing Slice Declaration
     - Subslicing Slice
     - Built-In Function: `copy()`
     - Converting Array To Slice: `copy()`
     - Converting Slice to Array: Type-Conversion
   - Strings, Runes, Bytes
     - UTF-8
   - Maps
     - Comparison With Slices
     - Map vs Hashmap
     - Reading and Writing a Map
     - *Comma-Ok* Idiom
     - Built-In Functions: `delete()`, `clear()`
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
      - C-Style `for`
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
      - Choosing Between `if-else` and `switch`
    - `goto`
      - Use of `goto` in Go
5. [Functions](./05-Functions/)
    - Declaring and Calling Function
    - Simulating Named and Optional Parameters
    - Variadic Function and Slices
    - Multiple Return Values
      - Ignoring Returned Values
      - Named Return Values
      - Blank Return
    - Functions Are Values
      - Function Type Declarations
      - Anonymous Functions
    - Closures
      - Benefits of Closures
    - Passing Functions As Arguments
    - Returning Functions From Functions
    - `defer`
    - Call By Value
6. [Pointers](./06-Pointers/)
7. Types, Methods, Interfaces
8. Generics
9. Errors
10. Modules, Packages, Imports
11. Go Tooling
12. Concurrency
13. Standard Library
14. Context
15. Tests
16. Reflect, Unsafe, Cgo
