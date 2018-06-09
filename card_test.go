package deck

import "fmt"

func ExampleCard_String() {
	fmt.Println(Card{Suit: Joker})
	fmt.Println(Card{Heart, Ace})
	fmt.Println(Card{Diamond, Two})
	fmt.Println(Card{Spade, Three})
	fmt.Println(Card{Club, Four})
	fmt.Println(Card{Heart, Five})
	fmt.Println(Card{Diamond, Six})
	fmt.Println(Card{Spade, Seven})
	fmt.Println(Card{Club, Eight})
	fmt.Println(Card{Heart, Nine})
	fmt.Println(Card{Diamond, Ten})
	fmt.Println(Card{Spade, Jack})
	fmt.Println(Card{Club, Queen})
	fmt.Println(Card{Heart, King})

	// Output:
	// 🃏
	// Ace of Hearts
	// Two of Diamonds
	// Three of Spades
	// Four of Clubs
	// Five of Hearts
	// Six of Diamonds
	// Seven of Spades
	// Eight of Clubs
	// Nine of Hearts
	// Ten of Diamonds
	// Jack of Spades
	// Queen of Clubs
	// King of Hearts
}
