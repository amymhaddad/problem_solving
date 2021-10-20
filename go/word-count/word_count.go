package wordcount

import (
	"fmt"
	"regexp"
	"strings"
)

//Frequency keeps track of the times a word appears in a given string
type Frequency map[string]int

//WordCount counts the number of times a word appears in a given string
func WordCount(s string) Frequency {
	//always need to init a map which is what make() does.
	f := make(Frequency)
	//f := Frequency{} -- less common - use when i have map vals I want to init w/
	s = strings.ToLower(s)

	//match an apostrophe and any letter w/in the range [a-z] as a group
	//Use group -- not just []. If put apostrophe inside brackets ['a-z]? -- this matches any of these chars
	//vs what I have now is a specifc order
	re := regexp.MustCompile("[0-9A-Za-z]+('[a-z])?")

	//Works in Python, not Go. unclear how to capture numbers and letters
	//re := regexp.MustCompile("\d*\w['a-z]?")
	words := re.FindAllString(s, -1)
	fmt.Println("words: ", words, len(words))

	for _, word := range words {
		count, found := f[word]

		if found {
			//	fmt.Println("found", word, count)
			f[word] = count + 1
		} else {
			f[word] = 1
		}
	}
	return f
}
