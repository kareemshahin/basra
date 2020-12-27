package main

import (
	"sort"
	"strconv"
)

const KING = "K"
const QUEEN = "Q"
const JACK = "J"

const DIAMONDS = "D"

// TODO: refactor this module, its ugly

func CalculateScore(cardPlayed Card, cardsOnTable []Card) HandScore {
	// if no cards on table, can't win any cards or score
	if len(cardsOnTable) == 0 {
		return HandScore{Score: 0, CardsWon: cardsOnTable}
	}

	// all cards match
	if allEqual(cardPlayed, cardsOnTable) {
		if cardPlayed.Rank == JACK {
			return HandScore{Score: 30, CardsWon: cardsOnTable}
		}
		return HandScore{Score: 10, CardsWon: cardsOnTable}
	}

	// Jack takes all
	if cardPlayed.Rank == JACK {
		return HandScore{Score: 0, CardsWon: cardsOnTable}
	}

	// King or Queen, collect matching faces
	if cardPlayed.Rank == KING || cardPlayed.Rank == QUEEN {
		return HandScore{Score: 0, CardsWon: getEqual(cardPlayed, cardsOnTable)}
	}

	// 7 of diamonds can collect or score basra if total of card values
	// are 10 or under
	if cardPlayed.Value == 7 && cardPlayed.Suit == DIAMONDS {
		if !hasFaceCard(cardsOnTable) {
			possibleInts := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

			for _, v := range possibleInts {
				possibleCardsCollected := getBestCombo(
					Card{Value: v, Suit: DIAMONDS, Rank: strconv.Itoa(v)},
					cardsOnTable,
				)

				if len(possibleCardsCollected) == len(cardsOnTable) {
					return HandScore{Score: 10, CardsWon: cardsOnTable}
				}
			}
		}

		return HandScore{Score: 0, CardsWon: cardsOnTable}
	}

	// If numerical card, find all possible totals that equal card played so we can
	// collect those cards
	if !isFaceCard(cardPlayed) {
		// get all possible combination of totals of the card being played
		bestCardValues := getBestCombo(cardPlayed, cardsOnTable)

		// create lookup indexed by card value
		matchLookup := make(map[int]Card)
		for _, v := range cardsOnTable {
			matchLookup[v.Value] = v
		}

		// get any of the values that match a card to collect
		var cardsToCollect []Card
		for _, v := range bestCardValues {
			if c, ok := matchLookup[v]; ok {
				cardsToCollect = append(cardsToCollect, c)
			}
		}

		if len(cardsToCollect) == len(cardsOnTable) {
			return HandScore{Score: 10, CardsWon: cardsToCollect}
		}

		return HandScore{Score: 0, CardsWon: cardsToCollect}
	}

	return HandScore{Score: 0, CardsWon: []Card{}}
}

/*
	Given an array of possible total arrays (from subsetSum), sorts them and finds the best combo
	of totals that includes the largest set of numbers.

	Example:
	target sum: 7 (car played)
	possible numbers: [1 4 2 3 5 7] (cardsOnTable)
	possible combo of numbers that sum up to 7 [[1 4 2] [4 3] [2 5] [7]]
	result of this function: [4 2 3 5 7]  (4 + 3, 5 + 2, 7)
*/
func getBestCombo(cardPlayed Card, cardsOnTable []Card) []int {
	var possibleTotals [][]int
	var partials []int

	candidateValues := extractPotentialValues(cardPlayed, cardsOnTable)
	subsetSum(candidateValues, cardPlayed.Value, partials, &possibleTotals)

	sort.Slice(possibleTotals, func(i, j int) bool { return len(possibleTotals[j]) < len(possibleTotals[i]) })
	mostTaken := []int{0}

	for i, v := range possibleTotals {

		remainder := arrayDifference(candidateValues, v)
		if (len(candidateValues) - len(remainder)) != len(v) {
			continue
		}

		for j := i + 1; j < len(possibleTotals); j++ {
			newRem := arrayDifference(remainder, possibleTotals[j])
			if (len(remainder) - len(newRem)) != len(possibleTotals[j]) {
				continue
			}
			remainder = newRem
		}
		newTaken := arrayDifference(candidateValues, remainder)

		if len(newTaken) > len(mostTaken) {
			mostTaken = newTaken
		}
	}
	return mostTaken
}

// Get the difference between two arrays
func arrayDifference(a, b []int) (diff []int) {
	m := make(map[int]bool)

	for _, item := range b {
		m[item] = true
	}

	for _, item := range a {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}
	return
}

// Calculates the sum of an array of values
func sumArray(values []int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}

	return sum
}

/*
	Recursively calculate all possible sums and only append to res parameter
	if it equals the target. The result is passed in by reference and matching combinations
	are then appended

	TODO: change to use array of cards rather than integers
*/
func subsetSum(numbers []int, target int, partial []int, res *[][]int) {
	sum := sumArray(partial)

	if sum == target {
		*res = append(*res, partial)
	}

	if sum > target {
		return
	}

	for i, v := range numbers {
		remaining := numbers[i+1:]
		subsetSum(remaining, target, append(partial, v), res)
	}
}

// Check if face card (K, Q, J)
func isFaceCard(card Card) bool {
	return card.Rank == KING || card.Rank == QUEEN || card.Rank == JACK
}

// check if all card ranks are equal to the card played
func allEqual(cardPlayed Card, cards []Card) bool {
	for _, v := range cards {
		if cardPlayed.Rank != v.Rank {
			return false
		}
	}

	return true
}

// Get cards whose rank match the card played
func getEqual(cardPlayed Card, cards []Card) []Card {
	var matchingCards []Card

	for _, v := range cards {
		if cardPlayed.Rank == v.Rank {
			matchingCards = append(matchingCards, v)
		}
	}

	return matchingCards
}

// check if there is a face card in the cards on the table
func hasFaceCard(cards []Card) bool {
	for _, v := range cards {
		if isFaceCard(v) {
			return true
		}
	}

	return false
}

/* Get all values that can be evaluated for sums for the
card that was played.  So face cards and cards who's values
are greater than the card played are eliminated. For example, if an
8 is played, any face cards and anything greater than 8 are eliminated (9, 10)
*/
func extractPotentialValues(cardPlayed Card, cards []Card) []int {
	var values []int

	for _, v := range cards {
		if !isFaceCard(v) && v.Value <= cardPlayed.Value {
			values = append(values, v.Value)
		}
	}

	return values
}
