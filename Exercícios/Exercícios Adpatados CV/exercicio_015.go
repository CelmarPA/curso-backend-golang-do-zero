// Exercício 015: Calcule o preço a pagar por um aluguel de carro: R$60 por dia + R$0,15 por km rodado.

//go:build ignore

package main

import (
	"fmt"
)

func main() {
	var day int
	var distance float64

	fmt.Print("Digite quantos dias de aluguel: ")
	fmt.Scan(&day)
	fmt.Print("Digite quantos quilômetros rodados: ")
	fmt.Scan(&distance)

	cost := float64(60 * day) + (distance * 0.15)

	fmt.Printf("O valor a pagar por %d dias e %.2f quilômetros rodados é de R$%.2f!\n", day, distance, cost)
}
