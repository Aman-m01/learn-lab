package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

/*
> Loops
- Golang provides a simple yet powerful way to handle loops. Unlike other languages that have multiple looping constructs (for, while, do-while), Go has only one looping construct: for. However, it can be used in various ways to achieve the functionality of while and do-while loops as well.
*/
