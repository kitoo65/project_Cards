package main

import "fmt"

func main() {

	cards := newDeck()
	fmt.Println(cards.toString())

}

//For Dealing:
//hand, rest := deal(cards, 5)
//fmt.Println(hand.toString())
//fmt.Println(rest.toString())
