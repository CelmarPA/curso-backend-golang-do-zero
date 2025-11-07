// Exercício 014: Converta metros para milhas (use constante de conversão).

//go:build ignore

package main

import (
	"fmt"
)

func main() {
	const fator = 1609

	var distMetros float64

	fmt.Print("Digite a distância em metros para converter em milhas: ")
	fmt.Scan(&distMetros)

	distMilhas := distMetros / fator

	fmt.Printf("A distância de %.2f metros convertida é igual a %.4f milhas.\n", distMetros, distMilhas)
}
