package main

import (
	"fmt"
	"os"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "King", "Queen"}

	for _, s := range cardSuits {
		for _, v := range cardValues {
			cards = append(cards, v+" of "+s)
		}
	}
	return cards
}

// use a receiver (d deck) to add a function as a method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize uint) (deck, deck) {
	return d[0:handSize], d[handSize:]
}

func (d deck) saveDeck(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
}

func (d deck) toString() string {
	sDeck := []string(d)

	return strings.Join(sDeck, ",")
}
