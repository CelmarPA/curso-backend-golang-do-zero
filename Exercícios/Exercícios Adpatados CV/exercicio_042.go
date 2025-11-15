// Exercício 0042: Refaça o Exercício 35 dos triângulos, acrescentando o recurso de mostrar que tipo de triângulo será formado.

//go:build ignore

package main

import (
	"fmt"
)

func main() {
	var a, b, c float64
	var triangleType string

	fmt.Print("Digite o primeiro lado: ")
	fmt.Scan(&a)
	fmt.Print("Digite o segundo lado: ")
	fmt.Scan(&b)
	fmt.Print("Digite o terceiro lado: ")
	fmt.Scan(&c)

	// Para um triângulo existir, a soma de quaisquer dois de seus lados deve ser sempre maior que o terceiro lado.

	if a + b > c && a + c  > b && b + c > a {
		if a == b && b == c {
			triangleType = "Equilátero"
		} else if a == b || a == c || b == c {
			triangleType = "Isosceles"
		} else {
			triangleType = "Escaleno"
		}
		fmt.Printf("Os lados podem formar um triângulo %s.\n", triangleType)
	} else {
		fmt.Println("Os lados não pode formar um triângulo")
	}
}
