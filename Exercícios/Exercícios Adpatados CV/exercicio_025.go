// Exercício 025: Leia um nome e diga se contém 'Silva' (substring).

//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var contain bool

	fmt.Print("Digite um nome: ")
	
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	splitName := strings.Split(name, " ")

	for _, part := range(splitName) {
		part = strings.ToLower(part)
		if part == "silva" {
			contain = true
			break
		}
	}

	if contain {
		fmt.Println("O nome possui Silva!")
	} else {
		fmt.Println("O nome não possui Silva!")
	}
}
