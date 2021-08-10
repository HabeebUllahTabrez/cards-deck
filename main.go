package main

func main() {
	// cards := newDeck()

	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()

	// cards.saveToFile("My_Cards")

	cards := newDeckFromFile("My_Cards")
	cards.shuffle()
	cards.print()
}
