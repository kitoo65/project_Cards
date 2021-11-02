package main

func main() {

	cards := newDeck()
	cards.print()

	hand, remainingCards := deal(cards, 5)
	hand.print()

	remainingCards.print()

}

func newCard() string {
	return "Five of Diamonds"
}

func (d deck) toString() string { //Its going to take a deck and return a complete string representation of it
	//hi

}
