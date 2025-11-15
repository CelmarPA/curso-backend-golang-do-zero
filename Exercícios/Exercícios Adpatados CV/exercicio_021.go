// Exercício 021: Faça um programa que leia o preço de um produto e mostre se ele está com promoção (aplique condições). 

//go:build ignore

package main

import (
	"fmt"
)

func main() {
	var value float64

	fmt.Print("Digite o valor do produto: R$")
	fmt.Scan(&value)

	if value <= 5 {
		fmt.Println("O produto não está em promoção!")
		return
	} else if value <= 10 {
		fmt.Println("O produto está em promoção com 5% desconto!")
		return
	} else if value <= 50 {
		fmt.Println("O produto está em promoção com 10% desconto!")
		return
	}

	fmt.Println("O produto está em promoção com 15% desconto!")
}
