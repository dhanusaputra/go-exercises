package main

import (
	"fmt"

	"github.com/dhanusaputra/go-exercises/aiblackjack/blackjack"
)

func main() {
	game := blackjack.New()
	winnings := game.Play(blackjack.HumanAI())
	fmt.Println(winnings)
}
