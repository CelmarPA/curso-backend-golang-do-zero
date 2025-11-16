// Exercício 0046: Faça um programa que mostre na tela uma contagem regressiva para o estouro de fogos de artifício, indo de 10 até 0, com uma pausa de 1 segundo entre eles.

//go:build ignore

package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 10; i >= 0; i-- {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("BUM BUM BUM BUM * * * *!!!")
}
