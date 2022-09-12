package main

import (
	"fmt"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected Four of Clubs got %v", d[len(d)-1])
	}
}

func TestNewDeckSaveToFileAndLoadFromFile(t *testing.T) {
	d := newDeck()
	d.saveToFile("_testing")
	loadedDeck := readFile("_testing")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards but got %d", len(loadedDeck))
	}
	cleanUp()
}

func cleanUp() {
	err := os.Remove("_testing")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
