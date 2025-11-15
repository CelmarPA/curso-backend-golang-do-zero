// Exercício 011: Leia a largura e altura de uma parede, calcule a área e a quantidade de tinta necessária (1L = 2m²).

//go:build ignore

package main

import (
	"fmt"
)

func main() {
	var width, length float64
	fmt.Print("Digite a largura da parede (m): ")
	fmt.Scan(&width)
	fmt.Print("Digite a altura da parede (m): ")
	fmt.Scan(&length)

	area := width * length
	litre := area / 2 

	fmt.Printf("A área da parede é de %.2f m² e são necessários %.2f litros de tinta para pinta-la.\n", area, litre)
}
