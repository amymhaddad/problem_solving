//package bob contains phrases and sentences and Bob's responses to them
package bob

import (
	"regexp"
	"strings"
)

// Receive a string that contains a phrase or sentence and return a string that
// responds to the given phrase or sentence
func shout(remark string) bool {
	containLowercaseLetters, _ := regexp.MatchString("[a-z]+", remark)
	containUppercaseLetters, _ := regexp.MatchString("[A-Z]+", remark)

	return !containLowercaseLetters && containUppercaseLetters
}

func containsNumbers(remark string) bool {
	found, _ := regexp.MatchString("[0-9]+", remark)

	return found
}

func silence(remark string) bool {
	return strings.TrimSpace(remark) == ""
}

func askQuestion(remark string) bool {
	found, _ := regexp.MatchString("\\?$", strings.TrimSpace(remark))

	return found
}

func Hey(remark string) string {
	switch {
	case askQuestion(remark) && shout(remark):
		return "Calm down, I know what I'm doing!"
	case askQuestion(remark) || askQuestion(remark) && containsNumbers(remark):
		return "Sure."
	case shout(remark) || shout(remark) && containsNumbers(remark):
		return "Whoa, chill out!"
	case silence(remark):
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}
