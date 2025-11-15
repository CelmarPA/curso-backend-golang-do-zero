// Exercício 030: Leia um número inteiro e diga se ele é par ou ímpar.

//go:build ignore

package main

import (
	"fmt"
)

func main() {
	var num int
	
	fmt.Print("Digite um número inteiro para saber se é par ou ímpar: ")
	fmt.Scan(&num)

	if num % 2 == 0 {
		fmt.Printf("O número %d é PAR!\n", num)
	} else {
		fmt.Printf("O número %d é ÍMPAR!\n", num)
	}
}
