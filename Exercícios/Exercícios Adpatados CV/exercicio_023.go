// Exercício 023: Leia um número e mostre seus divisores.

//go:build ignore

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int
	var divisors string
	
	fmt.Print("Digite um número inteiro para ver seus divisores: ")
	fmt.Scan(&num)
	
	for i := 1; i <= num; i++ {
		if num % i == 0 {
			if i < num {
				divisors += fmt.Sprintf("%s, ", strconv.Itoa(i))
			} else {
				divisors += fmt.Sprintf("%s", strconv.Itoa(i))
			}
			
		}
	}
	fmt.Printf("Os divisores de %d são: %s\n",num, divisors)
}
