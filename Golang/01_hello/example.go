package main

import (
	"fmt"
	"math"
	"math/rand" // we use this package to generate random numbers and typical functionalities like that.
)

func main() {
	fmt.Println("number is : ", rand.Intn(10)) // here we print the random number b/w 1 to 10
	fmt.Println(math.Pi)                       // 3.141592653589793
}

/*

- In this program, we import the "math" package to access mathematical constants and functions. When accessing a variable from another package, it must be exported (i.e., its name must start with a capital letter).
- In Go, identifiers that start with a lowercase letter are unexported and cannot be accessed outside their package. Only variables, functions, and types whose names start with a capital letter are accessible from outside their package.

- By default, Go restricts access to variables, functions, and types within their defining package unless explicitly exported.

*/
