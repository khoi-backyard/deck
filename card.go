//go:generate stringer -type=Suit,Rank
package deck

import "fmt"

type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

type Rank uint8

const (
	_ Rank = iota
	Ace
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

const (
	minRank = Ace
	maxRank = King
)

type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return "🃏"
	}
	return fmt.Sprintf("%s of %ss", c.Rank, c.Suit)
}

func absRank(c Card) int {
	return int(c.Suit) * int(maxRank) + int(c.Rank)
}