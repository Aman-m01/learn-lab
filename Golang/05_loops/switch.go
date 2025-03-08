// 1
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")

	today := time.Now().Weekday()

	switch time.Saturday {
	case today:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away")
	}
}

/*
//2
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Go runs on:")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s\n", os)
	}
}
*/

/*
//3
package main

import (
	"fmt"
	"time"
)
func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")

	}

}
*/
