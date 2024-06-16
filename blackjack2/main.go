package main

import (
	"fmt"
	"strings"

	"github.com/jarangutan/gophercises/blackjack2/deck"
)

type Hand []deck.Card

func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := range h {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ", ")
}

func (h Hand) DealerString() string {
	return h[0].String() + ", **HIDDEN**"
}

func (h Hand) Score() int {
	minScore := h.MinScore()
	if minScore > 11 {
		return minScore
	}
	for _, c := range h {
		if c.Rank == deck.Ace {
			// Ace is worth 1 point, this changes to worth to 11
			return minScore + 10
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Shuffle(gs GameState) GameState {
	ret := clone(gs)
	ret.Deck = deck.New(deck.Deck(3), deck.Shuffle)
	return ret
}

func Deal(gs GameState) GameState {
	ret := clone(gs)
	ret.Player = make(Hand, 0, 6)
	ret.Dealer = make(Hand, 0, 6)
	var card deck.Card
	for i := 0; i < 2; i++ {
		for _, hand := range []*Hand{&ret.Player, &ret.Dealer} {
			card, ret.Deck = draw(ret.Deck)
			*hand = append(*hand, card)
		}
	}
	ret.State = StatePlayerTurn
	return ret
}

func Hit(gs GameState) GameState {
	ret := clone(gs)
	hand := ret.CurrentPlayer()
	var card deck.Card
	card, ret.Deck = draw(ret.Deck)
	*hand = append(*hand, card)
	if hand.Score() > 21 {
		return Stand(gs)
	}
	return ret
}

func Stand(gs GameState) GameState {
	ret := clone(gs)
	switch ret.State {
	case StatePlayerTurn:
		ret.State = StateDealerTurn
	case StateDealerTurn:
		ret.State = StatePlayerTurn
	}
	return ret
}

func EndHand(gs GameState) GameState {
	ret := clone(gs)
	pScore, dScore := ret.Player.Score(), ret.Dealer.Score()
	fmt.Println("==FINAL HAND==")
	fmt.Println("Dealer:", ret.Dealer, "\nScore:", dScore)
	fmt.Println("Player:", ret.Player, "\nScore:", pScore)
	switch {
	case pScore > 21:
		fmt.Println("You busted")
	case dScore > 21:
		fmt.Println("Dealer busted")
	case pScore > dScore:
		fmt.Println("You win!")
	case dScore > pScore:
		fmt.Println("You lose!")
	case dScore == pScore:
		fmt.Println("Draw")
	}
	fmt.Println()

	ret.Player = nil
	ret.Dealer = nil
	return ret
}

func main() {
	var gs GameState
	gs = Shuffle(gs)
	for {

		gs = Deal(gs)

		var input string
		for gs.State == StatePlayerTurn {
			fmt.Println("Dealer", gs.Dealer.DealerString())
			fmt.Println("Player", gs.Player)
			fmt.Println("What will you do? (h)it, (s)tand")
			fmt.Scanf("%s\n", &input)
			switch input {
			case "h":
				gs = Hit(gs)
			case "s":
				gs = Stand(gs)
			default:
				fmt.Println("Invalid option:", input)
			}
		}

		for gs.State == StateDealerTurn {
			if gs.Dealer.Score() <= 16 || (gs.Dealer.Score() == 17 && gs.Dealer.MinScore() != 17) {
				gs = Hit(gs)
			} else {
				gs = Stand(gs)
			}
		}

		gs = EndHand(gs)
		fmt.Println("(n)ew game? (q)uit?")
		fmt.Scanf("%s\n", &input)
		switch input {
		case "n":
			continue
		case "q":
			return
		}
	}
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

type State int8

const (
	StatePlayerTurn State = iota
	StateDealerTurn
	StateHandOver
)

type GameState struct {
	Deck   []deck.Card
	State  State
	Player Hand
	Dealer Hand
}

func (gs *GameState) CurrentPlayer() *Hand {
	switch gs.State {
	case StatePlayerTurn:
		return &gs.Player
	case StateDealerTurn:
		return &gs.Dealer
	default:
		panic("It's not currently any player's turn")
	}
}

func clone(gs GameState) GameState {
	ret := GameState{
		Deck:   make([]deck.Card, len(gs.Deck)),
		State:  gs.State,
		Dealer: make(Hand, len(gs.Dealer)),
		Player: make(Hand, len(gs.Player)),
	}
	copy(ret.Deck, gs.Deck)
	copy(ret.Dealer, gs.Dealer)
	copy(ret.Player, gs.Player)

	return ret
}
