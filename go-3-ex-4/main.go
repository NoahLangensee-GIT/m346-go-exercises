package main

import "fmt"

const (
	Diamonds = '\u25c6' // Karo
	Spades   = '\u2660' // Pik
	Clubs    = '\u2663' // Kreuz
	Hearts   = '\u2665' // Herz

	Six   = '\u2465'
	Seven = '\u2466'
	Eight = '\u2467'
	Nine  = '\u2468'
	Ten   = '\u2469'
	Jack  = 'J'
	Queen = 'Q'
	King  = 'K'
	Ace   = 'A'
)

func main() {
	suits := []rune{Diamonds, Spades, Clubs, Hearts}
	ranks := []rune{Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}
	cards := map[rune][]rune{
		Diamonds: ranks,
		Spades:   ranks,
		Clubs:    ranks,
		Hearts:   ranks,
	}

	for i := 0; i < len(ranks); i++ {
		for j := 0; j < len(suits); j++ {
			fmt.Printf("%c%c\t", suits[j], ranks[i])
		}
		fmt.Println()
	}

}
