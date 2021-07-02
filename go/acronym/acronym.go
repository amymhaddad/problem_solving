package acronym

import (
	"regexp"
	"strings"
)

// Return the acronym of a given phrase
func Abbreviate(s string) string {
	acronymMatches := regexp.MustCompile("[A-Z]+'?[A-Z]*")
	words := acronymMatches.FindAllString(strings.ToUpper(s), -1)

	var abbreviations []byte

	for _, word := range words {
		firstLetter := word[0]
		abbreviations = append(abbreviations, firstLetter)
	}
	return string(abbreviations)
}
