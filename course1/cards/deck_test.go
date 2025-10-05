package main

import (
	"os"
	"testing"
)

func TestCreateNewDeck(t *testing.T) {
	deck := createNewDeck()
	if len(deck) != 32 {
		t.Errorf("Expected deck length of 32, but got %d", len(deck))
	}

	if deck[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be 'Ace of Spades', but got %s", deck[0])
	}

	if deck[len(deck)-1] != "King of Clubs" {
		t.Errorf("Expected last card to be 'King of Clubs', but got %s", deck[len(deck)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("test_deck.txt") // Clean up before test

	deck := createNewDeck()
	err := deck.saveToFile("test_deck.txt")
	if err != nil {
		t.Errorf("Error saving deck to file: %v", err)
	}

	loadedDeck := newDeckFromFile("test_deck.txt")
	if len(loadedDeck) != len(deck) {
		t.Errorf("Expected loaded deck length of %d, but got %d", len(deck), len(loadedDeck))
	}

	os.Remove("test_deck.txt") // Clean up after test
}
