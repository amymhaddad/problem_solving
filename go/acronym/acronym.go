//package main
package acronym

import (
	//"fmt"
	"regexp"
	"strings"
)

// Return the acronym of a given phrase
func Abbreviate(s string) string {
	r := regexp.MustCompile("[^A-Za-z]*")
	//r := regexp.MustCompile("[a-zA-Z]*")
	words := r.FindAllString(s, -1)
	 
	standardizeInput := strings.Title(strings.ToLower(strings.Join(words, " ")))

	capitalLetters := regexp.MustCompile("[A-Z]+")
	acronym := capitalLetters.FindAllString(standardizeInput, -1)
	return strings.Join(acronym, "")
 
}

// func main() {
// 	fmt.Println(Abbreviate("Hello world-Here"))
// }

