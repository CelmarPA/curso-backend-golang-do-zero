// Exercício 032: Faça um programa que leia um ano qualquer e mostre se ele é BISSEXTO.

//go:build ignore

package main

import (
	"fmt"
	"time"
)

func main() {
	var year int

	fmt.Print("Digite o ano para saber se é bissexto: ")
	fmt.Scan(&year)

	if year == 0 {
		year = time.Now().Year()
	}

	if (year % 4 == 0 && year % 100 != 0) || year % 400 == 0 {
		fmt.Printf("O ano de %d é BISSEXTO!\n", year)
	} else {
		fmt.Printf("O ano de %d NÃO é BISSEXTO!\n", year)
	}
}
