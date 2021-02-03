package main

import "encoding/json"

type player struct {
	FirstName string
	LastName  string
	Rank      int
	Hand      deck
	Deck      deck
}

func createPlayer(firstName string, lastName string) player {
	return player{
		FirstName: firstName,
		LastName:  lastName,
	}
}

func (p player) toJson() string{
	b, _ := json.Marshal(p)
	return string(b)
}
