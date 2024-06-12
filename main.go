package main

func main() {
	// cards := newDeck()
	d := getDeckFromFile("firstDeck.txt")
	d.print()
	d.shuffle()
	d.print()
}
