// Create a function that takes an integer as an argument and returns "Even" for even numbers or "Odd" for odd numbers.

//go:build ignore

package main

import "fmt"

func EvenOrOdd(num int) string {
	if num % 2 == 0 {
		return "Even"
	}
	return "Odd"
}

func main() {
	fmt.Println(EvenOrOdd(25))
}
