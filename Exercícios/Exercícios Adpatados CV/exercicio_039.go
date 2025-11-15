// Exercício 039: Faça um programa que leia o ano de nascimento de um jovem e informe, de acordo com sua idade: 
// Se ele ainda vai se alistar ao serviço militar, se é a hora de se alistar, se já passou do tempo do alistamento. 
// Seu programa também deverá mostrar o tempo que falta ou que passou do prazo.

//go:build ignore

package main

import (
	"fmt"
	"time"
)

func main() {
	var birthYear int
	currentYear := time.Now().Year()

	fmt.Print("Digite o ano do seu nascimento: ")
	fmt.Scan(&birthYear)

	age := currentYear - birthYear

	if age < 18 {
		fmt.Printf("Você faltam %d ano(s) para se alistar.\n", 18 - age)
	} else if age > 18 {
		fmt.Printf("Seu alistamento deveria ter sido feito à %d ano(s).\n", age - 18)
	} else {
		fmt.Println("Está na hora de se alistar procure um agência mais próxima de você!")
	}
}
