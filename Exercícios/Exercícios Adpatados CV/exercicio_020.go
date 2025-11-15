// 020. Exercício 020: Leia quatro nomes e gere uma ordem aleatória.

//go:build ignore

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var name string
	var names [4]string
	
	for i := 0; i < 4; i++ {
		fmt.Printf("Digite o %dº nome: ", i + 1)
		fmt.Scan(&name)	
		names[i] = name
	}

	rand.Shuffle(len(names), func(i, j int) {
		names[i], names[j] = names[j], names[i]
	})

	
	fmt.Printf("Os nomes em  ordem aleatória são: %s\n", names)
}
