// Exercício 004: Faça um programa que leia uma entrada do usuário e identifique se ela representa um número inteiro, número real ou texto, exibindo o tipo identificado.

//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main () {
	fmt.Print("Digite algo para verifica o que ela representa: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Verifica se é inteiro
	if _, err := strconv.Atoi(input); err == nil {
		fmt.Println("Tipo: número inteiro!")
		return
	}

	// Verifica se é float
	if _, err := strconv.ParseFloat(input, 64); err == nil {
		fmt.Println("Tipo: número real (float)")
		return
	}

	// Se não é inteiro nem float é texto
	fmt.Println("Tipo: texto")
}
