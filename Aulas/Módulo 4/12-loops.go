//go:build ignore

package main

import (
	"fmt"
	"time"
)

// Loops
// Laços de repetição
// Repetir tarefas

func main() {

	sum := 0
	// i++ -> soma 1 -> i = i + 1
	// i-- -> subtrai 1
	for i := 0; i < 10; i++ {
		fmt.Println("Sum:", sum)
		fmt.Println("Índice:", i)

		sum += i
		// é a mesma coisa que: sum = sum + i
		// sum -= i -> sum = sum - i
	}
	fmt.Println(sum)

	// loop infinito
	for {
		fmt.Println("Golang do zero")
		time.Sleep(2 * time.Second)
		break
	}
	
	// for range
	frutas := []string {"laranja", "maça", "banana", "uva", "kiwi"}
	for _, fruta := range frutas {
		fmt.Println("Fruta", fruta)
	}
}
