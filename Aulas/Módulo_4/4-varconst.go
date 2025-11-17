//go:build ignore

package main

import "fmt"

func main() {
	// Variáveis

	// var + nome_da_variável + tipo
	var nome string
	nome = "bento"
	fmt.Println(nome)

	nome = "Celmar"
	fmt.Println(nome)

	var idade int
	idade = 4
	fmt.Println(idade)

	var a = "Celmar"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	numero := 1
	fmt.Println(numero)

	// Constantes

	const idadeBento = 5
	fmt.Println(idadeBento)
}