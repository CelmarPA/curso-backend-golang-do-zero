// Exercício 037: Escreva um programa que leia um número inteiro qualquer e peça para o usuário escolher qual será a base de conversão:
// 1 para binário, 2 para octal, 3 para hexadecimal

//go:build ignore

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num, opc int

	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&num)
	fmt.Print("Qual para qual base seja converter [1 - Binário | 2 - Octal | 3 - Hexadecimal]: ")
	fmt.Scan(&opc)

	switch 	opc {
	case 1:
		binaryStr := strconv.FormatInt(int64(num), 2)

		fmt.Printf("O número %d em binário é: %s\n", num, binaryStr)

	case 2:
		octalStr := strconv.FormatInt(int64(num), 8)

		fmt.Printf("O número %d em octal é: %s\n", num, octalStr)

	case 3:
		hexadecimalStr := strconv.FormatInt(int64(num), 12)

		fmt.Printf("O número %d em hexadecimal é: %s\n", num, hexadecimalStr)

	default:
		fmt.Println("Opção inválida tente novamente!")
	}
}
