package main

import "fmt"

var c, python, java bool // these global variables are initialized to their zero values and must be declared using var, because in GO a variable is never in a uninitialized state. this is the only way, we defined variables outside the functions body.

func main() {
	var k string // var k string declares k with the zero value "", then it's assigned "Aman".
	k = "Aman"
	i := 100 // The := short declaration syntax is used inside functions. It declares and initializes i in one step.

	fmt.Println(i, c, python, java, k) // 0 false false false
}

/*
> Datatypes
- Go has several built-in data types categorized into basic types, composite types, and reference types.

1. Basic datatypes
  - Integar
  - unsigned integar
  - floating point
  - Complex number
  - Boolean
  - String

2. Composite data types
  - Array
  - Slice
  - Map
  - Struct
  - Pointer

3. Reference type
  - Functions
  - Interface
  - Channels
*/

/*
> NOTE
- Explicit types > The type of the variable is explicitly mentioned during declaration, This is useful when you want to ensure type safety and prevent unintended type conversions.
                     ex:   var x int = 10  // Explicitly stating `x` is an `int`

- Implicit types > The type is inferred automatically by the Go compiler based on the assigned value. This is done using the := syntax and Only allowed inside functions.
                   ex: y := 3.14    // Go infers `y` as a `float64`
                       name := "Go" // Go infers `name` as a `string`
*/
