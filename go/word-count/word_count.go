package wordcount

import (
	"regexp"
	"strings"
)

//Frequency keeps track of the times a word appears in a given string
type Frequency map[string]int

//WordCount counts the number of times a word appears in a given string
func WordCount(s string) Frequency {
	foundWords := make(Frequency)
	s = strings.ToLower(s)
	regex := regexp.MustCompile("[a-z0-9]+('[a-z])?")
	found := regex.FindAllString(s, -1)

	for _, word := range found {
		_, found := foundWords[word]
		if !found {
			foundWords[word] = 1
		} else {
			foundWords[word]++
		}
	}
	return foundWords
}
