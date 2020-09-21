package main

func main() {
	//var card string = "Ace of Spades"
	cards := newDeck()
	cards.print()
	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()
}

func newCard() string {
	return "Five of Dimonds"
}
