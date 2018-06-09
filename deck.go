package deck

import (
	"math/rand"
	"sort"
	"time"
)

type Deck []Card

func New(opts ...func([]Card) []Card) Deck {
	var cards []Card

	for _, s := range suits {
		for r := minRank; r <= maxRank; r++ {
			cards = append(cards, Card{Suit: s, Rank: r})
		}
	}

	for _, o := range opts {
		cards = o(cards)
	}

	return cards
}

func Sort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

// Shuffle the cards using Fisher–Yates algorithm.
func Shuffle(cards []Card) []Card {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i := len(cards) - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}

	return cards
}

func Jokers(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{Rank: Rank(i), Suit: Joker})
		}
		return cards
	}
}

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}
