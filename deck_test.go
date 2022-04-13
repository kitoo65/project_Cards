package main

import "testing"

//t *testing.T is the go testing runner.
//When anything goes wrong, we tell the "t" value that smth just went wrong

func TestNewDeck(t *testing.T) {
	d := newDeck()
	l := 52
	//The number 52, is the quantity of cards we may have created.
	if len(d) != l {
		//Errorf, is the error print in the terminal.
		t.Errorf("Expected deck length of %v, but got %d", l, len(d))
	}
}
