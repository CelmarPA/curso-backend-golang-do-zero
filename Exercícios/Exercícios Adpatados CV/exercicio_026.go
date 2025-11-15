// Exercício 026: Leia uma frase e mostre a posição da primeira e última ocorrência da palavra 'golang' (ou outra).

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

	firstIndex := strings.Index(strings.ToLower(phrase), "golang")
	lastIndex := strings.LastIndex(strings.ToLower(phrase), "golang")

	if firstIndex != -1{
		fmt.Printf("Na frase '%s' a palavra golang pela primeira vez na posição %d e pela última vez na posição %d!\n", phrase, firstIndex, lastIndex)
	} else {
		fmt.Println("A palavra golang não aparece na frase!")
	}
}
