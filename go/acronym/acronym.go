package main

import (
	//"strings"
	"fmt"
	"regexp"
)

// Return the acronym of a given phrase
func Abbreviate(s string) string {
	r := regexp.MustCompile("[a-zA-Z]*")
	words := r.FindAllString(s, -1)
	fmt.Println(words)
//	standardizeInput := strings.Title(strings.ToLower(strings.Join(words, "")))
	return "made it"
//	fmt.Println(standardizeInput)	
	 
	// var acronym []string
	// for _, word:= range splitString {
	// 	firstLetter:= string(word[0])
	// 	acronym = append(acronym, firstLetter)
	// }
	// return strings.Join(acronym, "") 
    //
}

func main() {
	fmt.Println(Abbreviate("Hello World"))
}

