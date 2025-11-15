// Exercício 038: Leia dois números inteiros e compare-os dizendo qual é maior ou se são iguais.

//go:build ignore

package main

import (
	"fmt"
)

func main() {
	var num1, num2 int

	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&num1)
	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&num2)

	if num1 > num2 {
		println("O primero número é maior!")
	} else if num2 > num1 {
		println("O segundo número é maior!")
	} else {
		println("Os números são iguais!")
	}
}
