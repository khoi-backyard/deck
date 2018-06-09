package deck

import "testing"

func TestNew(t *testing.T) {
	deck := New()

	if cardCount := len(deck); cardCount != 52 {
		t.Errorf("Deck should have 52 cards, got %d", cardCount)
	}
}

func TestDefaultSort(t *testing.T) {
	deck := New(DefaultSort)

	first := Card{Rank: Ace, Suit: Spade}
	last := Card{Rank: King, Suit: Heart}

	if c := deck[0]; c != first {
		t.Errorf("Expecting %s, got %s", first, c)
	}

	if c := deck[len(deck)-1]; c != last {
		t.Errorf("Expecting %s, got %s", first, c)
	}
}

func TestSort(t *testing.T) {
	deck := New(Sort(Less))
	exp := New(DefaultSort)

	if len(deck) != len(exp) {
		t.Errorf("Deck count: %d, Expected count: %s", len(deck), len(exp))
	}

	for i := 0; i < len(deck); i++ {
		if deck[i].Suit != exp[i].Suit || deck[i].Rank != exp[i].Rank {
			t.Errorf("Expected %s, got %s", exp[i], deck[i])
		}
	}
}
