package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type deck struct {
	Cards []card
}

type card struct {
	Value int
	Suit  string
}

func newDeck() deck {
	deck := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			card := card{
				Suit:  suit,
				Value: value,
			}
			deck.Cards = append(deck.Cards, card)
		}
	}

	return deck
}

func (d deck) print() {
	for i, card := range d.Cards {
		fmt.Println(i+1, card)
	}
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d.Cards {
		newPosition := r.Intn(len(d.Cards) - 1)

		d.Cards[i], d.Cards[newPosition] = d.Cards[newPosition], d.Cards[i]
	}
}

func (d deck) toJson() string {
	b, _ := json.Marshal(d)
	return string(b)
}

func dealHand(d deck, handSize int) (deck, deck) {
	handCards := d.Cards[:handSize]
	remainingCards := d.Cards[handSize:]

	return deck{Cards: handCards}, deck{Cards: remainingCards}
}
