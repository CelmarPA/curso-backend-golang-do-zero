// Task
// You get an array of numbers, return the sum of all of the positives ones.

// Example
// [1, -4, 7, 12] => 1 + 7 + 12 = 20

// Note
// If there is nothing to sum, the sum is default to 0.

//go:build ignore

package main

import "fmt"

func PositiveSum(numbers []int) (sumPositive int) {

	for _, number := range(numbers) {
		if number > 0 {
			sumPositive += number
		}
	}
	
	return sumPositive
}

func main() {
	fmt.Println(PositiveSum([]int{1, 2, 3, 4, 5}))
	fmt.Println(PositiveSum([]int{-1, -2, -3, -4, -5}))
}
