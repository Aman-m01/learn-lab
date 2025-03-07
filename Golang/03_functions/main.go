package main

import "fmt"

// basic functions
func hello() {
	fmt.Println("Hello, GO!")
}

// function with params
func add(x, y int) int {
	return x + y
}
func greet(name string) string {
	return "Hello " + name
}

// multiple return values
func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

// Function with named return values
func rectangle(l, w int) (area int, perimeter int) {
	area = l * w
	perimeter = 2 * (l + w)
	return // No need to explicitly return variables
}

func swap(x, y string) (string, string) {
	return y, x
}

// Variadic Functions (Multiple Arguments) > Go supports variadic functions, which take a variable number of arguments.
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Defining and calling an anonymous function
var result = func(a, b int) int {
	return a + b
}(4, 6)

// Function that accepts another function as an argument
func applyOperation(a, b int, op func(int, int) int) int {
	return op(a, b)
}
func subtract(x, y int) int {
	return x - y
}

func main() {
	hello()
	fmt.Println(add(2, 3))
	fmt.Println(greet("Aman"))

	q, r := divide(10, 3)
	fmt.Println("Quotient:", q, "Remainder:", r)

	a, p := rectangle(5, 10)
	fmt.Println("Area:", a, "Perimeter:", p)

	fmt.Println(swap("Hello", "World"))
	fmt.Println(sum(12, 13, 14))
	fmt.Println(sum(12, 13, 14, 12, 1343, 45, 566))

	fmt.Println("Sum:", result)

	applyOperationResult := applyOperation(10, 3, subtract)
	fmt.Println("Result:", applyOperationResult)
}

/*
- In Go, a function is a self-contained block of code that performs a specific task. Functions allow code reuse, modularity, and better maintainability.
*/
