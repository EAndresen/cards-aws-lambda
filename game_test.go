package main

import (
	"strconv"
	"testing"
)

func TestSummeryCardsInHand(t *testing.T) {
	deck := newDeck()

	if summeryCardsInHand(deck) != 220 {
		t.Errorf("Expected summery on card values to be 220, but got: %v", summeryCardsInHand(deck))
	}
}

func TestPlayCards(t *testing.T) {
	testPlayer1 := createTestPlayer(1)
	testPlayer2 := createTestPlayer(2)

	winner := playCards(testPlayer1, testPlayer2)

	if winner.FirstName != "Name2" {
		t.Errorf("Expected winner to be Player two, but got: %v", winner.FirstName)
	}

}

func createTestPlayer(value int) player {
	player := createPlayer("Name"+strconv.Itoa(value), "Last")
	playerHnd := deck{}
	playerCard := card{
		Value: value,
		Suit:  "",
	}
	playerHnd.Cards = append(playerHnd.Cards, playerCard)
	player.Hand = playerHnd
	return player
}
