package main

func main() {
	cards := deck{"Ace of Spades", "Two of Hearts", "Three of Diamonds"}
	cards = append(cards, createCard())

	cards.printDeck()
}

func createCard() string {
	return "Four of Clubs"
}
