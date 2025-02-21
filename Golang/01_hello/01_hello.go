package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
}

/*
> Note
- In GO, every file belongs to a package, package main is a special package in GO. It tells the compiler that this file is an executable program. Only package named main can produce a standalone executable when compiled.
- Packages are a way to organize our code, divide our code in a different-different files.

- import "fmt" > it imports a package form the standard library, all the package from the standard library are already defined in the Go library, we dont need to explicitely add it from the other sources. "fmt" is built in GO package that provides functions for formatted input/output operations.

- main() is the special function in GO, it is the entry point of the program. every go program must have a main() functions inside a main package to execute.

- whenever you access functions, variables, interfaces, structs and any data types form a package, it always be in capital from the first letter.
_________________________________________________________________________________________________________________________
Q1. why did the china language is written in the hello world program(in official GO Docs) ?
 - Go has built-in support for Unicode, meaning it can handle characters from almost all languages, including Chinese, Japanese, Arabic, etc.
 - The phrase 你好, 世界 means "Hello, World!" in Chinese, and using it highlights Go’s ability to handle multilingual text seamlessly.

_________________________________________________________________________________________________________________________
Q2. go build vs go run
- The go build command compiles a GO program into an executable binary, which can run on the any supported OS.
- One of GO's powerful features is the ability to generate a standalone executable that runs on different operating systems without needing additional dependencies. This means that a Go program can be compiled once and run on any system, making it truly cross-platform.

- go run Compiles and runs the program immediately without creating a permanent binary.
__________________________________________________________________________________________________________________________

Q3. Building an executable for another OS
- by default, `go build` compiles for the current OS and architecture, however we can explicitly specify the target OS and architecture.
ex-
    SET GOOS=darwin   // target OS
    SET GOARCH=amd64  // architecture is 64-bit
	go build main.go
- The generated executable can then be transferred and run on a macOS system.
______________________________________________________________________________________________________

Q4. How `go run` works?
 - `go run` goes through the compilation process, compiling all dependencies and packages. However, it does not store the compiled binary permanently; it creates a temporary binary somewhere in the system and executes it. However, it does not store the compiled binary permanently; it creates a temporary binary somewhere in the system and executes it. This makes go run relatively faster than go build, especially for repeated executions.

Q5. How `go build` works?
- `go build` compiles, links, and optimizes the program, producing an executable file. it does not cache results. Every time you run go build, it starts from scratch, making it slightly slower compared to `go run`.
_____________________________________________________________________________________________________________________

Q6. what is go.mod file?
- The go.mod file manage dependencies in a Go project. It defines the module’s name and tracks its required dependencies.
- It defines the module’s name, tracks required dependencies, and ensures reproducible builds by specifying exact versions of external libraries. The file typically includes the module path, the Go version, and a list of dependencies with their versions.

*/
