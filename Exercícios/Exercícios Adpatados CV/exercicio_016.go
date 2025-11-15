//  Exercício 016: Leia um número real e mostre sua parte inteira, descartando a parte decimal.

//go:build ignore

package main

import (
	"fmt"
	"math"
)

func main() {
	var real float64
	
	fmt.Print("Digite um número real para obter sua parte inteira: ")
	fmt.Scan(&real)

	integer := math.Trunc(real)
	
	fmt.Printf("O parte inteira do número real %f é: %.0f\n", real, integer)
}
