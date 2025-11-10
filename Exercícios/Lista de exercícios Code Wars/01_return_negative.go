// In this simple assignment you are given a number and have to make it negative. But maybe the number is already negative?

// Examples

// make_negative(1);  # return -1
// make_negative(-5); # return -5
// make_negative(0);  # return 0

// Notes
// The number can be negative already, in which case no change is required.
// Zero (0) is not checked for any specific sign. Negative zeros make no mathematical sense.

//go:build ignore

package main

import (
	"fmt"
)

func MakeNegative(num int) int {
	if num <= 0 {
		return num
	} 

	return num * -1
}

func main() {
	var num int
	fmt.Print("Digite um número: ")
	fmt.Scan(&num)

	fmt.Println(MakeNegative(num))
}
