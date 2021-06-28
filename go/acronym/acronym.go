package acronym

import (
	"regexp"
	"strings"
)

// Return the acronym of a given phrase
func Abbreviate(s string) string {
	match := regexp.MustCompile("[A-Za-z]+'?[A-Za-z]*")
	words := match.FindAllString(s, -1)

	var acronym = []string{}

	for _, word := range words {
		firstLetter := string(word[0])
		capitalLetter := strings.Title(strings.ToLower(firstLetter))
		acronym = append(acronym, capitalLetter)
	}
	return strings.Join(acronym, "")
}
