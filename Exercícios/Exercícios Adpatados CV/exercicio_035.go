// Exercício 035: Desenvolva um programa que leia o comprimento de três retas e diga ao usuário se elas podem ou não formar um triângulo.

//go:build ignore

package main

import (
	"fmt"
)

func main() {
	var a, b, c float64
	
	fmt.Print("Digite o primeiro lado: ")
	fmt.Scan(&a)
	fmt.Print("Digite o segundo lado: ")
	fmt.Scan(&b)
	fmt.Print("Digite o terceiro lado: ")
	fmt.Scan(&c)

	// Para um triângulo existir, a soma de quaisquer dois de seus lados deve ser sempre maior que o terceiro lado.

	if a + b > c && a + c  > b && b + c > a {
		fmt.Println("Os lados podem formar um triângulo.")
	} else {
		fmt.Println("Os lados não pode formar um triângulo")
	}
}
