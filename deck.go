package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	// perm: 0666 means anyone can read and write this file
	// * perm -> permission
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	// bs, err := ioutil.ReadFile(filename)
	// bs: byte slice
	// error: Value of type 'error'.
	//        If nothing went wrong, it will have a value of 'nil'
	bs, err := ioutil.ReadFile(filename)

	// Error handling
	// For handling error,
	// ask yourself hey if something goes wrong here what do I really want to happen
	// Example:
	// Option #1 - Log the error and return a call to newDeck()
	// Option #2 - Log the error and entirely quit the program
	if err != nil {
		// Take option #2
		fmt.Println("Error:", err)
		// code zero indicates success, non-zero an error
		os.Exit(1)
	}

	// string(bs): Type conversion from 'byte[]' to 'string'
	// => Ace of Spades,Two of Spades,Three of Spades, ...
	// Split the string into a slice of string
	s := strings.Split(string(bs), ",")

	// Convert the slice into deck then return
	return deck(s)
}

func (d deck) shuffle() {
	// To be truly random, it's required to change a seed used for generating a random number
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		// rand.Intn(len(d) - 1): Generate a random number between '0' and 'the length of the deck - 1'
		newPosition := r.Intn(len(d) - 1)

		// Swap the position
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
