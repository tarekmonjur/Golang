package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of 'deck'
// which is a slice of strings
type deck []string

// A receiver on a function that is associated with the 'deck' type
// which is receiving a deck type variable reference
// so any variable of type 'deck' now gets access to the 'printDeck' function
// (d deck) -> receiver
// d -> A variable called d is a reference to the actual copy of the deck type variable
// which is available in the function
// deck -> every variable of type 'deck' can call this function on itself
func (d deck) printDeck() {
	for i, card := range d {
		fmt.Println("Card", i, ":", card)
	}
}

// Function that creates and returns a new deck of cards
func createNewDeck() deck {
	var cards deck
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace of", "Two of", "Three of", "Four of", "Five of", "Jack of", "Queen of", "King of"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" "+suit)
		}
	}
	return cards
}

// Function that deals a hand of cards
// and returns two values of type 'deck'
// The first return value is the hand of cards
// The second return value is the remaining cards in the deck
func createDeal(cards deck, handSize int) (deck, deck) { // multiple return values
	// Slice Range Syntax
	return cards[:handSize], cards[handSize:]
}

// Converting a deck to a string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Saving a deck to a file
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

// Loading a deck from a file
func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

// Shuffling a deck of cards
func (d deck) shuffleDeck() {
	// for i := range d {
	// 	newPosition := rand.Intn(len(d))
	// 	d[i], d[newPosition] = d[newPosition], d[i]
	// }

	// rand.Shuffle(len(d), func(i, j int) {
	// 	d[i], d[j] = d[j], d[i]
	// })

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d))
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
