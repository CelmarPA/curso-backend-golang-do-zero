// Exercício 017: Leia os comprimentos dos catetos de um triângulo retângulo e calcule a hipotenusa.

//go:build ignore

package main

import (
	"fmt"
	"math"
)

func main() {
	var adjacentLeg, oppositeLeg float64

	fmt.Print("Digite o valor do cateto adjacente: ")
	fmt.Scan(&adjacentLeg)
	fmt.Print("Digite o valor do cateto oposto: ")
	fmt.Scan(&oppositeLeg)

	hypotenuse := math.Sqrt(math.Pow(adjacentLeg, 2) + math.Pow(oppositeLeg, 2))

	fmt.Printf("O valor da hipotenusa calculado é de %.2f!\n", hypotenuse)
}
