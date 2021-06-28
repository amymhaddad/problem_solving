package main
//package acronym

import (
	"fmt"
	"regexp"
	"strings"
)

// Return the acronym of a given phrase
func Abbreviate(s string) string {
	r := regexp.MustCompile("[a-zA-Z]+([a-z'a-z])")
	words := r.FindAllString(s, -1)
	fmt.Println(r, words)
	standardizeInput := strings.Title(strings.ToLower(strings.Join(words, " ")))

	capitalLetters := regexp.MustCompile("[A-Z]+")
	acronym := capitalLetters.FindAllString(standardizeInput, -1)
	return strings.Join(acronym, "")
 
}

func main() {
	fmt.Println(Abbreviate("There's Hello world-Here"))
}

