package main

import "fmt"

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
