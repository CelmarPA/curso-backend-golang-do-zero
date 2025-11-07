// Exercício 004: Leia uma variável e mostre seus caracteres separados (ex.: 'abc' -> 'a b c').

//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Forma mais compacta
// func main() {
// 	fmt.Print("Digite uma palavra: ")
// 	reader := bufio.NewReader(os.Stdin)
// 	word, _ := reader.ReadString('\n')
// 	word = strings.TrimSpace(word)
// 	fmt.Println(strings.Join(strings.Split(word, ""), " "))
// }


func main () {
	fmt.Print("Digite uma palavra: ")
	reader := bufio.NewReader(os.Stdin)
	word, _ := reader.ReadString('\n')
	word = strings.TrimSpace(word)

	// Quebra a string em runas (suporta acentos e emojis)
	var chars []string	
	for _, letter := range word {
		chars = append(chars, string(letter))
	}
	
	fmt.Println(strings.Join(chars, " "))
}

