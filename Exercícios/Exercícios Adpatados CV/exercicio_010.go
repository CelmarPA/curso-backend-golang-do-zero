// Exercício 010: Leia quanto dinheiro a pessoa tem na carteira e mostre quantos dólares ela pode comprar (considere uma taxa fixa, ou permita que o usuário informe a taxa)

//go:build ignore

package main

import (
	"fmt"
)

func main() {
	var dollarCost, real float64
	fmt.Print("Digite a cotação atual do dólar: R$")
	fmt.Scan(&dollarCost)
	fmt.Print("Digite o valor em reais para calcular a quantidade de dólares que podem ser adquiridos: R$")
	fmt.Scan(&real)

	dollars := real / dollarCost

	fmt.Printf("Com R$%.2f você pode comprar US$%.2f.\n", real, dollars)
}
