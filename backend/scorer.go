package main

const KING = "K"
const QUEEN = "Q"
const JACK = "J"

// TODO: test the heck out of this logic

func calculateBasra(cardPlayed Card, cardsOnTable []Card) HandScore {

	// all cards match
	if allEqual(cardPlayed, cardsOnTable) {
		if cardPlayed.Rank == JACK {
			return HandScore{Score: 30, CardsWon: cardsOnTable}
		}
		return HandScore{Score: 10, CardsWon: cardsOnTable}
	}

	/*
		TODO: Use subsetSum to figure out the combos that sum up to the card played
		and determined if its a basra or not and what cards they win. Might have to filter out
		any face cards or cards whose value is greater than the card played.
		It's possible that this new logic can eliminate the check above

		if card played isFaceCard
			only pick up face cards
			basra should be covered in case above
		else
			filter out face cards and any cards greater than card played value
			check for totals for cards won using subsetSum and check for basra (no cards left)
		end
	*/

	return HandScore{Score: 0, CardsWon: []Card{}}
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

func isFaceCard(cardRank string) bool {
	return cardRank == KING || cardRank == QUEEN || cardRank == JACK
}

func allEqual(cardPlayed Card, cards []Card) bool {
	for _, v := range cards {
		if cardPlayed.Rank != v.Rank {
			return false
		}
	}

	return true
}
