package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) <= 0 {
		t.Errorf("Expected deck to have some cards, bot got %v", len(d))
	}

	if d[0] != "Ace of Diamonds" {
		t.Errorf("Expected the first card to be Ace of Diamonds, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected the last card to be Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) <= 0 {
		t.Errorf("Expected deck to have some cards, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
