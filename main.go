package main

func newCard() string {
	return "Five of Diamonds"
}

func main() {
	cards := newDeck()

	// save deck to file locally
	cards.saveDeck("firstDeck.txt")
}
