// 016. Exercício 016: Leia três números e mostre o maior e o menor.

//go:build ignore

package main

import (
	"fmt"
	"slices"
)

// Solução utilizando if / else if / else
// func main() {
// 	var num1, num2, num3, maior, menor float64

// 	fmt.Print("Digite o primeiro número: ")
// 	fmt.Scan(&num1)
// 	fmt.Print("Digite o segundo número: ")
// 	fmt.Scan(&num2)
// 	fmt.Print("Digite o terceiro número: ")
// 	fmt.Scan(&num3)

// 	if num1 > num2 && num1 > num3 {
// 		maior = num1
// 	} else if num2 > num1 && num2 > num3 {
// 		maior = num2
// 	} else {
// 		maior = num3
// 	}

// 	if num1 < num2 && num1 < num3 {
// 		menor = num1
// 	} else if num2 < num1 && num2 < num3 {
// 		menor = num2
// 	} else {
// 		menor = num3
// 	}

// 	fmt.Printf("O maior número é %.2f e o menor número é %.2f.\n", maior, menor)
// }

// Solução utilizando listas
func main() {
	var num1, num2, num3, maior, menor float64
	numeros := []float64{}

	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&num1)
	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&num2)
	fmt.Print("Digite o terceiro número: ")
	fmt.Scan(&num3)

	numeros = append(numeros, num1, num2, num3)

	maior = slices.Min(numeros)
	menor = slices.Max(numeros)

	fmt.Printf("O maior número é %.2f e o menor número é %.2f.\n", maior, menor)
}
