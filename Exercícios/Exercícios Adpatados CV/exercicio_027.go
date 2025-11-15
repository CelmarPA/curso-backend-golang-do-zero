// Exercício 027: Ler o nome completo de uma pessoa e mostrar apenas o primeiro e o último nome.

//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	print("Digite seu nome completo: ")

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	splitName := strings.Split(name, " ")

	if len(splitName) < 2 {
		fmt.Println("Você digitou apenas o primeiro nome!")
		return
	}

	fmt.Printf("Seu primeiro nome é %s e seu último nome é %s!\n", splitName[0], splitName[len(splitName) - 1])
}
