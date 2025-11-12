// Exercício 009: Leia um número e mostre sua tabuada (1 a 10).

//go:build ignore

package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Print("Digite um número para ver sua tabuada: ")
	fmt.Scan(&num)

	fmt.Println("======= TABUADA DO NÚMERO", num, "=======")
	fmt.Println(num, "X", 1, "=", num * 1)
	fmt.Println(num, "X", 2, "=", num * 2)
	fmt.Println(num, "X", 3, "=", num * 3)
	fmt.Println(num, "X", 4, "=", num * 4)
	fmt.Println(num, "X", 5, "=", num * 5)
	fmt.Println(num, "X", 6, "=", num * 6)
	fmt.Println(num, "X", 7, "=", num * 7)
	fmt.Println(num, "X", 8, "=", num * 8)
	fmt.Println(num, "X", 9, "=", num * 9)
	fmt.Println(num, "X", 10, "=", num * 10)
}
