// Write a function that accepts a non-negative integer n and a string s as parameters, and returns a string of s repeated exactly n times.

// Examples (input -> output)
// 6, "I"     -> "IIIIII"
// 5, "Hello" -> "HelloHelloHelloHelloHello"

//go:build ignore

package main

import (
	"fmt"
	"strings"
)

func StringRepeat(n int, s string) string {
	str := ""
	for i := 0; i < n; i++ {
		str += s
	}
	return str
}

func RepeatStr(repetitions int, value string) string {
	return strings.Repeat(value, repetitions)
}

func main() {
	fmt.Println(StringRepeat(5, "Hello"))
	fmt.Println(RepeatStr(5, "Hello"))
}