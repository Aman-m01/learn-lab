package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}

// This program is using the packages with import paths "fmt" and "math/rand".
// By convention, the package name is the same as the last element of the import path. For instance, the "math/rand" package comprises files that begin with the statement package rand.

/*
> Packages
- A package in Go is a directory that contains Go source files with a common purpose. Every Go program consists of packages, and the starting point for execution is always the main package.

1. Executable packages:
   - Contains a main functions
   - Used to create executable applications
   - The Go compiler will compile it into a binary.

2. Library (utility) packages
   - It doesn't have a main functions
   - Used to provide reusable functionalities
   - These packages can be imported into other programs

> Note
- Package names should be short, lowercase, and descriptive.
- Package names should match the directory name they reside in
- The name should be singular unless a plural makes sense and also avoid using underscore and camelCase
- Avoid generic names like utils or helpers.

- Variables and functions declared in a package are visible only within that package unless they are exported (start with an uppercase letter).
- Export functions/variables using uppercase letters.
- Use init() for package initialization
- Use go mod for dependency management.
*/
