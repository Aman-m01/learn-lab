package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T, Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T, Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T, Value: %v\n", z, z)
}

/*
%v - %v prints a value in its default format, It works for any data type.

%T - Prints the type of variables

%d - Prints an integer in decimal format (base 10)

%b - Prints an integer in binary format.

%o - Prints an integer in octal format.

%t - Prints a boolean (true or false)

%c - Prints an ASCII character from an integer.

%s - Prints a String

%q - Print a string with the double quotes

%p - Prints the memory address of a variable.

%% - Prints % literally.

*/
