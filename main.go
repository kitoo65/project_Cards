package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}

//For Dealing:
//hand, rest := deal(cards, 5)
//fmt.Println(hand.toString())
//fmt.Println(rest.toString())
