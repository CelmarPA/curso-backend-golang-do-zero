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
	for i := range(10) {
		fmt.Println(num, "X", i + 1, "=", num * (i + 1))
	}
}
