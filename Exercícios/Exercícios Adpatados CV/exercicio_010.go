// Exercício 010: Leia a temperatura em Celsius e converta para Fahrenheit.

//go:build ignore

package main

import (
	"fmt"
)

func main() {
	var celsius float64
	fmt.Print("Digita a temperatura em Celsius para converter em Fahrenheit: ")
	fmt.Scan(&celsius)
	fahreneit := (celsius * 9 / 5) + 32
	
	fmt.Println("Convertendo em fahreneit temos:", fahreneit, "F.")
}
