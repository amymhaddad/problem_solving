package bob

import (
	"regexp"
	"strings"
	"unicode"
)

//Solution 1:
const (
	responseToSilence         = "Fine. Be that way!"
	responseToYellingQuestion = "Calm down, I know what I'm doing!"
	responseToQuestion        = "Sure."
	responseToShouting        = "Whoa, chill out!"
	defaultResponse           = "Whatever."
)

//Hey returns Bob's response
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	switch {
	case silence(remark):
		return responseToSilence
	case containsLetters(remark) && yellingQuestion(remark):
		return responseToYellingQuestion
	case question(remark):
		return responseToQuestion
	case containsLetters(remark) && shout(remark):
		return responseToShouting
	default:
		return defaultResponse
	}
}

func silence(remark string) bool {
	return len(remark) == 0
}

func yellingQuestion(remark string) bool {
	return strings.HasSuffix(remark, "?") && remark == strings.ToUpper(remark)
}

func question(remark string) bool {
	return strings.HasSuffix(remark, "?")
}

func shout(remark string) bool {
	return remark == strings.ToUpper(remark)
}

func containsLetters(remark string) bool {
	return strings.IndexFunc(remark, unicode.IsLetter) >= 0
}

//Solution 2:

//Hey returns Bob's response
func Hey(remark string) string {
	switch {
	case askQuestion(remark) && shout(remark):
		return "Calm down, I know what I'm doing!"
	case askQuestion(remark) || askQuestion(remark) && containsNumbers(remark):
		return "Sure."
	case shout(remark) || shout(remark) && containsNumbers(remark):
		return "Whoa, chill out!"
	case silence(remark):
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}
func shout(remark string) bool {
	containLowercaseLetters, _ := regexp.MatchString("[a-z]+", remark)
	containUppercaseLetters, _ := regexp.MatchString("[A-Z]+", remark)

	return !containLowercaseLetters && containUppercaseLetters
}

func containsNumbers(remark string) bool {
	found, _ := regexp.MatchString("[0-9]+", remark)

	return found
}

func silence(remark string) bool {
	return strings.TrimSpace(remark) == ""
}

func askQuestion(remark string) bool {
	found, _ := regexp.MatchString("\\?$", strings.TrimSpace(remark))

	return found
}
