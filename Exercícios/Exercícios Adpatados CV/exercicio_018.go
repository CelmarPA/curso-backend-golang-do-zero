// Exercício 018: Leia um ângulo em graus e mostre o seno, cosseno e tangente.

//go:build ignore

package main

import (
	"fmt"
	"math"
)

func main() {
	var angle float64

	fmt.Print("Digite o ângulo para obter seu seno, cosseno e tangente: ")
	fmt.Scan(&angle)

	radians := angle * (math.Pi / 180)

	sine := math.Sin(radians)
	cosine := math.Cos(radians)
	tangent := math.Tan(radians)

	fmt.Printf("O ângulo de %.2f graus tem seno de %.3f, cosseno de %.3f e tangente de %.3f!\n", angle, sine, cosine, tangent)
}