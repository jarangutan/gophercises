package deck

import (
	"math/rand"
	"testing"
)

func TestCardToString(t *testing.T) {
	var tests = []struct {
		input Card
		want  string
	}{
		{Card{Suit: Spade, Rank: Ace}, "Ace of Spades"},
		{Card{Suit: Heart, Rank: Five}, "Five of Hearts"},
		{Card{Suit: Club, Rank: Jack}, "Jack of Clubs"},
		{Card{Suit: Diamond, Rank: King}, "King of Diamonds"},
		{Card{Suit: Joker}, "Joker"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			cardName := tt.input.String()
			if cardName != tt.want {
				t.Errorf("got \"%s\", want \"%s\"", cardName, tt.want)
			}
		})
	}

}

func TestNew(t *testing.T) {
	cards := New()
	cardLen := len(cards)
	wantLen := 13 * 4 // rank * suit
	if cardLen != wantLen {
		t.Errorf("Wrong number of cards in deck. Got %d, want %d", cardLen, wantLen)
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	want := Card{Suit: Spade, Rank: Ace}

	if cards[0] != want {
		t.Errorf("Wrong first card. Got \"%s\", want \"%s\"", cards[0], want)
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	want := Card{Suit: Spade, Rank: Ace}

	if cards[0] != want {
		t.Errorf("Wrong first card. Got \"%s\", want \"%s\"", cards[0], want)
	}
}

func TestJokers(t *testing.T) {
	want := 5
	cards := New(Jokers(want))

	got := 0
	for _, c := range cards {
		if c.Suit == Joker {
			got++
		}
	}

	if got != want {
		t.Errorf("Wrong number of jokers. Got \"%d\", want \"%d\"", got, want)
	}
}

func TestFilter(t *testing.T) {
	cards := New(Filter(func(card Card) bool {
		return card.Suit == Spade
	}))

	for _, c := range cards {
		if c.Suit == Spade {
			t.Errorf("Filtered out card found. Got \"%s\", want no \"%s\" cards", c, Spade)
		}
	}
}

func TestDeck(t *testing.T) {
	cards := New(Deck(3))
	want := 13 * 4 * 3 // rank * suit * deck
	got := len(cards)

	if got != want {
		t.Errorf("Deck length incorrect. Got %d, want %d", got, want)
	}
}

func TestShuffle(t *testing.T) {
	// firstCall to shuffleRand.Perm() should be:
	// [40 35 ...]
	shuffleRand = rand.New(rand.NewSource(0))

	orig := New()
	first := orig[40]
	second := orig[35]
	cards := New(Shuffle)

	if cards[0] != first {
		t.Errorf("Cards not shuffled right. Got %s, want first card to be %s", cards[0], first)
	}
	if cards[1] != second {
		t.Errorf("Cards not shuffled right. Got %s, want second card to be %s", cards[1], second)
	}

}

// TODO: Not sure how // Output: below works but it's the same as above
//
// func ExampleCard() {
// 	fmt.Println(Card{Suit: Heart, Rank: Ace})
//
// 	// Output:
// 	// Ace of Hearts
// }
