// Exercício 002: Leia uma string do usuário e responda com 'Olá, <nome>!'.

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
