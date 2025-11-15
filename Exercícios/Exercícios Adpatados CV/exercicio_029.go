// Exercício 029: Escreva um programa que leia a velocidade de um carro. Se ele ultrapassar 80Km/h, mostre uma 
// mensagem dizendo que ele foi multado. A multa vai custar R$7,00 por cada Km/h acima do limite.

//go:build ignore

package main

import (
	"fmt"
)

func main() {
	var speed int

	fmt.Print("Digite a velocidade do veículo em Km/h: ")
	fmt.Scan(&speed)

	if speed > 80 {
		fine := float64((speed - 80) * 7)
		fmt.Printf("Você ultrapassou o limite de velocidade de 80 Km/h, será multado em R$%.2f!\n", fine)
		return
	} 

	fmt.Println("Você está dentro do limite de velocidade permitido. Continue viajando com segurança!")
}
