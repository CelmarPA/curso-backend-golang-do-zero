// Exercício 015: Leia o nome completo e mostre apenas o primeiro nome.

//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Digite seu nome completo: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	
	fmt.Printf("Seu primeiro nome é %s.\n", strings.Split(name, " ")[0])
}
