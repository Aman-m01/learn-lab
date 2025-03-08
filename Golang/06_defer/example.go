package main

import "fmt"

func main() {
	fmt.Println("Counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Done !")
}

/*
-> Execution of Program
	1. Print "Counting" → This is printed immediately.
	2. Loop Execution (for i := 0; i < 10; i++)
		->  On each iteration, defer fmt.Println(i) is executed.
		->  Since defer statements are stacked, the numbers are pushed in order: 0, 1, 2, ..., 9.

	3. Print "Done !" → This is printed before any of the deferred statements execute.

	4. Deferred Statements Execute (LIFO Order)
		-> The values are printed in reverse order because they were pushed onto the stack.
		-> The output starts from 9 and goes down to 0.

output:
		Counting
		Done !
		9
		8
		7
		6
		5
		4
		3
		2
		1
		0
*/
