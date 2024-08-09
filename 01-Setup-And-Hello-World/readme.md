# Setup and Hello World

---

- [Installing Go](#installing-go)
  - [Detailed Installation for Linux](#detailed-installation-for-linux)
    - [Uninstalling Go from Linux](#uninstalling-go-from-linux)
  - [Detailed Installation for Windows](#detailed-installation-for-windows)
    - [Uninstalling Go from Windows](#uninstalling-go-from-windows)
  - [Detailed Installation for Mac](#detailed-installation-for-mac)
    - [Uninstalling Go from Mac](#uninstalling-go-from-mac)
- [Commonly Used Go Commands](#commonly-used-go-commands)
- [Hello World](#hello-world)
  - [Go Module](#go-module)
  - [`main.go`](#maingo)
    - [Package Declaration](#package-declaration)
    - [Import Declaration](#import-declaration)
  - [`main` Function](#main-function)
  - [Build and Execution](#build-and-execution)
- [`go fmt`](#go-fmt)
  - [Automatic Semicolon Insertion Rule](#automatic-semicolon-insertion-rule)
- [`go vet`](#go-vet)
- [Go Dev Tools](#go-dev-tools)
  - [VS Code Setup](#vs-code-setup)
  - [Makefiles For Build Automation](#makefiles-for-build-automation)
- [Go Compatibility Promise](#go-compatibility-promise)
- [Staying Up-To-Date](#staying-up-to-date)

---

- Every programming language needs a development environment
- We will setup a development environment for Go

## Installing Go

- We only need to install Go: No additional tools
  - [Check here](https://go.dev/doc/install) for installation instructions
  - [Check here](https://go.dev/dl/) for all available releases
- **Go programs compile to a single native binary**
  - No VM needed to run
  - No additional software required to run
  - Go apps can also be [packaged inside Docker images](https://blog.baeke.info/2021/03/28/distroless-or-scratch-for-go-apps/)

### Detailed Installation for Linux

- Download latest `.tar.gz` file from [official website](https://go.dev/dl/) into your *Downloads* directory

```sh
cd ~/Downloads
wget https://go.dev/dl/{{go_version}}.tar.gz
```

- Remove any previous installation, if any
- Extract the new archive into installation location

```sh
sudo rm -rf /usr/local/go
tar -C /usr/local -xzf {{go-version}}.tar.gz
```

- Add `/usr/local/go/bin` to the `PATH` environment variable

```sh
export PATH=$PATH:/usr/local/go/bin
source $HOME/.bash_profile
```

- Verify the installation

```sh
go version
which go
```

#### Uninstalling Go from Linux

```sh
sudo rm -rf /usr/local/go
sudo rm /etc/paths.d/go
```

- Remove the Go paths from these files, if any
  - `/etc/profile`
  - `$HOME/.bash_profile`
  - `$HOME/.profile`
  - `$HOME/.bashrc`

### Detailed Installation for Windows

- Download latest `.msi` file from [official website](https://go.dev/dl/) into *Downloads* directory
- Open the `.msi` file and follow installation prompts
- Close and reopen any open command prompts or shell prompts
- Verify the installation

```sh
go version
where go
```

#### Uninstalling Go from Windows

- Remove from *Control Panel*
- Or remove with `msiexec` via the `.msi` installation file

```ps1
msiexec /x go{{version}}.windows-{{cpu_arch}}.msi /q
```

### Detailed Installation for Mac

- Download latest `.pkg` file from [official website](https://go.dev/dl/) into *Downloads* directory
- Open the `.pkg` file and follow installation prompts
- Close and reopen any open command prompts or shell prompts
- Verify the installation

```sh
go version
which go
```

#### Uninstalling Go from Mac

```sh
sudo rm -rf /usr/local/go
sudo rm /etc/paths.d/go
```

- Remove the Go paths from these files, if any
  - `/etc/profile`
  - `$HOME/.bash_profile`
  - `$HOME/.profile`
  - `$HOME/.bashrc`

## Commonly Used Go Commands

Command | Action
:-|:-
`go version`|Check installed Go **version**
`go build`|**Compiles** multiple Go source code files into executable binaries
`go run`|**Compiles and executes** multiple Go source code files (build & execute)<br>*But does not produce an actual executable*
`go fmt`|When ran in a directory with `.go` files, **formats** all the code in each file in the directory
`go vet`|When ran in a directory with `.go` files, **scans** for common coding mistakes
`go install`|**Compiles** and **installs** a package and dependencies
`go get`|**Download** raw source code of someone else's package as dependency
`go test`|Runs any **tests** associated with the current projects
`go mod`|For **dependency management** and **module management**
`go mod tidy`|Verify existing module dependencies<br>Adds missing and removes unused modules
`go bug`|Start a bug report
`go clean`|Remove object files and cached files
`go doc`|Show documentation for package or symbol
`go env`|Print Go environment information
`go fix`|Update packages to use new APIs
`go generate`|Generate Go files by processing source
`go list`|List packages or modules
`go work`|Workspace maintenance
`go tool`|Run specified go tool

- [Additional Go Commands](https://pkg.go.dev/cmd/go)

## Hello World

- **You are free to organize your project as you see fit**
- Here is a suggested project structure

```tree
Go-Module/
├── bin/
    ├── debug/
    ├── release/
├── src/
|   ├── main.go
├── tests/
├── go.mod
├── makefile
└── readme.md
```

### Go Module

- **A Go project is called a *Go Module***
  - A *Go Module* is a directory with a `go.mod` file in it
  - `go.mod` can be auto-generated with the following command in the module directory

```sh
go mod init unique/module/path/typically/github
```

- A `go.mod` file declares:
  - The name of the module and its unique path
  - Minimum supported Go version
  - Any additional module dependencies
- **The `go.mod` file should not be edited directly**
  - Use `go` command to manage its contents
  - `go mod` creates it
  - `go get` adds dependencies
  - `go mod tidy` verifies dependencies
- The module is then defined in `go.mod` file as follow

```sh
module unique/module/path/typically/github

go 1.22.6
```

### `main.go`

- Define *executable code* in `main.go`

```go
package main

import "fmt"

// This is the main entry of the application.
func main() {
    fmt.Println("Hello World!")
}
```

#### Package Declaration

- Within a Go Module, codes are organized into one or more *packages*
  - **`main` package contains *executable* code**
  - **Any other package names are *reusable libraries***
  - **A package is essentially a directory (except for `main`)**

#### Import Declaration

- Allows to import external packages referenced in the file
  - **Go `import` only imports whole packages**
  - We cannot limit the imports to specific functions or atomic elements
- `fmt` is from the standard library so it is already available
  - No need to use `go get`
  - For 3rd-Party libraries, we use `go get` first

### `main` Function

- **All executable Go programs start execution from the `main()` function in the `main` package**

### Build and Execution

Command | Action
:-|:-
`go build`|**Compiles** multiple Go source code files into executable binaries
`go run`|**Compiles and executes** multiple Go source code files (build & execute)<br>*But does not produce an actual executable*

- **To run**

```sh
go run Go-Module-Name/src/main.go
```

- **To build/compile into binaries**

```sh
go build -o Go-Module-Name/bin/Hello-World Go-Module-Name/src/main.go
```

- **To run after compile**

```sh
./Go-Module-Name/bin/Hello-World
```

- **To build/compile into binaries AND run**

```sh
go build -o Go-Module-Name/bin/Hello-World Go-Module-Name/src/main.go && ./Go-Module-Name/bin/Hello-World
```

- **NOTE: By default, `go build` uses the name of the module  as the name of the binary output**
  - To specify a new name or path, we use the `-o` flag

## `go fmt`

- Go allows to write code efficiently
- Simple syntax and fast compile
- **Go also has an *enforced* standard code-formatting**
  - This simplifies the compiler
  - Makes easier collaborations
  - More productive
  - Avoid unnecessary arguments over styles and formatting
  - E.g. It is a syntax error if the opening brace is not on the same line as the declaration or block command
  - Also allows some clever tools to generate code
- **`go fmt` command can be run on all files in the project**
  - Only run this from within a module directory

```sh
# Fix formatting everywhere
go fmt ./...
```

- **NOTE: Remember to always run `go fmt` before compiling/building and pushing to version control**
  - This is one of the first step to ensure quality code
  - It is a good practice to make a separate commit that only does `go fmt ./...`
  - Helps to avoid combining logic changes and formatting changes
- **NOTE: In VS Code, it is possible to activate auto-formatting by `gofmt` via the Go extension settings**

### Automatic Semicolon Insertion Rule

- Go normally requires `;` at the end of every statement
- **However, Go developers should never put those `;` manually**
- The Go compiler inserts them automatically [by following a rule](https://go.dev/doc/effective_go#semicolons)
- The short insertion rule is: ***If the newline comes after a token that could end a statement, insert a semicolon***
  - That is, *a semicolon is inserted at the newline if the last token before the newline is either:*
    - *An identifier,*
    - *A basic literal,*
    - *Or any of `break`, `continue`, `fallthrough`, `return`, `--`, `++`, `)`, `}`*
- **Because of this rule, braces must be placed strategically**
  - **Following the correct code-formatting is required**
- *Automatic Semicolon Insertion* rule results in
  - Go compiler running simpler and faster
  - Enforced coding-style standard

```go
// The following breaks because...
func main() // ...a semicolon is automatically added here
{
    fmt.Println("This breaks!")
}

// However, the following works because..
func main() { // ...no automatic semicolon insertion here
    fmt.Println("This works!")
}
```

## `go vet`

- It is possible for code to be syntactically valid yet likely incorrect
- `go vet` can detect some of these errors
- **Run it in a directory with `.go` files to scan for common coding mistakes**

```go
package main

import "fmt"

// This is the main entry of the application.
func main() {
    fmt.Printf("Hello %s!\n")
}
```

- Running `go vet` on the above code capture the error
  - No value was specified for the placeholder of `Printf()`
  - The code will still compile and run
  - However, it is not correct
- **NOTE: Always run `go vet` before compiling/building and pushing to version control**
  - This is one of the first step to ensure quality code
- **NOTE: In VS Code, it is possible to activate *Vet On Save* with `go vet` via the Go extension settings**
- *NOTE:*
  - *`fmt.Printf()` allows Print Formatting*
  - *`fmt.Sprintf()` allows String Interpolation*

```go
package main

import "fmt"

// This is the main entry of the application.
func main() {
    fmt.Printf("Hello %s!\n", "world")
}
```

- **NOTE: `go vet` can capture many bugs, but it cannot capture all bugs**
  - We can use additional 3rd-party quality-code tools
  - Follow additional advices from [*Effective Go*](https://go.dev/doc/effective_go) and [*Go Wiki: Go Code Review Comments*](https://go.dev/wiki/CodeReviewComments)

## Go Dev Tools

- A simple Text Editor and Go commands can be used to write small programs
- However, better IDEs or Advanced Editors are necessary for larger projects
  - Auto-format on save (`go fmt`)
  - Auto-error-checking (`go vet`)
  - Code completion
  - Type checking
  - Code navigation
  - Integrated debugging
- There are [multiple options for Go IDEs and Advanced Editors](https://go.dev/wiki/IDEsAndTextEditorPlugins)
  - [VSCode](https://code.visualstudio.com/) + [Go Extension](https://marketplace.visualstudio.com/items?itemName=golang.Go)
  - [JetBrains GoLand](https://www.jetbrains.com/go)
  - [Vim](https://www.vim.org/)
  - [Neovim](https://neovim.io/)
  - [GNU Emacs](https://www.gnu.org/software/emacs/)
- **Here, we are using VS Code**

### VS Code Setup

- Download [VS Code](https://code.visualstudio.com/download)
- Open any `.go` file with VS Code for the first time
  - The Go-related extensions will automatically install
  - If not, search the `Go` extension and install
- The installation includes [all the tools listed here](https://github.com/golang/vscode-go/wiki/tools), which includes
  - A language server `gopls`
  - A debugger `dlv`
  - Unit Test `gotests`
  - Linter `staticcheck`
- Tools can be updated via the VS Code command

```sh
Go: Install/Update Tools
```

### Makefiles For Build Automation

- An IDE is nice to use but hard to automate
- **Automated build is essential**
  - Requiring this tooling is good software engineering practice
  - Use scripts to specify the build steps
- **Go developers can use `make` as the principal build solution**
  - `make` has been used to build programs since 1976
  - Specify a set of operations for building the program
  - Specify the order in which to perform the steps

```sh
sudo apt-get install build-essential
```

- To use `make`
  - Create a `makefile` in the project/module directory
  - Add the content for building the project
- **NOTE: Make sure to use *tabs* for the whitespaces for indentation in a makefile**
- **NOTE: In VSCode, it is helpful to also install a Makefile-support extensions**
  - The *Makefile Tools* by Microsoft is a good extension to have

```makefile
# NOTE: Make sure all indentations use tabs

# DEFAULT_GOAL specifies the default target
# This is run when no target is provided during the call
.DEFAULT_GOAL := try

# Target definitions
# .PHONY helps avoid possible name-collisions with other directory or file names on the computer
.PHONY: fmt vet build build-release run run-release try

# Target
fmt:
    # Task: Format all source files
    cd src
    go fmt ./...
    cd ..

# Target
vet: fmt
    # Task: Verify any possible errors
    cd src
    go vet ./...
    cd ..

# Target
build: vet
    # Task: Build module
    go build -o bin/debug/Hello-World src/main.go

# Target
build-release: vet
    # Task: Build module for release, strip symbols
    go build -ldflags="-s -w" -o bin/release/Hello-World src/main.go

# Target
run: build
    # Task: Build module then run
    bin/debug/Hello-World

# Target
run-release: build-release
    # Task: Build module for release then run
    bin/release/Hello-World

# Target
try: vet
    # Task: Build module, run, then remove built binary
    if test -f bin/debug/Hello-World-Temp; then \
        rm -f bin/debug/Hello-World-Temp; \
    fi
    go build -o bin/debug/Hello-World-Temp src/main.go
    bin/debug/Hello-World-Temp
    rm -f bin/debug/Hello-World-Temp
```

- *Target*
  - Each possible operation
  - *`DEFAULT_GOAL` specifies the default target*
    - This is run when no target is provided during the call
- *Target Definitions*
  - **The word before `:` is the name of the target**
  - **After `:` is a requirement target that must run before running the current target**
- `.PHONY` helps avoid possible name-collisions with other directory or file names on the computer

Command|Description
:-|:-
`make`|Runs the *target* specified under `DEFAULT_GOAL`
`make <target>`|Runs the requirements and steps specific for that *target*

- **Ensuring that *formatting* and *vetting* always happen before *building* is a great improvement in the build process**
- Drawbacks of Makefiles
  - They are picky: *Must indent with tabs only*
  - Not supported out-of-the-box on Windows
    - *Must install `make` first*: `choco install make`
- [Additional resources for learning Makefiles](https://makefiletutorial.com/)

## Go Compatibility Promise

- Go is periodically updated
  - New release roughly occurs every 6 months
  - Patches and security fixes are as-needed
  - Releases tend to be incremental than expansive
- The Go team plans to avoid breaking changes by following the [*Go Compatibility Promise*](https://go.dev/doc/go1compat)
  - **For any version with *Major 1*, there is no expected breaking changes to the language or standard library**
  - **Unless a breaking change is required for bug or security fix**
  - **But this guarantee does not apply to [Go commands](https://pkg.go.dev/cmd/go)**

## Staying Up-To-Date

- **Updating Go environment does not affect previously-compiled programs**
  - Go compiles to standalone native binary
  - There is no dependency on any VMs
  - The version of Go does not affect the runtime of compiled programs

System|Option
:-|:-
Windows|`chocolatey` can be used to install and update Go
Mac|`brew` can be used to install and update Go
Linux|- Move the current installation in a backup location<br>- Download and unpack the new installation<br>- If all is good, delete the backup
Others|Follow the installer's option for removing and re-installing
