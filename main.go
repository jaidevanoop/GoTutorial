package main

import "fmt"

func main() {

	cards := newDeck()

	// hand, remaningDeck := deal(cards, 5)
	//
	// hand.print()
	// remaningDeck.print()

	fmt.Println(cards.toString())
}
