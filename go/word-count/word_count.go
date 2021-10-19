package wordcount

import (
	"regexp"
	"strings"
)

//Frequency keeps track of the times a word appears in a given string
type Frequency map[string]int

//WordCount counts the number of times a word appears in a given string
func WordCount(s string) Frequency {
	//always need to init a map which is what make() does.
	f := make(Frequency)
	s = strings.ToLower(s)

	re := regexp.MustCompile("[A-Za-z]*['a-z]?[0-9]*")
	words := re.FindAllString(s, -1)

	for _, word := range words {
		count, found := f[word]

		if found {
			count++
		} else {
			f[word] = 1
		}
	}
	return f
}
