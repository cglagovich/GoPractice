package main

import (
  "fmt"
  "github.com/cglagovich/GoPractice/deck"
  "strings"
)

type Hand []deck.Card

func (h Hand) String() string {
  strs := make([]string, len(h))
  for i := range h {
    strs[i] = h[i].String()
  }
  return strings.Join(strs,", ")
}

func (h Hand) DealerString() string {
  return h[0].String() + ", **HIDDEN**"
}

func (h Hand) Score() int {
  minScore := h.MinScore()
  if (minScore > 11) {
    return minScore
  }
  for _, c:= range h {
    if c.Rank == deck.Ace{
      // Ace is worth 1, and we are changing it to be worth 11
      // Only one Ace can be changed to 11 without busting
      return minScore + 10;
    }
  }
  return minScore
}

func (h Hand) MinScore() int {
  score := 0
  for _, c := range h {
    score += min(int(c.Rank), 10)
  }
  return score
}

func min(a,b int) int {
  if a < b {
    return a
  }
  return b
}

func main() {
  cards := deck.New(deck.Deck(3), deck.Shuffle)
  var card deck.Card
  var player, dealer Hand
  // Dealing. Use a slice of pointers to hands that are the player and dealer's
  // hands so the for loop can modify their Hands.
  for i := 0; i < 2; i++ {
    for _, hand := range []*Hand{&player, &dealer} {
      card, cards = draw(cards)
      *hand = append(*hand, card)
    }
  }
  var input string
  for input != "s" {
    fmt.Println("Player:", player)
    fmt.Println("Dealer:", dealer.DealerString())
    fmt.Println("What now? (h)it, (s)tand")
    fmt.Scanf("%s\n", &input)
    switch input {
    case "h":
      card, cards = draw(cards)
      player = append(player, card)
    default:
      fmt.Println("That is not a valid option. Try again")
    }
  }
  // Strategy for the dealer:
  // If dealer score <= 16, hit
  // If dealer has soft 17, hit -- aka score==17, minscore==7
  for dealer.Score() <= 16 || (dealer.Score() == 17 && dealer.MinScore() != 17) {
    card, cards = draw(cards)
    dealer = append(dealer, card)
  }

  pScore, dScore := player.Score(), dealer.Score()
  fmt.Println("----Final Hands----")
  fmt.Println("Player:", player, "\nScore:", pScore)
  fmt.Println("Dealer:", dealer, "\nScore:", dScore)
  switch {
  case pScore > 21:
    fmt.Println("You busted")
  case dScore > 21:
    fmt.Println("Dealer busted")
  case pScore > dScore:
    fmt.Println("You win :)")
  case dScore > pScore:
    fmt.Println("Dealer won :(")
  case dScore == pScore:
    fmt.Println("Draw")
  }
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
    return cards[0], cards[1:]
}
