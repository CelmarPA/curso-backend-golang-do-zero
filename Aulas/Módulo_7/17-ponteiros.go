//go:build ignore

package main

import "fmt"

// Ponteiros: Um ponteiro nada mais é do que uma variável que ao invés
// de armazenar um valor (1, "olá", true..), armazena o endereço de memória.

func zeroValue(i int) {
	i = 0
	fmt.Println("Endereço de memória do i dentro da função:", &i)
}

func zeroPointer(i *int) {
	*i = 0
	fmt.Println("Endereço de memória do i dentro da função:", &i)
}

func main() {
	// Memória -> essa memória tem um endereço -> esse endereço guarda um valor

	i := 1
	fmt.Println("Valor inicial: ", i)
	fmt.Println("Valor endereço de memória:", &i) // &var -> apontando para o endereço de memória

	zeroValue(i)
	fmt.Println("zeroVal:", i)

	zeroPointer(&i)
	fmt.Println("zeroPointer:", i)

	fmt.Println("pointer:", &i)

	a := &i
	fmt.Println("Variável a:", a)
	fmt.Println("Variável *a:", *a) // *var -> pega o valor contido no endereço apontando por a, ou seja, pega endereço e acessa o valor
	fmt.Println("Variável &a:", &a)
}
