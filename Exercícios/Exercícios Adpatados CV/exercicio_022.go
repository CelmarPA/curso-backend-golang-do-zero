// Exercício 022: Faça um programa que leia um número e gere sua tabuada usando loop 'for'.

//go:build ignore

package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Print("Digite um número para ver sua tabuada: ")
	fmt.Scan(&num)

	fmt.Println("======= TABUADA DO NÚMERO", num, "=======")
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d X %d = %d\n", num, i, num * i)
	} 
}
