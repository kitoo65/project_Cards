package main

import (
	"fmt"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of strings

type deck []string

func newDeck() deck { // I put 'deck', will tell go that any time someone calls newDeck, this will return a value of type deck
	//Must create and return a list of playing cards.
	//An array of strings
	cards := deck{}
	cardSuits := []string{"Espada", "Corazon", "Diamante", "Picas"}
	cardValues := []string{"As de ", "2 de ", "3 de ", "4 de ", "5 de ", "6 de ", "7 de ", "8 de", "9 de ", "10 de ", "J de ", "Q de ", "K de "}
	var addedCard string

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			addedCard = value + suit
			cards = append(cards, addedCard)
		}

	}
	return cards

}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) { // We expect to call this deal func with an existing list of cards which will refer as d deck, we are going to list out the number of cards that we want to have in the hands:
	// We put two times deck, deck in parenthesis cuz we are going to return 2 values.

	return d[:handSize], d[handSize:] //two values.
}
func (d deck) toString() string { //Its going to take a deck and return a complete string representation of it
	return strings.Join([]string(d), ",")
}
