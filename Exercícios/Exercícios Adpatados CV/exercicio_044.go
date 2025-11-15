// Exercício 044: Elabore um programa que calcule o valor a ser pago por um produto, considerando o seu preço normal e condição de pagamento: 
// À vista dinheiro/ pix /cheque: 10% de desconto; À vista no cartão: 5% de desconto; Em até 2x no cartão: preço normal; 3x ou mais no cartão: 20% de juros.

//go:build ignore

package main

import (
	"fmt"
)

func main() {
	var cost, discount, interest, finalCost float64
	var option int

	fmt.Print("Informe o valor do produto: R$")
	fmt.Scan(&cost)

	fmt.Println("======== Opções de Pagamento ========")
	fmt.Println(" 1 -> À vista dinheiro | Pix | Cheque")
	fmt.Println(" 2 -> À vista cartão")
	fmt.Println(" 3 -> 2 x no cartão")
	fmt.Println(" 4 -> 3x ou mais no cartão")
	fmt.Println("=====================================")

	fmt.Print("Informe a opção de pagamento: ")
	fmt.Scan(&option)

	switch option {
	case 1:
		discount = 0.10 * cost
		finalCost = cost - discount
		fmt.Println("10% de desconto aplicado!")

	case 2:
		discount = 0.05 * cost
		finalCost = cost - discount
		fmt.Println("5% de desconto aplicado!")

	case 3:
		finalCost = cost

	case 4:
		interest = 0.20 * cost
		finalCost = cost + interest
		fmt.Println("20% de juros aplicados!")
		
	default:
		fmt.Println("Opção inválida!")
		return
	}

	fmt.Printf("O valor final da sua compra será de R$%.2f volte sempre!\n", finalCost)
}
