package main

import (
	"fmt"
	"sort"
)

const KING = "K"
const QUEEN = "Q"
const JACK = "J"

const DIAMONDS = "D"

// TODO: test the heck out of this logic

func CalculateScore(cardPlayed Card, cardsOnTable []Card) HandScore {

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

	// TODO: handle 7 of diamonds
	if cardPlayed.Value == 7 && cardPlayed.Suit == DIAMONDS {
		fmt.Println("TODO: handle 7 of diamonds")

		return HandScore{Score: 0, CardsWon: cardsOnTable}
	}

	/*
		TODO: Use subsetSum to figure out the combos that sum up to the card played
		and determined if its a basra or not and what cards they win. Might have to filter out
		any face cards or cards whose value is greater than the card played.
		It's possible that this new logic can eliminate the check above

		if card played isFaceCard
			only pick up face cards
			basra should be covered in case above
			jack should pick everything up!
		else
			filter out face cards and any cards greater than card played value
			check for totals for cards won using subsetSum and check for basra (no cards left)
			check for 7 diamonds
		end
	*/

	// If numerical card, find all possible totals that equal card played so we can
	// collect those cards
	if !isFaceCard(cardPlayed) {
		// get all possible combination of totals of the card being played
		var possibleTotals [][]int
		var partials []int
		candidateValues := extractPotentialValues(cardPlayed, cardsOnTable)
		subsetSum(candidateValues, cardPlayed.Value, partials, &possibleTotals)

		bestCardValues := getBestCombo(possibleTotals, candidateValues)
		fmt.Println(bestCardValues)

		// create lookup indexed by card value
		var matchLookup map[int]Card
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
	target sum: 7
	possible numbers: [1 4 2 3 5 7]
	possible combo of numbers that sum up to 7 [[1 4 2] [4 3] [2 5] [7]]
	result of this function: [4 2 3 5 7]  (4 + 3, 5 + 2, 7)
*/
func getBestCombo(totals [][]int, original []int) []int {
	sort.Slice(totals, func(i, j int) bool { return len(totals[j]) < len(totals[i]) })
	mostTaken := []int{0}

	for i, v := range totals {

		remainder := Difference(original, v)
		if (len(original) - len(remainder)) != len(v) {
			continue
		}
		fmt.Println("remaining", remainder)
		for j := i + 1; j < len(totals); j++ {
			newRem := Difference(remainder, totals[j])
			if (len(remainder) - len(newRem)) != len(totals[j]) {
				continue
			}
			remainder = newRem
		}
		newTaken := Difference(original, remainder)

		if len(newTaken) > len(mostTaken) {
			mostTaken = newTaken
		}
	}
	return mostTaken
}

// Set Difference: A - B
func Difference(a, b []int) (diff []int) {
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

func isFaceCard(card Card) bool {
	return card.Rank == KING || card.Rank == QUEEN || card.Rank == JACK
}

func allEqual(cardPlayed Card, cards []Card) bool {
	for _, v := range cards {
		if cardPlayed.Rank != v.Rank {
			return false
		}
	}

	return true
}

func getEqual(cardPlayed Card, cards []Card) []Card {
	var matchingCards []Card

	for _, v := range cards {
		if cardPlayed.Rank == v.Rank {
			matchingCards = append(matchingCards, v)
		}
	}

	return matchingCards
}

func extractPotentialValues(cardPlayed Card, cards []Card) []int {
	var values []int

	for _, v := range cards {
		if !isFaceCard(cardPlayed) && v.Value >= cardPlayed.Value {
			values = append(values, v.Value)
		}
	}

	return values
}
