//go:generate stringer -type=Suit,Rank

package deck

import (
  "fmt"
  "sort"
  "math/rand"
  "time"
)

type Suit uint8

const (
  // Constants to enumerate suits with integers.
  // Iota gives the first const in a const list value 0, the next 1...
  Spade Suit = iota
  Diamond
  Club
  Heart
  Joker // Special case of suit
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
    return c.Suit.String()
  }
  return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

func New(opts ...func([]Card) []Card) []Card {
  var cards []Card
  // For each suit
  for _, suit := range suits {
    // For each rank
    for rank := minRank; rank <= maxRank; rank++ {
        // Add card to cards
        cards = append(cards, Card{Suit: suit, Rank: rank})
    }
  }
  for _, opt := range opts {
    cards = opt(cards)
  }
  return cards
}

func DefaultSort(cards []Card) []Card {
  sort.Slice(cards, Less(cards))
  return cards
}

func Sort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
  return func(cards []Card) []Card {
    sort.Slice(cards, less(cards))
    return cards
  }
}

func Less(cards []Card) func(i, j int) bool {
  return func(i, j int) bool {
    return absRank(cards[i]) < absRank(cards[j])
  }
}

func absRank(c Card) int {
  return int(c.Suit) * int(maxRank) + int(c.Rank)
}

type Permer interface {
  Perm(n int) []int
}

var shuffleRand = rand.New(rand.NewSource(time.Now().Unix()))
func Shuffle(cards []Card) []Card {
  // This shuffle method will make a new slice to shuffle in. Could be shuffled
  // in place, but this will be easier.
  ret := make([]Card, len(cards))
  perm := shuffleRand.Perm(len(cards))
  for i, j := range perm {
    // i is index in perm int slice, j is value at index i
    ret[i] = cards[j]
  }
  return ret
}

// Option for New function.. called like New(Jokers(3))
func Jokers(n int) func([]Card) []Card {
  return func(cards []Card) []Card {
    for i := 0; i < n; i++ {
      cards = append(cards, Card{
        Rank: Rank(i),
        Suit: Joker,
      })
    }
    return cards
  }
}

// Filter takes a function f as an argument. This function returns
// a boolean given a card if it should be in the deck.
// Filter returns a function on a Card slice that returns a filtered
// card slice.
// This way, it will be passed as an option to New() that will be run
// on a new deck to filter it.
func Filter(f func(card Card) bool) func([]Card) []Card {
  return func(cards []Card) []Card {
    var ret []Card
    for _, c := range cards {
      if !f(c) {
        ret = append(ret, c)
      }
    }
    return ret
  }
}

func Deck(n int) func([]Card) []Card {
  return func(cards []Card) []Card {
    var ret []Card
    for i := 0; i < n; i++ {
        ret = append(ret, cards...)
    }
    return ret
  }
}
