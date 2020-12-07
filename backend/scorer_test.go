package main

import (
	"fmt"
	"testing"
)

func TestCalculateScoreJackBasra(t *testing.T) {
	cardPlayed := Card{
		Suit:  "C",
		Rank:  "J",
		Value: 0,
	}

	cardsOnTable := []Card{
		{Suit: "D", Rank: "J", Value: 0},
	}

	handScore := CalculateScore(cardPlayed, cardsOnTable)
	fmt.Println(handScore)

	if handScore.Score != 30 {
		t.Errorf("Hand score = %d; want 30", handScore.Score)
	}

	if len(handScore.CardsWon) != 1 {
		t.Errorf("Cards won = %d; want 1", len(handScore.CardsWon))
	}
}

func TestCalculateScoreFaceCardBasra(t *testing.T) {
	cardPlayed1 := Card{
		Suit:  "C",
		Rank:  "K",
		Value: 0,
	}

	cardsOnTable1 := []Card{
		{Suit: "D", Rank: "K", Value: 0},
		{Suit: "H", Rank: "K", Value: 0},
	}

	handScore := CalculateScore(cardPlayed1, cardsOnTable1)

	if handScore.Score != 10 {
		t.Errorf("Hand score = %d; want 10", handScore.Score)
	}

	if len(handScore.CardsWon) != 2 {
		t.Errorf("Cards won = %d; want 2", len(handScore.CardsWon))
	}

	cardPlayed2 := Card{
		Suit:  "C",
		Rank:  "Q",
		Value: 0,
	}

	cardsOnTable2 := []Card{
		{Suit: "D", Rank: "Q", Value: 0},
	}

	handScore2 := CalculateScore(cardPlayed2, cardsOnTable2)

	if handScore2.Score != 10 {
		t.Errorf("Hand score = %d; want 10", handScore2.Score)
	}

	if len(handScore2.CardsWon) != 1 {
		t.Errorf("Cards won = %d; want 1", len(handScore2.CardsWon))
	}
}
