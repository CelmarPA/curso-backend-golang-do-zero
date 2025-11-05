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
	var name string
	fmt.Print("Digite seu nome: ")
	fmt.Scan(&name)
	fmt.Println("Olá", name, "!")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Digite seu nome completo: ")
	nome, _ := reader.ReadString('\n')
	nome = strings.TrimSpace(nome)

	fmt.Println("Olá", nome, "!")
}
