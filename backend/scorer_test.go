package main

import (
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

func TestCalculateScoreJackTakesAll(t *testing.T) {
	cardPlayed := Card{
		Suit:  "H",
		Rank:  "J",
		Value: 0,
	}

	cardsOnTable := []Card{
		{Suit: "D", Rank: "9", Value: 9},
		{Suit: "H", Rank: "K", Value: 0},
		{Suit: "C", Rank: "2", Value: 2},
	}

	handScore := CalculateScore(cardPlayed, cardsOnTable)

	if handScore.Score != 0 {
		t.Errorf("Hand score = %d; want 0", handScore.Score)
	}

	if len(handScore.CardsWon) != 3 {
		t.Errorf("Cards won = %d; want 3", len(handScore.CardsWon))
	}
}

func TestCalculateScoreFaceCardMatches(t *testing.T) {
	cardPlayedKing := Card{
		Suit:  "H",
		Rank:  "K",
		Value: 0,
	}

	kingToCollect := Card{
		Suit: "H", Rank: "K", Value: 0,
	}

	cardsOnTable := []Card{
		{Suit: "D", Rank: "9", Value: 9},
		{Suit: "C", Rank: "2", Value: 2},
		kingToCollect,
	}

	handScore := CalculateScore(cardPlayedKing, cardsOnTable)

	if handScore.Score != 0 {
		t.Errorf("Hand score = %d; want 0", handScore.Score)
	}

	if len(handScore.CardsWon) != 1 {
		t.Errorf("Cards won = %d; want 1", 1)
	}

	if handScore.CardsWon[0] != kingToCollect {
		t.Errorf("Cards won = %v; want King of Hearts", handScore.CardsWon[0])
	}

	cardPlayedQueen := Card{
		Suit:  "H",
		Rank:  "Q",
		Value: 0,
	}

	queenToCollect := Card{
		Suit: "C", Rank: "Q", Value: 0,
	}

	cardsOnTable2 := []Card{
		{Suit: "D", Rank: "9", Value: 9},
		{Suit: "C", Rank: "K", Value: 0},
		queenToCollect,
	}

	handScore2 := CalculateScore(cardPlayedQueen, cardsOnTable2)

	if handScore2.Score != 0 {
		t.Errorf("Hand score = %d; want 0", handScore2.Score)
	}

	if len(handScore2.CardsWon) != 1 {
		t.Errorf("Cards won = %d; want 1", 1)
	}

	if handScore2.CardsWon[0] != queenToCollect {
		t.Errorf("Cards won = %v; want King of Hearts", handScore2.CardsWon[0])
	}
}

func TestCalculateScoreNumberCardMatchesBasra(t *testing.T) {
	cardPlayed := Card{
		Suit:  "H",
		Rank:  "8",
		Value: 8,
	}

	// all should total 8
	// 6+2, 4+3+1, 8
	cardsOnTable1 := []Card{
		{Suit: "D", Rank: "6", Value: 6},
		{Suit: "C", Rank: "2", Value: 2},
		{Suit: "D", Rank: "8", Value: 8},
		{Suit: "D", Rank: "4", Value: 4},
		{Suit: "S", Rank: "3", Value: 3},
		{Suit: "H", Rank: "1", Value: 1},
	}

	handScore1 := CalculateScore(cardPlayed, cardsOnTable1)

	if handScore1.Score != 10 {
		t.Errorf("Hand score = %d; want 10", handScore1.Score)
	}

	if len(handScore1.CardsWon) != len(cardsOnTable1) {
		t.Errorf("Cards won = %d; want %d", len(handScore1.CardsWon), len(cardsOnTable1))
	}

	// onl card is 8 should be basra
	cardsOnTable2 := []Card{
		{Suit: "D", Rank: "8", Value: 8},
	}

	handScore2 := CalculateScore(cardPlayed, cardsOnTable2)

	if handScore1.Score != 10 {
		t.Errorf("Hand score = %d; want 10", handScore1.Score)
	}

	if len(handScore2.CardsWon) != len(cardsOnTable2) {
		t.Errorf("Cards won = %d; want %d", len(handScore2.CardsWon), len(cardsOnTable2))
	}
}

func TestCalculateScoreNumberCardMatches(t *testing.T) {
	cardPlayed := Card{
		Suit:  "H",
		Rank:  "8",
		Value: 8,
	}

	cardsOnTable1 := []Card{
		{Suit: "D", Rank: "6", Value: 6},
		{Suit: "C", Rank: "2", Value: 2},
		{Suit: "D", Rank: "9", Value: 9},
		{Suit: "D", Rank: "Q", Value: 0},
		{Suit: "S", Rank: "3", Value: 3},
		{Suit: "H", Rank: "1", Value: 1},
	}

	handScore1 := CalculateScore(cardPlayed, cardsOnTable1)

	if handScore1.Score != 0 {
		t.Errorf("Hand score = %d; want 0", handScore1.Score)
	}

	// should only collect 6 + 2 (6+2=8)
	if len(handScore1.CardsWon) != 2 {
		t.Errorf("Cards won = %d; want 2", len(handScore1.CardsWon))
	}
}

func TestCalculateScoreSevenDiamondsCollectAll(t *testing.T) {
	cardPlayed := Card{
		Suit:  "D",
		Rank:  "7",
		Value: 7,
	}

	cardsOnTable1 := []Card{
		{Suit: "D", Rank: "6", Value: 6},
		{Suit: "C", Rank: "2", Value: 2},
		{Suit: "D", Rank: "9", Value: 9},
		{Suit: "D", Rank: "Q", Value: 0},
		{Suit: "H", Rank: "1", Value: 1},
	}

	handScore1 := CalculateScore(cardPlayed, cardsOnTable1)

	if handScore1.Score != 0 {
		t.Errorf("Hand score = %d; want 0", handScore1.Score)
	}

	// 7 diamonds should collect all cards
	if len(handScore1.CardsWon) != len(cardsOnTable1) {
		t.Errorf("Cards won = %d; want %d", len(handScore1.CardsWon), len(cardsOnTable1))
	}

	cardsOnTable2 := []Card{
		{Suit: "D", Rank: "6", Value: 6},
		{Suit: "C", Rank: "2", Value: 2},
		{Suit: "D", Rank: "9", Value: 9},
		{Suit: "H", Rank: "2", Value: 2},
	}

	handScore2 := CalculateScore(cardPlayed, cardsOnTable2)

	if handScore2.Score != 0 {
		t.Errorf("Hand score = %d; want 0", handScore2.Score)
	}

	// 7 diamonds should collect all cards
	if len(handScore2.CardsWon) != len(cardsOnTable2) {
		t.Errorf("Cards won = %d; want %d", len(handScore2.CardsWon), len(cardsOnTable2))
	}
}

func TestCalculateScoreSevenDiamondsBasra(t *testing.T) {
	cardPlayed := Card{
		Suit:  "D",
		Rank:  "7",
		Value: 7,
	}

	// Should all total 8 and get a basra
	// 6 + 2, 3 + 1 + 4, 8
	cardsOnTable := []Card{
		{Suit: "D", Rank: "6", Value: 6},
		{Suit: "C", Rank: "2", Value: 2},
		{Suit: "D", Rank: "8", Value: 8},
		{Suit: "H", Rank: "3", Value: 3},
		{Suit: "H", Rank: "1", Value: 1},
		{Suit: "S", Rank: "4", Value: 4},
	}

	handScore := CalculateScore(cardPlayed, cardsOnTable)

	if handScore.Score != 10 {
		t.Errorf("Hand score = %d; want 0", handScore.Score)
	}

	// 7 diamonds should collect all cards
	if len(handScore.CardsWon) != len(cardsOnTable) {
		t.Errorf("Cards won = %d; want %d", len(handScore.CardsWon), len(cardsOnTable))
	}
}
