// Given a set of numbers, return the additive inverse of each. Each positive becomes negatives, and the negatives become positives.

// [1, 2, 3, 4, 5] --> [-1, -2, -3, -4, -5]
// [1, -2, 3, -4, 5] --> [-1, 2, -3, 4, -5]
// [] --> []
// You can assume that all values are integers. Do not mutate the input array.

//go:build ignore

package main

import "fmt"

func Invert(arr []int) []int {
	var res [] int

	for _, number := range(arr) {
		res = append(res, -number)
	}

	return res
}


func main() {
	fmt.Println(Invert([]int {1, 2, 3, 4, 5}))
	fmt.Println(Invert([]int {1, -2, 3, -4, 5}))
	fmt.Println(Invert([]int {}))
	
}