package main

import "fmt"

func main() {
	fmt.Println("Start")

	defer fmt.Println("First Deferred")  // Deferred 2nd (executes 2nd-last)
	defer fmt.Println("Second Deferred") // Deferred 1st (executes last)

	fmt.Println("End of function")
	// The defer statements are stacked. When the function completes execution, the deferred functions are popped and executed in reverse order of how they were deferred.
}

/*
-> Defer in Golang
  -> The defer statement in Go is a powerful feature that ensures a function or expression is executed only after the surrounding function completes, regardless of whether it finishes successfully or encounters an error. Think of it as a way to schedule cleanup tasks that must run at the end of a function, like closing files, unlocking mutexes, or recovering from panics

-> How defer Works
  -> When a defer statement is encountered in a function, the expression or function call following defer is pushed onto a stack. The function then continues executing normally. Once the function is about to return (either normally or due to an error), all the deferred function calls are executed in Last In, First Out (LIFO) order -> meaning the most recent defer statement runs first.

-> Key Uses of defer
 -> 1. Resource Management (Closing Files, Unlocking Mutexes, Freeing Memory) > Without defer, you’d have to remember to manually close the file at every possible return point. defer ensures the file is closed no matter how the function exits (error or success).

 -> 2.  Panic Recovery (Handling Unexpected Errors) > A panic in Go stops execution unless it is recovered. defer is useful in panic handling to ensure proper cleanup before the program crashes.

 -> 3. Stacked Defers (Handling Multiple Cleanup Actions) > Multiple defer statements stack up in LIFO order.

 -> 4. Code Simplification (Ensuring Cleanup in One Place) > Instead of manually handling cleanup in multiple return paths, defer lets you handle it once at the beginning.

*/
