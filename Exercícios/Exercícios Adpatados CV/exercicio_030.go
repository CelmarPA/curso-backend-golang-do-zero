// Exercício 030: Leia um número e imprima os n primeiros termos da sequência de Fibonacci.

//go:build ignore

package main

import (
	"fmt"
)

func main() {
	var num int
	var sequenceFibonacci string

	fmt.Print("Digite um número para calcular sua sequencia de Fibonacci: ")
	fmt.Scan(&num)

	fa := 1
	fb := 1

	// Se o número for 1 ou 2, já tratamos direto
	if num ==0 {
		fmt.Println("Nenhum termo a exibir.")
		return
	}
	if num == 1 {
		fmt.Printf("A sequencia de Fibonacci para o número %d é: 1\n", num)
		return
	}
	if num == 2 {
		fmt.Printf("A sequencia de Fibonacci para o número %d é: 1, 1\n", num)
		return
	}

	sequenceFibonacci = "1, 1"

	for i := 3; i <= num; i++ {
		fn := fa + fb
		fb = fa
		fa = fn
		
		sequenceFibonacci += fmt.Sprintf(", %d", fn)
	}

	fmt.Printf("A sequencia de Fibonacci para o número %d é: %s\n", num, sequenceFibonacci)
}
