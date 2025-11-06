// Exercício 004: Leia uma variável e mostre seus caracteres separados (ex.: 'abc' -> 'a b c').

//go:build ignore

package main

import (
	"fmt"
	"strings"
)

// Forma mais compacta
// func main() {
// 	var word string
// 	fmt.Print("Digite uma palavra: ")
// 	fmt.Scanln(&word)
// 	fmt.Println(strings.Join(strings.Split(word, ""), " "))
// }


func main () {
	var word string
	fmt.Print("Digite uma palavra: ")
	fmt.Scanln(&word)

	// Quebra a string em runas (suporta acentos e emojis)
	var chars []string	
	for _, letter := range word {
		chars = append(chars, string(letter))
	}
	
	fmt.Println(strings.Join(chars, " "))
}

