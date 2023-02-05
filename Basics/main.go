package main

import "fmt"

var deckSize int = 20

func main() {
	fmt.Println("Hello World")

	var card string = "Spades"
	secondCard := "Hearts"
	fmt.Println(card, " ", secondCard, " ", deckSize)

	//slices
	cards := []string{"Hearts", "Spades", "Clubs"}
	fmt.Println(cards)

	trees := make([]string, 0)
	trees = append(trees, "mango")
	trees = append(trees, "coconut", "jak-fruit")
	fmt.Println(trees)

	for i, tree := range trees {
		fmt.Println(i, " = ", tree)
	}

	//creating a deck of cards
	deckOfCards := NewDeck()
	deckOfCards.printCardsInDeck()

	deck1, deck2 := deckOfCards.deal(3)
	deck1.printCardsInDeck()
	deck2.printCardsInDeck()

	deckOfCards.saveToFile("my_cards")

	loadedDeck := loadDeckFromFile("my_cards")
	loadedDeck.printCardsInDeck()

	loadedDeck.shuffleDeck()
	loadedDeck.printCardsInDeck()

}
