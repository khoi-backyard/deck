package deck

type Deck []Card

func New() Deck {
	var cards []Card

	for _, s := range suits {
		for r := minRank; r <= maxRank; r++ {
			cards = append(cards, Card{Suit: s, Rank: r})
		}
	}

	return cards
}
