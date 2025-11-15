// Exercício 019: Sorteie um dos quatro nomes informados pelo usuário utilizando slices e geração de números aleatórios. 

//go:build ignore

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var name string
	var names [4]string
	
	for i := 0; i < 4; i++ {
		fmt.Printf("Digite o %dº nome: ", i + 1)
		fmt.Scan(&name)	
		names[i] = name
	}

	index := rand.Intn(4)
	
	fmt.Printf("O nome sorteado foi: %s\n", names[index])
}
