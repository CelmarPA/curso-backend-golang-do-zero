// Exercício 029: Jogo da adivinhação: o computador pensa em um número e o usuário tenta acertar.

//go:build ignore

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var cpuGuess, userGuess, chances int
	var winner bool

	fmt.Println("Tente adivinha o número que eu pensei entre 1 a 10 você tem 3 chances!!")

	min := 1
	max := 10
	cpuGuess = rand.Intn(max - min + 1) + min // Número aleatório em  um intervalo personalizado
	chances = 3
	fmt.Println(cpuGuess)
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
