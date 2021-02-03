package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"strings"
)

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest() (string, error) {
	const HandSize = 7
	returnString := "--- Game Beguines --- \n"

	firstDeck := newDeck()
	secondDeck := newDeck()

	returnString += "--- New Player - Paquito ---\n"
	paquito := createPlayer("Paquito", "Navarro")
	paquito.Deck = firstDeck
	paquito.Deck.shuffle()
	paquito.Hand, paquito.Deck = dealHand(paquito.Deck, HandSize)
	returnString += paquito.toJson() + "\n\n"

	returnString += "--- New Player - Gaben ---\n"
	gaben := createPlayer("Gaben", "Newell")
	gaben.Deck = secondDeck
	gaben.Deck.shuffle()
	gaben.Hand, paquito.Deck = dealHand(paquito.Deck, HandSize)
	returnString += gaben.toJson() + "\n\n"

	returnString +="--- AND THE BIG WINNER IS!\n"
	winner := playCards(paquito, gaben)
	returnString += strings.ToUpper(winner.FirstName + "!!!")

	return fmt.Sprintln(returnString), nil
}
