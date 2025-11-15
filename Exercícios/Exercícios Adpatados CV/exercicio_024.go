// Exercício 024: Leia o nome de uma cidade e diga se ela começa com 'Santo'. 

//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Digite nome de uma cidade: ")
	reader := bufio.NewReader(os.Stdin)
	city, _ := reader.ReadString('\n')
	city = strings.TrimSpace(city)

	if strings.ToLower(city[:5]) == "santo" {
		fmt.Println("O nome da cidade começa com Santo!")
	}
}
