package main

import "fmt"

// This is the main entry of the application.
func main() {
    fmt.Println("Hello world!")
}

// FOR WINDOWS:
//  To Run:                 make run-win
//                          go run Functions\src\main.go
//  To Build:               make build-win
//                          go build -o Functions\bin\Functions.exe Functions\src\main.go
//  To Run after Build:     .\bin\Functions.exe
//                          .\Functions\bin\Functions.exe
//  Try Build + Run:        make try-win
//                          go build -o Functions\bin\Functions.exe Functions\src\main.go && .\Functions\bin\Functions.exe && rm .\Functions\bin\Functions.exe

// FOR LINUX:
//  To Run:                 make run
//                          go run Functions/src/main.go
//  To Build:               make build
//                          go build -o Functions/bin/Functions Functions/src/main.go
//  To Run after Build:     ./bin/Functions
//                          ./Functions/bin/Functions
//  Try Build + Run:        make try
//                          go build -o Functions/bin/Functions Functions/src/main.go && ./Functions/bin/Functions && rm ./Functions/bin/Functions


