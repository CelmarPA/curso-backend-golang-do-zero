// Exercício 040: Crie um programa que leia duas notas de um aluno e calcule a sua média, mostrando uma mensagem no final, de acordo com a média atingida: 
// Média abaixo de 5.0: REPROVADO; Média entre 5.0 e 6.9: RECUPERAÇÃO; Média 7.0 ou superior: APROVADO.

//go:build ignore

package main

import (
	"fmt"
)

func main() {
	var grade1, grade2 float64
	var status string

	fmt.Print("Digite a primeira nota: ")
	fmt.Scan(&grade1)
	fmt.Print("Digite a segunda nota: ")
	fmt.Scan(&grade2)

	mean := (grade1 + grade2) / 2

	if mean < 5 {
		status = "REPROVADO!"
	} else if mean <= 6.9 {
		status = "RECUPERAÇÃO!"
	} else {
		status = "APROVADO!"
	}

	fmt.Printf("A média do aluno foi de %.2f pontos: %s\n", mean, status)
}
