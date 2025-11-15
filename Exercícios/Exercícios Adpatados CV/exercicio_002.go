// Exercício 002 - Faça um programa que leia o nome de uma pessoa e mostre uma mensagem de boas-vindas personalizada.

//go:build ignore

package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Digite seu nome: ")
	nome, _ := reader.ReadString('\n')
	nome = strings.TrimSpace(nome)

	fmt.Printf("Olá %s!\n", nome)
}
