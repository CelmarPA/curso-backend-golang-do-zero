// Exercício 028: Jogo: o computador 'escolhe' um número de 0 a 5 e o usuário tenta adivinhar.

//go:build ignore

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var cpuGuess, userGuess, chances int
	var winner bool

	fmt.Println("Tente adivinha o número que eu pensei entre 1 a 5 você tem 2 chances!!")

	min := 1
	max := 5
	cpuGuess = rand.Intn(max - min + 1) + min // Número aleatório em  um intervalo personalizado
	chances = 2

	for {		
		if chances > 0 {
			fmt.Printf("Você possui %d chances!\n", chances)
			fmt.Print("Qual seu palpite? ")
			fmt.Scan(&userGuess)
			
			if userGuess == cpuGuess {
				winner = true
				break
			} else {
				chances--
			}
		} else {
			break
		}
	}

	if winner {
		fmt.Printf("Você ganhou eu pensei no número %d!!!\n", cpuGuess)
	} else {
		fmt.Printf("Você perdeu eu pensei no número %d!!!\n", cpuGuess)
	}
}
