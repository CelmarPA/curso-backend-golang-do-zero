// Exercício 022: Faça um programa que leia uma frase e conte quantas letras 'a' aparecem, a posição da primeira e da última ocorrência

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

	fmt.Print("Digite uma frase para contar quantas letras 'a's ela possui e a primeira e última ocorrência: ")
	reader := bufio.NewReader(os.Stdin)
	phrase, _ := reader.ReadString('\n')
	phrase = strings.TrimSpace(phrase)

	
	for _, letter := range(phrase) {
		if strings.ToLower(string(letter)) == "a" {
			count++
		}
	}

	firstIndex := strings.Index(phrase, "a")
	lastIndex := strings.LastIndex(phrase, "a")

	fmt.Printf("A frase %s; possui %d letras 'a's. Primeira ocorrência no índice %d e a última no índice %d \n", phrase, count, firstIndex, lastIndex)
}
