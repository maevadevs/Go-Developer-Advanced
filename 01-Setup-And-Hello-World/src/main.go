package main

import "fmt"

// This is the main entry of the application.
func main() {
	fmt.Println("Hello world!")
	fmt.Printf("Hello again %s!\n", "world")
}

// FOR WINDOWS:
//  To run:                 go run 01-Setup-And-Hello-World\src\main.go
//  To compile:             go build -o 01-Setup-And-Hello-World\bin\Hello-World.exe 01-Setup-And-Hello-World\src\main.go
//  To run after compile:   .\01-Setup-And-Hello-World\bin\Hello-World.exe
//  Compile + Run:          go build -o 01-Setup-And-Hello-World\bin\Hello-World.exe 01-Setup-And-Hello-World\src\main.go && .\01-Setup-And-Hello-World\bin\Hello-World.exe

// FOR LINUX:
//  To run:                 go run 01-Setup-And-Hello-World/src/main.go
//  To compile:             go build -o 01-Setup-And-Hello-World/bin/Hello-World 01-Setup-And-Hello-World/src/main.go
//  To run after compile:   ./01-Setup-And-Hello-World/bin/Hello-World
//  Compile + Run:          go build -o 01-Setup-And-Hello-World/bin/Hello-World 01-Setup-And-Hello-World/src/main.go && ./01-Setup-And-Hello-World/bin/Hello-World
