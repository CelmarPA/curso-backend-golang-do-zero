// Exercício 048: Faça um programa que calcule a soma entre todos os números que são múltiplos de três e que se encontram no intervalo de 1 até 500.

//go:build ignore

package main

import (
	"fmt"
)

func main() {
	sum := 0

	for i := 0; i <= 500; i ++ {
		if i % 3 == 0 {
			sum += i	
		}
	}

	fmt.Printf("A soma dos números ímpares entre 0 e 500 é igual %d\n", sum)
}
