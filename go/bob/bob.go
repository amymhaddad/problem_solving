//package bob contains phrases and sentences and Bob's responses to them
package bob

import (
	"regexp"
	"strings"
)

// Receive a string that contains a phrase or sentence and return a string that
// responds to the given phrase or sentence

func shout(remark string) bool {
	match := regexp.MustCompile("[A-Z]+'?[A-Z]*")
	words := match.FindAllString(remark, -1)
	splitRemark := strings.Fields(remark)


	return len(words) == len(splitRemark)
}

func askQuestion(remark string) bool {
	return strings.ContainsAny(remark, "?")
}

func Hey(remark string) string {
	switch {
	case askQuestion(remark) && shout(remark):
		return "Calm down, I know what I'm doing!"
	case shout(remark):
		return "Whoa, chill out!"
	case askQuestion(remark):
		return "Sure."
	default:
		return "Whatever."
	}

}
