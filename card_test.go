package deckofcards

import (
	"fmt"
	"math/rand"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Jack, Suit: Spade})
	fmt.Println(Card{Rank: Nine, Suit: Club})
	fmt.Println(Card{Rank: Queen, Suit: Diamond})
	fmt.Println(Card{Rank: King, Suit: Heart})

	// Output:
	// Ace of Hearts
	// Jack of Spades
	// Nine of Clubs
	// Queen of Diamonds
	// King of Hearts
}

func TestNew(t *testing.T) {
	cards := New()
	if len(cards) != 13*4 {
		t.Errorf("Error with new cards")
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp {
		t.Error("Expected Ace of Spades as first card. Received:", cards[0])
	}
}

func TestShuffle(t *testing.T) {
	// make shuffleRand deterministic
	// First call to shuffleRand.Perm(52) should be:
	// [40 35 ... ]
	shuffleRand = rand.New(rand.NewSource(0))
	// shuffleRand.Perm(52)

	orig := New()
	first := orig[40]
	second := orig[35]
	cards := New(Shuffle)
	if cards[0] != first {
		t.Errorf("Expected the first card to be %s, received %s.", first, cards[0])
	}
	if cards[1] != second {
		t.Errorf("Expected the first card to be %s, received %s.", second, cards[1])
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(4))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}
	if count != 4 {
		t.Error("Expected 4 Jokers, received:", count)
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
	// 13 ranks * 4 suits * 3 decks
	if len(cards) != 13*4*3 {
		t.Errorf("Expected %d cards, received %d cards.", 13*4*3, len(cards))
	}
}
