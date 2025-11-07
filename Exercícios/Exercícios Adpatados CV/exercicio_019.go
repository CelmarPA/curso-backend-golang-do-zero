// Exercício 019: Leia a idade de uma pessoa e classifique: criança, adolescente, adulto, idoso.

//go:build ignore

package main

import (
	"fmt"
)

func main() {
	var idade int
	
	fmt.Print("Digite sua idade: ")
	fmt.Scan(&idade)

	switch {
	case idade < 12:
		fmt.Println("Criança")
	case idade < 18:
		fmt.Println("Adolescente")
	case idade < 60:
		fmt.Println("Adulto")
	default:
		fmt.Println("Idoso")
	}
}
