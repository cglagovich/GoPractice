package deck

import (
  "testing"
  "math/rand"
)
func TestExampleCard(t *testing.T) {
  AceHeart := Card{Rank: Ace, Suit: Heart}.String()
  if AceHeart != "Ace of Hearts" {
    t.Errorf("Failed")
  }

  TwoSpade := Card{Rank: Two, Suit: Spade}.String()
  if TwoSpade != "Two of Spades" {
    t.Errorf("Failed")
  }

  NineDiamond := Card{Rank: Nine, Suit: Diamond}.String()
  if NineDiamond != "Nine of Diamonds" {
    t.Errorf("Failed")
  }

  JackClub := Card{Rank: Jack, Suit: Club}.String()
  if JackClub != "Jack of Clubs" {
    t.Errorf("Failed")
  }

  Joker := Card{Suit: Joker}.String()
  if Joker != "Joker" {
    t.Errorf("Failed")
  }
}

func TestNew (t *testing.T) {
  cards := New()
  if len(cards) != 52 {
    t.Errorf("Card deck wrong size")
  }
}

func TestDefaultSort(t *testing.T) {
  cards := New(DefaultSort)
  exp := Card{Rank: Ace, Suit: Spade}
  if cards[0] != exp {
    t.Error("Expected Ace of Spades as first card. Received:", cards[0])
  }
}

func TestJokers(t *testing.T) {
  cards := New(Jokers(3))
  count := 0
  for _, c:= range cards {
    if c.Suit == Joker {
      count++
    }
  }
  if count != 3 {
    t.Error("Expected 3 Jokers, received:", count)
  }
}

func TestFilter(t *testing.T) {
  filter := func(card Card) bool {
    return card.Rank == 2 || card.Rank == Three
  }
  cards := New(Filter(filter))
  for _, c := range cards {
    if c.Rank == Two || c.Rank == Three {
      t.Error("Expected all twos and threes to be filtered")
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

func TestShuffle(t *testing.T) {
  // make shuffleRand deterministic
  // First call to shufleRand.Perm(52) should be:
  // [40 35 ... ]
  shuffleRand = rand.New(rand.NewSource(0))
  orig := New()
  first := orig[40]
  second := orig[35]
  cards := New(Shuffle)
  if cards[0] != first {
    t.Errorf("Expected the first card to be %s, received %s", first, cards[0])
  }

  if cards[1] != second {
    t.Errorf("Expected the second card to be %s, received %s", second, cards[1])
  }

}
