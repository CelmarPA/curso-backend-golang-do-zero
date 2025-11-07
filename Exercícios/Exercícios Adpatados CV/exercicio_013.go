// Exercício 013: Calcule a média ponderada de três valores com pesos 2,3,5.

//go:build ignore

package main

import (
	"fmt"
)

func main() {
	var v1, v2, v3 float64
	fmt.Print("Digite o primeiro valor: ")
	fmt.Scan(&v1)
	fmt.Print("Digite o segundo valor: ")
	fmt.Scan(&v2)
	fmt.Print("Digite o terceiro valor: ")
	fmt.Scan(&v3)


	weightedAverage := ((v1 * 2) + (v2 * 3) + (v3 * 5)) / (2 + 3 + 5)

	fmt.Printf("A média ponderada dos valores: %.2f, %.2f, %.2f é igual a %.2f.\n", v1, v2, v3, weightedAverage)
}
