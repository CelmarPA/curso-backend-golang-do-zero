// Exercício 006: Leia um número e mostre o seu dobro, triplo e raiz quadrada.

//go:build ignore

package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64
	fmt.Print("Digite um número para calcular seu dobro, triplo e raiz quadrada: ")
	fmt.Scan(&num)

	double := num * 2
	triple := num * 3
	sqrt := math.Sqrt(num)

	fmt.Println("Você digitou o número,", num, "o dobro é", double, ", o triple é", triple, "e a raiz quadrada é", sqrt)
}
