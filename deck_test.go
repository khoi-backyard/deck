package deck

import "testing"

func TestNew(t *testing.T) {
	deck := New()

	if cardCount := len(deck); cardCount != 52 {
		t.Errorf("Deck should have 52 cards, got %d", cardCount)
	}
}
