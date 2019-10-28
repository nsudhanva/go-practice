package main

import "fmt"

func main() {
	// var card string = "Ace of spades"
	// card := "Ace of spades"

	// Not a new variable
	// card := "Five of Diamonds"

	// card = "Five of Diamonds"
	
	card := newCard()
	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}