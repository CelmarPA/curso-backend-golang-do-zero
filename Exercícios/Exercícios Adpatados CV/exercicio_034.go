// Exercício 034: Escreva um programa que pergunte o salário de um funcionário e calcule o valor do seu aumento. 
// Para salários superiores a R$1200,00, calcule um aumento de 10%. Para salários inferiores ou iguais a R$1200,00, o aumento é de 15%.

//go:build ignore

package main

import (
	"fmt"
)

func main() {
	var wage float64

	fmt.Print("Digite o salário do funcionário: R$ ")
	fmt.Scan(&wage)

	if wage <= 1200 {
		increase := wage * 0.15

		fmt.Printf("O novo salário do funcionário com aumento de 15%% será de R$%.2f!\n", wage + increase)
	} else {
		increase := wage * 0.10

		fmt.Printf("O novo salário do funcionário com aumento de 10%% será de R$%.2f!\n", wage + increase)
	}
}
