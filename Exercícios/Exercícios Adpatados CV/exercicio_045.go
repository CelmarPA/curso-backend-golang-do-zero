// Exercício 045: Crie um programa que faça o computador jogar Jokenpô (pedra, papel e tesoura) com você.

//go:build ignore

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	var option int

	// Cria um gerador de números aleatórios independente, com seu próprio estado interno.
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	fmt.Println("#=#=# Bem vindo ao Jokenpô!!! #=#=#")
	fmt.Println(strings.Repeat("=", 35))
	fmt.Println("  Escolha uma da seguintes opções: ")
	fmt.Println("  1 -> PEDRA ")
	fmt.Println("  2 -> PAPEL ")
	fmt.Println("  3 -> TESOURA ")
	fmt.Println(strings.Repeat("=", 35))
	fmt.Print("Sua jogada: ")
	fmt.Scan(&option)

	if option > 3 {
		fmt.Println("Opção inválida, comece novamente!")
		return
	}

	moves := []string {"PEDRA", "PAPEL", "TESOURA"}

	playerMove := moves[option - 1]
	
	// Faz o computador escolher um jogada aleatória
	min := 0
	max := 2
	cpuMove := moves[r.Intn(max - min + 1) + min]

	fmt.Printf("Você jogou: %s\n", playerMove)
	fmt.Printf("CPU jogou: %s\n", cpuMove)

	if cpuMove == "PEDRA" && playerMove == "TESOURA" {
		fmt.Printf("Você perdeu: %s ganha de %s\n", cpuMove, playerMove)
	} else if cpuMove == "PAPEL" && playerMove == "PEDRA" {
		fmt.Printf("Você perdeu: %s ganha de %s\n", cpuMove, playerMove)
	} else if cpuMove == "TESOURA" && playerMove == "PAPEL" {
		fmt.Printf("Você perdeu: %s ganha de %s\n", cpuMove, playerMove)
	} else if cpuMove == playerMove {
		fmt.Printf("Empate: %s igual %s\n", cpuMove, playerMove)
	} else {
		fmt.Printf("Você ganhou: %s ganha de %s\n", playerMove, cpuMove)
	}
}
