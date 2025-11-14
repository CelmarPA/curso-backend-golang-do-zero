// Exercício 032: Leia o ano de nascimento e diga se a pessoa já pode se alistar (considere ano atual).

//go:build ignore

package main

import (
	"fmt"
	"time"
)

func main() {
	var yearOfBirth int
	year := time.Now().Year()

	fmt.Print("Digite o ano do seu nascimento: ")
	fmt.Scan(&yearOfBirth)

	if year - yearOfBirth >= 18 {
		fmt.Println("Você já pode se alistar esse ano!")
	} else {
		fmt.Println("Você ainda não pode alistar!")
	}
}
