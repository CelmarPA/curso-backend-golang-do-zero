// Exercício 008: Converta metros para centímetros; leia um valor em metros (float).

//go:build ignore

package main

import (
	"fmt"
)

func main() {
	var meter float64
	fmt.Print("Digite o valor em metros para converter para centímetros:")
	fmt.Scan(&meter)

	centimeter := meter * 100

	fmt.Println("A medida informada foi", meter, "metros, que convertida em centímetros é igual a", centimeter)
}
