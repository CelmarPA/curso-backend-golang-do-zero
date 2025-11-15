// Exercício 043: Desenvolva uma lógica que leia o peso e a altura de uma pessoa, calcule seu IMC e mostre seu status, de acordo com a tabela abaixo: 
// Abaixo de 18.5: Abaixo do peso; Entre 18.5 e 25: Peso ideal; 25 até 30: Sobrepeso; 30 até 40: Obesidade Acima de 40: Obesidade mórbida.

//go:build ignore

package main

import (
	"fmt"
)

func main() {
	var weight, height float64
	var status string
		
	fmt.Print("Digite seu peso em Kg: ")
	fmt.Scan(&weight)
	fmt.Print("Digite sua altura em metros: ")
	fmt.Scan(&height)

	if height < 0 {
		fmt.Println("Altura inválida!")
		return
	}

	imc := weight / (height * height)
	fmt.Println(imc)

	switch {
	case imc < 18.5:
		status = "Abaixo do peso"
	
	case imc < 25:
		status = "Peso ideal"
	
	case imc < 30:
		status = "Sobrepeso"
		
	case imc < 40:
		status = "Obesidade"	
		
	default:
		status = "Obesidade mórbida"
	}
	
	fmt.Printf("Seu IMC atual é de %.2f e seus status: %s\n", imc, status)
}
