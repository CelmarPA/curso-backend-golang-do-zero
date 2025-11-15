// Exercício 012: Leia o preço de um produto e mostre seu novo preço com 5% de desconto.

//go:build ignore

package main

import (
	"fmt"
)

func main() {
	var cost float64
	fmt.Print("Digite o valor do produto R$")
	fmt.Scan(&cost)

	desc := cost * 0.05
	newCost := cost - desc

	fmt.Printf("O valor do produto com desconto é de R$%.2f volte sempre.\n", newCost)
}
