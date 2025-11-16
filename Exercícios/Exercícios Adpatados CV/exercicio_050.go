// Exercício 050: Desenvolva um programa que leia seis números inteiros e mostre a soma apenas daqueles que forem pares. Se o valor digitado for ímpar, desconsidere-o.

//go:build ignore

package main

import (
	"fmt"
)

func main() {
	var num int
	sum := 0

	for i := 1; i <= 6; i++ {
		fmt.Printf("Digite o %dº número: ", i)
		fmt.Scan(&num)

		if num % 2 == 0 {
			sum += num
		}
	}

	fmt.Printf("A soma dos número pares é igual a %d.\n", sum)
}
