// Exercício 023: Leia um número de 0 a 9999 e mostre separadamente cada um dos dígitos (unidade, dezena, centena e milhar). 

//go:build ignore

package main

import (
	"fmt"
)

func main() {
	var num int

	fmt.Print("Digite um número de 0 a 9999: ")
	fmt.Scan(&num)

	unit := num / 1 % 10
	tens := num / 10 % 10
	hundreds := num / 100 % 10
	thousands := num / 1000 % 10

	fmt.Printf("O número %d possui:\n", num)
	fmt.Printf("Unidade: %d\n",unit)
	fmt.Printf("Dezena: %d\n",tens)
	fmt.Printf("Centena: %d\n",hundreds)
	fmt.Printf("Milhar: %d\n",thousands)
}
