// Exercício 007: Leia as duas notas de um aluno, calcule e mostre a sua média. 

//go:build ignore

package main

import (
	"fmt"
)

func main() {
	var grade1, grade2 float64
	fmt.Print("Digite a primeira nota: ")
	fmt.Scan(&grade1)
	fmt.Print("Digite a segunda nota: ")
	fmt.Scan(&grade2)
	mean := (grade1 + grade2) / 2

	fmt.Println("A média do aluno é", mean)
}
