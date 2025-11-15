// Exercício 003: Leia dois números e mostre a soma entre eles.

//go:build ignore

package main

import "fmt"

func main() {
	var x, y int

	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&x)

	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&y)

	fmt.Println("A soma é:", x + y)
}
