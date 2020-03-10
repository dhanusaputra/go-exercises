package aiblackjack

import (
	"fmt"

	"github.com/dhanusaputra/go-exercises/deck"
)

//State ...
type state uint8

//State details ...
const (
	statePlayerTurn state = iota
	stateDealerTurn
	stateHandOver
)

//Game ...
type Game struct {
	//unexported fields
	deck     []deck.Card
	state    state
	player   []deck.Card
	dealer   []deck.Card
	dealerAI dealerAI
}

func (g *Game) currentPlayer() []deck.Card {
	switch g.state {
	case statePlayerTurn:
		return g.player
	case stateDealerTurn:
		return g.dealer
	default:
		panic("it isn't currently any player's turn")
	}
}

func deal(g *Game) {
	g.player = make([]deck.Card, 0, 5)
	g.dealer = make([]deck.Card, 0, 5)
	var card deck.Card
	for i := 0; i < 2; i++ {
		card, g.player = draw(g.deck)
		g.player = append(g.player, card)
		card, g.deck = draw(g.deck)
		g.dealer = append(g.dealer, card)
	}
	g.state = statePlayerTurn
}

//Play ...
func (g *Game) Play(ai AI) {
	g.deck = deck.New(deck.Deck(3), deck.Shuffle)

	for i := 0; i < 10; i++ {
		deal(g)

		var input string
		for g.state == statePlayerTurn {
			hand := make([]deck.Card, len(g.player))
			copy(hand, g.player)
			move := ai.Play(hand, g.dealer[0])
			move(g)
		}

		for g.state == stateDealerTurn {
			hand := make([]deck.Card, len(g.player))
			copy(hand, g.dealer)
			move := g.dealerAI.Play(hand, g.dealer[0])
			move(g)
		}

		endHand(g, ai)
	}
}

//Move ...
type Move func(*Game)

//MoveHit ...
func MoveHit(g *Game) {
	hand := g.currentPlayer()
	var card deck.Card
	card, g.deck = draw(g.deck)
	hand = append(hand, card)
	if Score(hand...) > 21 {
		MoveStand(g)
	}
}

//MoveStand ...
func MoveStand(g *Game) {

}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

//EndHand ...
func endHand(g *Game, ai AI) {
	pScore, dScore := Score(g.player...), Score(g.dealer...)
	switch {
	case pScore > 21:
		fmt.Println("You busted")
	case dScore > 21:
		fmt.Println("Dealer busted")
	case pScore > dScore:
		fmt.Println("You win")
	case pScore < dScore:
		fmt.Println("You lose")
	case pScore == dScore:
		fmt.Println("Draw")
	}
	fmt.Println()
	ai.Results([][]deck.Card{g.player}, g.dealer)
	g.player = nil
	g.dealer = nil
}

//Score ...
func Score(hand ...deck.Card) int {
	minScore := minScore(hand...)
	if minScore > 11 {
		return minScore
	}
	for _, c := range hand {
		if c.Rank == deck.Ace {
			return minScore + 10
		}
	}
	return minScore
}

//Soft ...
func Soft(hand ...deck.Card) bool {
	minScore := minScore(hand...)
	score := Score(hand...)
	return minScore != score
}

//MinScore ...
func minScore(hand ...deck.Card) int {
	var score int
	for _, c := range hand {
		score += min(int(c.Rank), 10)
	}
	return score
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
