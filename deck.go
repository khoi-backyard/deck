package deck

import "sort"

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

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}
