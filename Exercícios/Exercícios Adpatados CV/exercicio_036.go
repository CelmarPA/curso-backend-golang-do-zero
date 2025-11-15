// Exercício 036: Escreva um programa para aprovar o empréstimo bancário para a compra de uma casa. O programa vai perguntar o valor da casa. 
// O salário do comprador e em quantos anos ele vaio pagar. Calculo o valor da prestação mensal, sabendo que ela não pode exceder 30% do 
// salário ou então o empréstimo será negado.

//go:build ignore

package main

import (
	"fmt"
)

func main() {
	var houseCost, wage float64
	var years int

	fmt.Print("Informe o valor do imóvel: R$")
	fmt.Scan(&houseCost)
	fmt.Print("Digite o seu salário atual: R$")
	fmt.Scan(&wage)
	fmt.Print("Digite em quantos anos deseja pagar: ")
	fmt.Scan(&years)

	if installments := houseCost / float64(years * 12); installments <= wage * 0.3 {
		fmt.Println("Empréstimo concedido!")
		return
	}

	fmt.Println("Empréstimo negado valor das prestações ultrapassa 30%% do salário!")
}
