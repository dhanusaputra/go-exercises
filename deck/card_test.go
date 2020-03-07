package deck

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Spade})
	fmt.Println(Card{Rank: Nine, Suit: Diamond})
	fmt.Println(Card{Rank: Jack, Suit: Club})
	fmt.Println(Card{Suit: Joker})

	//Output:
	//Ace of Hearts
	//Two of Spades
	//Nine of Diamonds
	//Jack of Clubs
	//Joker
}

func TestNew(t *testing.T) {
	cards := New()
	want := 13 * 4
	if len(cards) != want {
		t.Errorf("wrong number of cards got %v, want %v", len(cards), want)
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	if want := (Card{Rank: Ace, Suit: Spade}); !reflect.DeepEqual(cards[0], want) {
		t.Error("Expected ", want, " as first card. Got: ", cards[0])
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	if want := (Card{Rank: Ace, Suit: Spade}); !reflect.DeepEqual(cards[0], want) {
		t.Error("Expected ", want, " as first card. Got: ", cards[0])
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(3))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}
	if count != 3 {
		t.Error("Expected 3 Jokers, Got: ", count)
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}
	cards := New(Filter(filter))
	for _, c := range cards {
		if c.Rank == Two || c.Rank == Three {
			t.Error("Expected all twos and threes to be filtered out.")
		}
	}
}

func TestDeck(t *testing.T) {
	cards := New(Deck(3))
	if len(cards) != 13*4*3 {
		t.Errorf("Expected %d, got: %d", 13*4*4, len(cards))
	}
}

func TestShuffle(t *testing.T) {
	shuffleRand = rand.New(rand.NewSource(0))
	orig := New()
	cards := New(Shuffle)
	if first := orig[40]; cards[0] != first {
		t.Errorf("Expecetd %v, got %v", first, cards[0])
	}
	if second := orig[35]; cards[1] != second {
		t.Errorf("Expecetd %v, got %v", second, cards[1])
	}

}
