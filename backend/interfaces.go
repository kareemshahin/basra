package main

type Game struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	State string `json:"state"`
}

type Player struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Token string
	Score int `json:"score"`
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
