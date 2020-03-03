//go:generate stringer -type=Suit,Rank

package deck

//Suit ...
type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
)

//Rank ...
type Rank uint8

const (
	Ace Rank = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

//Card ...
type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	return "Ace of Hearts"
}
