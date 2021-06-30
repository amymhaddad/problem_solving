
//package bob contains phrases and sentences and Bob's responses to them
package bob

import (
	"regexp"
	"strings"
)

// Receive a string that contains a phrase or sentence and return a string that
// responds to the given phrase or sentence

func shouting(remark string) bool {
	match := regexp.MustCompile("[A-Z]+")
	words := match.FindAllString(remark, -1)

	matchedWords := strings.Join(words, " ")
	return len(remark) == len(matchedWords)
}

func Hey(remark string) string {
		

	return "Whatever."

}
