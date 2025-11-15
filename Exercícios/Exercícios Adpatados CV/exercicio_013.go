// Exercício 013: Leia o salário de um funcionário e mostre seu novo salário com 15% de aumento.

//go:build ignore

package main

import (
	"fmt"
)

func main() {
	var wage float64
	fmt.Print("Digite o salário do funcionário R$")
	fmt.Scan(&wage)

	raise := wage * 0.15
	newWage := wage + raise

	fmt.Printf("O valor do novo salário é de R$%.2f!\n", newWage)
}
