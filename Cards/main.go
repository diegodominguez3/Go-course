package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "six of spades")

	for i, card := range cards {
		fmt.Println(card)
	}
}

func newCard() string {
	return "Five of Dimonds"
}
