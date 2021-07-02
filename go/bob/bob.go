//package bob contains phrases and sentences and Bob's responses to them
package bob

import (
	//	"fmt"
	"regexp"
	//"strings"
)

// Receive a string that contains a phrase or sentence and return a string that
// responds to the given phrase or sentence
func shout(remark string) bool {
	found, _ := regexp.MatchString("[a-z]+", remark)
	return !found
}

func containsNumbers(remark string) bool {
	found , _ := regexp.MatchString("[0-9]+", remark)
	return found
}

func containsLetters(remark string) bool {

	found, _ := regexp.MatchString("[a-z]+", remark)
	return found
}


func askQuestion(remark string) bool {
	//option 1 - use regex 
	found, _ := regexp.MatchString("\\?$", remark)
	return found
	// question := fmt.Sprintf("%s", string(len(remark) -1))  == "?"
	// return question
    //
//	return string(len(remark) -1) == "?"
}

func Hey(remark string) string {
	switch {
	case askQuestion(remark) && shout(remark):
		return "Calm down, I know what I'm doing!"
	case askQuestion(remark) || askQuestion(remark) && containsNumbers(remark):
		return "Sure."
	case shout(remark) || shout(remark) && containsNumbers(remark):
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}

}
