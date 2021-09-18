package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	// create map
	m := make(map[string]int)
	// Put cards into it
	m["ace"] = 11
	m["two"] = 2
	m["three"] = 3
	m["four"] = 4
	m["five"] = 5
	m["six"] = 6
	m["seven"] = 7
	m["eight"] = 8
	m["nine"] = 9
	m["ten"] = 10
	m["jack"] = 10
	m["queen"] = 10
	m["king"] = 10
	m["other"] = 0

	return m[card]
}

// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
	if ParseCard(card1)+ParseCard(card2) == 21 {
		return true
	} else {
		return false
	}
}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
	if isBlackjack == true && dealerScore <= 9 {
		return "W"
	} else if isBlackjack == true && dealerScore > 9 {
		return "S"
	} else {
		return "P"
	}
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {
	if handScore >= 17 {
		return "S"
	} else if handScore <= 11 {
		return "H"
	} else if handScore >= 12 && handScore <= 16 && dealerScore >= 7 {
		return "H"
	} else {
		return "S"
	}
}

// FirstTurn returns the semi-optimal decision for the first turn, given the cards of the player and the dealer.
// This function is already implemented and does not need to be edited. It pulls the other functions together in a
// complete decision tree for the first turn.
func FirstTurn(card1, card2, dealerCard string) string {
	handScore := ParseCard(card1) + ParseCard(card2)
	dealerScore := ParseCard(dealerCard)

	if 20 < handScore {
		return LargeHand(IsBlackjack(card1, card2), dealerScore)
	}
	return SmallHand(handScore, dealerScore)
}
