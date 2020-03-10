package aiblackjack

import (
	"fmt"

	"github.com/dhanusaputra/go-exercises/deck"
)

//AI ...
type AI interface {
	Bet() int
	Play(hand []deck.Card, dealer deck.Card) Move
	Results(hand [][]deck.Card, dealer []deck.Card)
}

type dealerAI struct{}

func (ai *dealerAI) Bet() int {
	//noop
	return 1
}

func (ai *dealerAI) Play(hand []deck.Card, dealer deck.Card) Move {
	dScore := Score(hand...)
	if dScore <= 16 || (dScore == 17 && Soft(hand...)) {
		return MoveHit
	} else {
		return MoveStand
	}
}

//Results ...
func (ai *dealerAI) Results(hand [][]deck.Card, dealer []deck.Card) {
	//noop
}

//HumanAI ...
type HumanAI struct{}

//Bet ...
func (ai *HumanAI) Bet() int {
	return 1
}

//Play ...
func (ai *HumanAI) Play(hand []deck.Card, dealer deck.Card) Move {
	for {
		fmt.Println("Player:", hand)
		fmt.Println("Dealer:", dealer)
		fmt.Println("What will you do? (h)it, (s)tand")
		var input string
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			return MoveHit
		case "s":
			return MoveStand
		default:
			fmt.Println("Invalid option:", input)
		}
	}
}

//Results ...
func (ai *HumanAI) Results(hand [][]deck.Card, dealer []deck.Card) {
	fmt.Println("==FINAL HANDS==")
	fmt.Println("Player:", hand)
	fmt.Println("Dealer:", dealer)
}
