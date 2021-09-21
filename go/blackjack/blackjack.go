package blackjack

const blackjack = 21

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ace":
		return 11
	default:
		return 10
	}
}

// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
	return ParseCard(card1)+ParseCard(card2) == blackjack
}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
	containsAce := dealerScore%11 == 0
	containsFigureOrTen := dealerScore%10 == 0

	switch {
	case !isBlackjack && containsAce:
		return "P"
	case (isBlackjack) && (!containsAce && !containsFigureOrTen):
		return "W"
	default:
		return "S"
	}
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {
	switch {
	case (handScore >= 17) || (handScore >= 12 && handScore <= 16 && dealerScore < 7):
		return "S"
	default:
		return "H"
	}
}
