package main

func playCards(player1 player, player2 player) player {
	var winner player
	if summeryCardsInHand(player1.Hand) > summeryCardsInHand(player2.Hand) {
		player1.Rank = 1
		player2.Rank = 2
		winner = player1
	} else if summeryCardsInHand(player1.Hand) < summeryCardsInHand(player2.Hand) {
		player1.Rank = 2
		player2.Rank = 1
		winner = player2
	} else {
		winner.FirstName = "DRAW"
	}
	return winner
}

func summeryCardsInHand(Hand deck) int {
	var summery int
	for _, card := range Hand.Cards {
		summery += card.Value
	}

	return summery
}
