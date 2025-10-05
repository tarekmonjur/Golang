package main

func main() {
	cards := createNewDeck()
	// cards.printDeck()

	// hand, remainingCards := createDeal(cards, 5)
	// println("Hand:")
	// hand.printDeck()
	// println("Remaining Cards:")
	// remainingCards.printDeck()

	// cards.saveToFile("my_cards.txt")
	// newDeckFromFile("my_cards.txt").printDeck()

	cards.shuffleDeck()
	cards.printDeck()
}
