package main

import (
	"os"
	"testing"
)

// `t` is a test handler
// We tell `t` that something just went wrong
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

// For testing file I/O, you need to make sure that
// whenever we are writing test with Go
// we take care of any cleanup that needs to be done manually
// because there is no hand-holding by the Go testing framework
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// Need to handle error properly, but ignore this time
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
