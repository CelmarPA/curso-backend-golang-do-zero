// Exercício 031: Calcule o preço da passagem de acordo com a distância: até 200km R$0,50/km, senão R$0,45/km.

//go:build ignore

package main

import (
	"fmt"
)

func main() {
	var distance int 

	fmt.Print("Digite a distância em Km que deseja viajar para calcular o valor: ")
	fmt.Scan(&distance)

	if distance <= 200 {
		cost := float64(distance) * 0.50

		fmt.Printf("O valor da sua viagem de %d quilômetros será de R$%.2f!\n", distance, cost)
		return
	} 

	cost := float64(distance) * 0.45

	fmt.Printf("O valor da sua viagem de %d quilômetros será de R$%.2f!\n", distance, cost)
}