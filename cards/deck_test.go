package main

import (
	"os"
	"testing"
)

// writing tests in go doesn't require a testing framework
// you simply have to create a file with the same name as one of your go files
// put _test in the file name and then you can begin to write
// go test functions and then run $ go test in the terminal
// to see the results
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card to be Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected loaded deck length of 16, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
