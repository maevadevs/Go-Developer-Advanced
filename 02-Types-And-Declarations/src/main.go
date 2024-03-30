package main

import "fmt"

// This is the main entry of the application.
func main() {
    fmt.Println("Hello world!")
}

// FOR WINDOWS:
//  To Run:                 make run-win
//                          go run 02-Types-And-Declarations\src\main.go
//  To Build:               make build-win
//                          go build -o 02-Types-And-Declarations\bin\Types.exe 02-Types-And-Declarations\src\main.go
//  To Run after Build:     .\bin\Types.exe
//                          .\02-Types-And-Declarations\bin\Types.exe
//  Try Build + Run:        make try-win
//                          go build -o 02-Types-And-Declarations\bin\Types.exe 02-Types-And-Declarations\src\main.go && .\02-Types-And-Declarations\bin\Types.exe && rm .\02-Types-And-Declarations\bin\Types.exe

// FOR LINUX:
//  To Run:                 make run
//                          go run 02-Types-And-Declarations/src/main.go
//  To Build:               make build
//                          go build -o 02-Types-And-Declarations/bin/Types 02-Types-And-Declarations/src/main.go
//  To Run after Build:     ./bin/Types
//                          ./02-Types-And-Declarations/bin/Types
//  Try Build + Run:        make try
//                          go build -o 02-Types-And-Declarations/bin/Types 02-Types-And-Declarations/src/main.go && ./02-Types-And-Declarations/bin/Types && rm ./02-Types-And-Declarations/bin/Types
