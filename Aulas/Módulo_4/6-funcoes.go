//go:build ignore

package main

import "fmt"

func main() {
	somaDosValores := soma(42, 13)
	fmt.Println(somaDosValores)

	sub := subtracao(10, 5)
	fmt.Println(sub)

	nome, sobrenome := printNomeCompleto("Celmar", "Pereira")
	fmt.Println(nome, sobrenome)
}

// Função começando com letra minúscula:
// Função Privada
// Essa Função só poderá ser utilizada no próprio pacote 
func printNomeCompleto(nome, sobrenome string) (string, string) {
	return nome, sobrenome
}

func soma(x int, y int) int {
	// somaDosValores := x + y
	// return somaDosValores
	return x + y
}

func subtracao(x int, y int) int {
	return x - y
}

// Função começando com letra maiúscula
// Função Pública
// Essa Função poderá ser utilizada fora do próprio pacote
// Como utilizar fora do pacote: main.PrintNome(nome)
func PrintName(nome string) string {
	return nome
}
