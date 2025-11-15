// Exercício 041: A Confederação  Nacional de Natação precisa de um programa que leia o ano de nascimento de um atleta e mostre sua categoria, de acordo com a idade:
// Até 9 anos: MIRIM; Até 14 anos:  INFANTIL; Até 19 anos: JUNIOR; Até 25 anos: SÊNIOR e Acima : MASTER.

//go:build ignore

package main

import (
	"fmt"
	"time"
)

func main() {
	var birthYear int
	var category string
	currentYear := time.Now().Year()

	fmt.Print("Digite o ano de nascimento do atleta: ")
	fmt.Scan(&birthYear)

	age := currentYear - birthYear
	
	switch {
	case age <= 9:
		category = "MIRIM"

	case age <= 14:
		category = "INFANTIL"
	
	case age <= 19:
		category = "JUNIOR"

	case age <= 25:
		category = "SÊNIOR"
	
	default:
		category = "MASTER"
	}

	fmt.Printf("O atleta tem %d anos de idade e se enquadra na categoria %s!\n", age, category)
}
