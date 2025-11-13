// Exercício 024: Leia uma frase e conte quantas letras 'a' ela contém (case-insensitive).

//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var count int

	fmt.Print("Digite uma frase para contar quantas letras 'a's ela possui: ")
	reader := bufio.NewReader(os.Stdin)
	phrase, _ := reader.ReadString('\n')
	phrase = strings.TrimSpace(phrase)

	
	for _, letter := range(phrase) {
		if strings.ToLower(string(letter)) == "a" {
			count++
		}
	}
	
	fmt.Printf("A frase %s; possui %d letras 'a's.\n", phrase, count)
}
