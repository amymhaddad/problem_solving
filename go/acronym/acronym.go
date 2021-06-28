//package main
package acronym

import (
	//"fmt"
	"regexp"
	"strings"
)

// Return the acronym of a given phrase
func Abbreviate(s string) string {
	r := regexp.MustCompile("[A-Za-z]+'?[A-Za-z]*")
	words := r.FindAllString(s, -1)

	var acronym = []string{}
	for _, word := range words{
		letter := string(word[0])
		capitalLetter := strings.Title(strings.ToLower(letter))
		acronym = append(acronym,capitalLetter)
	}
	return strings.Join(acronym, "")
 
}

// func main() {
// 	fmt.Println(Abbreviate("There's Hello world-Here"))
// }
//
