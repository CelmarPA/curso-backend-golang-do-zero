// Exercício 008: Leia um valor em metros e o converta para centímetros e milímetros.

//go:build ignore

package main

import (
	"fmt"
)

func main() {
	var meter float64
	fmt.Print("Digite o valor em metros para converter para centímetros: ")
	fmt.Scan(&meter)

	centimeter := meter * 100
	millimeter :=  meter * 1000

	fmt.Println("A medida informada foi", meter, "metros, que convertida em centímetros é igual a", centimeter, "e em milímetros igual a", millimeter)
}
