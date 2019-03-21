package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// '_', blank identifier, means that there is a variable but we are not gonna use it
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// (d deck) -> receiver
// d: The actual copy of the deck we're working with is available
//    in the function as variable called 'd'.
// deck: Every variable of type 'deck' can call this function on itself.
// ----------------------------------------------------------------------------
// Any variable of type "deck" now gets access to the "print" method.
// The receiver sets up methods on variables that we create.
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
