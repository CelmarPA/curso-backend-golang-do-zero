// Exercício 033: Leia três números e mostre qual é o maior e o menor.

//go:build ignore

package main

import (
	"fmt"
	"slices"
)

func main() {
	var num1, num2, num3, maior, menor int
	numeros := []int{}

	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&num1)
	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&num2)
	fmt.Print("Digite o terceiro número: ")
	fmt.Scan(&num3)

	numeros = append(numeros, num1, num2, num3)

	menor = slices.Min(numeros)
	maior = slices.Max(numeros)

	fmt.Printf("O maior número é %d e o menor número é %d.\n", maior, menor)
}
