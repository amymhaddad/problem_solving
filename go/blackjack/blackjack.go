package blackjack

const blackjack = 21

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	cardValues := map[string]int{
		"ace":   11,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"ten":   10,
		"jack":  10,
		"queen": 10,
		"king":  10,
	}

	return cardValues[card]

}

// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
	return ParseCard(card1)+ParseCard(card2) == blackjack
}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
	containsAce := dealerScore%11 == 0
	containsFigureOrTen := dealerScore%10 == 0

	if !isBlackjack {
		if containsAce {
			return "P"
		}
		return "S"
	}

	if (isBlackjack) && (!containsAce && !containsFigureOrTen) {
		return "W"
	}

	return "S"
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {

	if handScore >= 17 {
		return "S"
	}

	if (handScore >= 12 && handScore <= 16) && dealerScore < 7 {
		return "S"
	}

	return "H"
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
