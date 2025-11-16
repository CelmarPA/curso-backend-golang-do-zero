// Exercício 049: Refaça o exercício 9, mostrando a tabuada de um número que o usuário escolher, só que agora utilizando um laço for.

//go:build ignore

package main

import (
	"fmt"
)

func main() {
	var num int

	fmt.Print("Digite um número para ver sua tabuada: ")
	fmt.Scan(&num)
	fmt.Printf("====== Tabuada do número %d ======\n", num)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d X %d = %d\n", i, num, num * i)	
	}
}
