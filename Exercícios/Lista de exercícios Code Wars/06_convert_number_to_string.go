// We need a function that can transform a number (integer) into a string.

// What ways of achieving this do you know?

// Examples (input --> output):
// 123  --> "123"
// 999  --> "999"
// -100 --> "-100"

//go:build ignore

package main

import (
	"fmt"
	"strconv"
)

func NumberToString(num int) string {
	return strconv.Itoa(num)
}

func main() {
	fmt.Println(NumberToString(1234))
}
