// Exercício 011: Leia as dimensões de um terreno (largura e comprimento) e calcule a área.

//go:build ignore

package main

import (
	"fmt"
)

func main() {
	var width, length float64
	fmt.Print("Digite a largura do terreno (m): ")
	fmt.Scan(&width)
	fmt.Print("Digite o comprimento do terreno (m): ")
	fmt.Scan(&length)

	area := width * length
	fmt.Printf("A área do terreno é de %.2f m².\n", area)
}
