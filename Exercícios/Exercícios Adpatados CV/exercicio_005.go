// Exercício 005: Faça um programa que leia um número inteiro e mostre o seu antecessor e sucessor.

//go:build ignore

package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Print("Digite um número para saber seu antecessor e sucessor: ")
	fmt.Scan(&num)
	fmt.Println("O número digitado foi", num, "seu antecessor é", num - 1, "e seu sucessor é", num + 1)
}
