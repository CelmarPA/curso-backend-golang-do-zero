// Exercício 027: Leia um número e verifique se é um número primo.

//go:build ignore

package main

import (
	"fmt"
)

func main() {
	var count, num int

	fmt.Print("Digite um número para verifica se é primo: ")
	fmt.Scan(&num)

	if num < 2 {
        fmt.Println(num, "não é primo")
        return
    }

	if num >= 2 {
		for i := 1; i <= num / 2; i++ {
			if num % i == 0 {
				count++
			}
		}
	}
	
	count++

	if count == 2 {
		fmt.Println(num, "é primo")
	} else {
		fmt.Println(num, "não é primo")
	}
}
