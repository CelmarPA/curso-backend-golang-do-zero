// Given an array of integers.

// Return an array, where the first element is the count of positives numbers and the second element is sum of negative numbers. 0 is neither positive nor negative.

// If the input is an empty array or is null, return an empty array.

// Example
// For input [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15], you should return [10, -65].

//go:build ignore

package main

import "fmt"

func CountPositivesSumNegatives(numbers []int) []int {
	var res []int
	var count, sumNegative int
	if len(numbers) == 0 {
		return res
	}
	for _, number := range(numbers) {	
		if number < 0 {
			sumNegative += number
		} else if number > 0 {
			count++
		}
	}

	res = append(res, count, sumNegative)
	return res
}

func main() {
	fmt.Println(CountPositivesSumNegatives([]int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}))
}