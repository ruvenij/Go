package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

type Deck []string

func NewDeck() Deck {
	cardSuites := []string{"hearts", "diamond", "clubs", "spades"}
	cardValue := []string{"ace", "two", "three", "four"}
	cards := Deck{}

	for _, suite := range cardSuites {
		for _, value := range cardValue {
			cards = append(cards, suite+" of "+value)
		}
	}

	return cards
}

func (d Deck) printCardsInDeck() {
	fmt.Println("*** Printing the cards in the deck ***")
	for i, card := range d {
		fmt.Println(i, " ", card)
	}
}

func (d Deck) deal(numberOfCards int) (Deck, Deck) {
	return d[:numberOfCards], d[numberOfCards:]
}

func (d Deck) saveToFile(fileName string) {
	ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func (d Deck) toString() string {
	return strings.Join([]string(d), ",")
}

func loadDeckFromFile(fileName string) Deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error : ", err.Error())
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return Deck([]string(s))
}

func (d Deck) shuffleDeck() {
	for i := range d {
		randN := rand.Intn(len(d))
		tempString := d[i]
		d[i] = d[randN]
		d[randN] = tempString
	}
}
