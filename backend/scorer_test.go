package main

import (
	"fmt"
	"testing"
)

func TestCalculateScore(t *testing.T) {
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
