package main

import (
	"strings"
	"fmt"
)

// Return the acronym of a given phrase
func Abbreviate(s string) string {
	//acronym, _ := regexp.MatchString("[A-Z]*", s)

	splitString := strings.Split(s," ")

	var acronym []string
	for _, word:= range splitString {
		firstLetter:= string(word[0])
		acronym = append(acronym, firstLetter)
	}
	return strings.Join(acronym, "") 

}

func main() {
	fmt.Println(Abbreviate("Hello World"))
}

