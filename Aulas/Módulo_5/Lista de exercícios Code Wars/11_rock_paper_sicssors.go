// Rules of the "Rock, Paper, Scissors" game are:

// Rock beats Scissors
// Scissors beat Paper,
// Paper beats Rock.
// Let's play! You have to return which player won! In case of a draw return Draw!.

// Examples:
// "scissors", "paper" --> "Player 1 won!"
// "scissors", "rock" --> "Player 2 won!"
// "paper", "paper" --> "Draw!"

//go:build ignore

package main

import (
	"fmt"
)

func Rps(p1, p2 string) string {
	if p1 == "rock" && p2 == "scissors" || p1 == "paper" && p2 == "rock" || p1 == "scissors" && p2 == "paper" {
		return "Player 1 won!"
	} else if p2 == "rock" && p1 == "scissors" || p2 == "paper" && p1 == "rock" || p2 == "scissors" && p1 == "paper" {
		return "Player 2 won!"
	}
	return "Draw!"
}

func main() {
	fmt.Println(Rps("rock","scissors"))
	fmt.Println(Rps("scissors","rock"))
	fmt.Println(Rps("rock","rock"))
}