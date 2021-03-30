package main

import "github.com/google/uuid"

type Game struct {
	ID      uuid.UUID    `json:"id"`
	Name    string       `json:"name"`
	State   string       `json:"state"`
	Players []GamePlayer `json:"players"`
}

type GamePlayer struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Score int       `json:"score"`
}

type Player struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Score          int       `json:"score"`
	Hand           []Card    `json:"hand"`
	CardsCollected []Card    `json:"cards_collectd"`
}

type PlayerToken struct {
	ID    uuid.UUID
	Token string
}

type Card struct {
	Suit  string `json:"suit"`
	Rank  string `json:"Rank"`
	Value int    `json:"value"`
}

type HandScore struct {
	Score    int
	CardsWon []Card
}
