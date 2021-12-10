package main

import (
	"fmt"
	"strconv"
)

func mainold() {
	// var card string = "Ace of Spades"
	card := "Ace of Spades" // only when defining a new variables.
	// card = "Five of Diamonds" // reassign a new value
	fmt.Print(card)

	card = newCard()
	fmt.Println(card)

	// Arrays and Slices...
	// Array -> Fixed length.
	// Slices -> variable length.
	// cards := []string{newCard(), newCard()}
	// cards := deck{newCard(), newCard()}

	cards := newDeck()
	fmt.Println("Create a new Deck")
	fmt.Println(cards)

	// append does not modify the exisitng cards array but creates and returns a new one.
	cards = append(cards, "Six of spades")

	// Iterator over the slice.
	// for & range -> take a slice and loop over it.
	// Also note the := to keep declaring and assigning a new variable.
	fmt.Println("Range for loop Begin")
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	// The above commented code was moved to cards.go
	cards.print()
	fmt.Println("Range for loop End")

	hand, remainingCards := deal(cards, 5)
	fmt.Println("HAND")
	hand.print()
	fmt.Println("REMAINING CARDS")
	remainingCards.print()

}

func newCard() string {
	return "Five of diamonds"
}


func errorVariables(args[] interface{}) {
	var total, nInts, nFloats int
	invalid := make([]string, 0)
	for _,k := range args[1:] {
		_, err := strconv.Atoi(k)
		if err == nil {
			total++
			nInts++
			continue
		}

		_, err = strconv.ParseFloat(k, )
		if err == nil {
			total++
			nFloats++
			continue
		}
	}

	fmt.Println(total)
	fmt.Println(nInts)
	fmt.Println(invalid)

}

func main()  {
	errorVariables(1, "2", 3)
}