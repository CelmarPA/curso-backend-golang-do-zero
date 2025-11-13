// Exercício 026: Leia uma frase e mostre-a em maiúsculas e minusculas.

//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Digite uma frase: ")
	
	reader := bufio.NewReader(os.Stdin)
	phrase, _ := reader.ReadString('\n')
	phrase = strings.TrimSpace(phrase)

	fmt.Printf("A frase em maiúsculas: %s\n", strings.ToUpper(phrase))
	fmt.Printf("A frase em minúsculas: %s\n", strings.ToLower(phrase))
}
