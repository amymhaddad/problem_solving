package wordcount

import (
	"regexp"
	"strings"
)

//Frequency keeps track of the times a word appears in a given string
type Frequency map[string]int

//WordCount counts the number of times a word appears in a given string
func WordCount(s string) Frequency {
	wordCounter := make(Frequency)
	re := regexp.MustCompile("[0-9a-z]+('[a-z])?")
	words := re.FindAllString(strings.ToLower(s), -1)

	for _, word := range words {
		wordCounter[word]++
	}
	return wordCounter
}
