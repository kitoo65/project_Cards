package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings

type deck []string

func newDeck() deck { // I put 'deck', will tell go that any time someone calls newDeck, this will return a value of type deck
	//Must create and return a list of playing cards.
	//An array of strings
	cards := deck{}
	cardSuits := []string{"Espada", "Corazon", "Diamante", "Picas"}
	cardValues := []string{"As", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" de "+suit)
		}

	}
	return cards

}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) { // We expect to call this deal func with an existing list of cards which will refer as d deck, we are going to list out the number of cards that we want to have in the hands:
	// We put two times deck, deck in parenthesis cuz we are going to return 2 values.

	return d[:handSize], d[handSize:] //two values.
}
func (d deck) toString() string { //Its going to take a deck and return a complete string representation of it
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {

	//Error: We should just return this error that might get produced when we attempt
	//to write something to the hard drive.
	//So, the func makes that the reciever can be casted with a fileName, and will return an error (if smth goes bad)
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
	//0666 in the perm param, means it is a very default permission. Anyone can read and write this file.
}

func newDeckFromFile(filename string) deck {
	//bs is ByteSlice.
	//err is the error.
	bs, err := ioutil.ReadFile(filename)
	if err != nil { //Here, we handle the error. We have 2 options:
		//Option #1 - Log the error and return a call to newDeck()
		//Option #2 - Maybe, if we are looking a deck, and it is no finding it, there's something extremely wrong with our program.
		//Log the error and entirely quit the program.
		fmt.Println("Error:", err)
		//Also, we are going to use the "os" package, to handle the exit.
		//This package provides me lots of functionalities that work in every OS.
		//It does not matter if it is Linux, or Mac, or windows.
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",") //This is the single string we stored.
	//If we use the strings package again, we can split the single string we have.
	//We have to pass the string, a comma.
	//The variable "s", represents the []string that we are reading from our hard drive.
	return deck(s)
}
func (d deck) shuffle() {
	//For doing the random code:
	//s for source.
	s := rand.NewSource(time.Now().UnixNano())
	//r is for the random number.
	r := rand.New(s)
	for i := range d {
		//newPosition
		np := r.Intn(len(d) - 1)
		//The r, has inside the new Intn func. From the type Rand.
		//Here we swap the positions in each element.
		//This, is going to work, but its going to be strange.
		d[i], d[np] = d[np], d[i]

	}

}
