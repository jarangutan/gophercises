package main

import (
	"fmt"
	"strings"

	"github.com/jarangutan/gophercises/blackjack/deck"
)

type Hand []deck.Card

func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := range h {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ", ")
}

func Score(cards []deck.Card) uint8 {
	var total uint8 = 0
	for _, c := range cards {
		if c.Rank >= deck.Two && c.Rank <= deck.Ten {
			total += uint8(c.Rank)
		} else if c.Rank == deck.Jack || c.Rank == deck.Queen || c.Rank == deck.King {
			total += 10
		} else if c.Rank == deck.Ace {
			if total+11 > 21 {
				total += 1
			} else {
				total += 11
			}
		}
	}

	return total
}

func Round() {
	cards := deck.New(deck.Shuffle)

	var player, dealer Hand

	for i := 0; i < 2; i++ {
		player = append(player, cards[0])
		cards = cards[1:]

		dealer = append(dealer, cards[0])
		cards = cards[1:]
	}

	for {
		fmt.Printf("Dealer: %s, %s\n", dealer[0], "-")
		fmt.Printf("Player: %v\n", player)

		fmt.Printf("(h)it or (s)tand?\n")
		var choice string
		fmt.Scanf("%s\n", &choice)

		if choice == "h" {
			player = append(player, cards[0])
			cards = cards[1:]
		}

		ps := Score(player)

		if ps == 21 {
			fmt.Println("Player blackjack!", ps)
			return
		} else if ps > 21 {
			fmt.Println("Player busted!", ps)
			return
		}

		if choice == "s" {
			break
		}
	}

	for {
		fmt.Printf("Dealer: %v\n", dealer)
		ds := Score(dealer)

		if ds == 21 {
			fmt.Println("Dealer blackjack!", ds)
			return
		} else if ds > 21 {
			fmt.Println("Dealer busted!", ds)
			return
		}

		if ds <= 16 {
			dealer = append(dealer, cards[0])
			cards = cards[1:]
		} else {
			break
		}
	}

	ps := Score(player)
	ds := Score(dealer)
	if ps == ds {
		fmt.Println("Blackjack tie!")

	} else if ps >= ds {
		fmt.Println("Player wins!")
		return
	} else if ds >= ps {
		fmt.Println("Dealer wins!")
		return
	}

}

func main() {
	for {
		Round()
		fmt.Println("---")
		fmt.Println("(n)ew game? (q)uit?")
		var choice string
		fmt.Scanf("%s\n", &choice)
		if choice == "n" {
			continue
		} else {
			break
		}
	}

	fmt.Println("Thanks for playing!")
}
