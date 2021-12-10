package main

import "fmt"

// new type of slices of string.
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"A", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}
	return cards
}

// (d deck) in front of the function name is called a **Receiver**
// In this case it means that any variable of type 'deck' gets a print method.
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// The following is valid code, Go does not make any mention of this or self though.
// So although valid this is bad convention.
func (this deck) getDeck() deck {
	return this
}
