// Exercício 031: Leia o comprimento e o ângulo e calcule a componente seno/cosseno (math package). 

//go:build ignore

package main

import (
	"fmt"
	"math"
)

func main() {
	var hypotenuse, angle float64

	fmt.Print("Digite o comprimento da hipotenusa: ")
	fmt.Scan(&hypotenuse)
	fmt.Print("Digite o angle: ")
	fmt.Scan(&angle)

	radian := angle * (math.Pi / 180) // Transforma Graus em Radianos
	oppositeLeg := math.Sin(radian) * hypotenuse
	adjacentLeg := math.Cos(radian) * hypotenuse
	
	fmt.Printf("O cateto oposto é igual a %.2f\n", oppositeLeg)
	fmt.Printf("O cateto adjacente é igual a %.2f\n", adjacentLeg)
}