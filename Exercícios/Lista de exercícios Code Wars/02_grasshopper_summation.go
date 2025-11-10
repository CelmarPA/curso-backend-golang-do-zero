// Summation
// Write a program that finds the summation of every number from 1 to num (both inclusive). The number will always be a positive integer greater than 0. Your function only needs to return the result, what is shown between parentheses in the example below is how you reach that result and it's not part of it, see the sample tests.

// For example (Input -> Output):

//go:build ignore

package main

import "fmt"

func Summation(num int) int {
	return num * (num + 1) / 2  // utiliza a fórmula da soma dos números naturais consecutivos
}

func main() {
	var num int
	fmt.Print("Digite um número: ")
	fmt.Scan(&num)

	fmt.Println(Summation(num))
}
