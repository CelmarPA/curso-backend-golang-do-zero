// Exercício 028: Leia várias idades e calcule quantas pessoas são maiores de idade.

//go:build ignore

package main

import (
	"fmt"
)

func main() {
	var count, major, idade int

	fmt.Println("Digite as idades desejadas [0 para parar]!")
	
	for {
		fmt.Print("Digite uma idade: ")
		fmt.Scan(&idade)
		if idade == 0 {
			break
		}

		count++
		if idade >= 18 {
			major++
		}
	}

	if count == 0 {
		return
	}

	fmt.Printf("Das %d idades que você digitou %d pessoas são maiores de idade.\n", count, major)
}
