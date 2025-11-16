// Exercício 047: Crie um programa que mostre na tela todos os números pares que estão no intervalo entre 1 e 50.

//go:build ignore

package main

import (
	"fmt"
	"strconv"
)

func main() {
	evenNumbers := ""

	fmt.Println("Os número pares de 0 a 50 são: ")
	
	for i := 0; i <= 50; i++ {
		if i % 2 == 0 {
			if i == 50 {
				evenNumbers += fmt.Sprintf("%s", strconv.Itoa(i))
				break
			} 
			evenNumbers += fmt.Sprintf("%s, ", strconv.Itoa(i))
		}
	}
	
	fmt.Println(evenNumbers)
}