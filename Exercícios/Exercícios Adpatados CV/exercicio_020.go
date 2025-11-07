// 020. Exercício 020: Leia números até o usuário digitar 0; ao final mostre soma, quantidade e média.

//go:build ignore

package main

import (
	"fmt"
)

func main() {
	var num, sum, count int 
	
	for {
		fmt.Print("Digite um número [0 para finalizar]: ")
		fmt.Scan(&num)

		if num == 0 {
			break
		} else {
			count++
			sum += num
		}
	}
	
	mean := sum / count

	fmt.Printf("A soma dos valores digitados é: %d\n", sum)
	fmt.Printf("Você digitou %d números.\n", count)
	fmt.Printf("A média dos valores digitados é: %.2f\n", float64(mean))
}
