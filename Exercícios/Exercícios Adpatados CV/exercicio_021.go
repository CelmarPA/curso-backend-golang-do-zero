// Exercício 021: Faça um programa que jogue par ou ímpar com o usuário (use rand).

//go:build ignore

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var player int

	fmt.Print("Digite um número entre 0 e 10 para jogar par ou ímpar: ")
	fmt.Scan(&player)

	cpu := rand.Intn(11)
	result := player + cpu

	if result % 2 == 0 {
		fmt.Printf("Resultado %d PAR!\n", result)
	} else {
		fmt.Printf("Resultado %d ÍMPAR!\n", result)
	}
}
