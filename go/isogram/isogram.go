// Package isogram determines if a string is an isogram
package isogram

import (
	"unicode"
)

// IsIsogram determines if a given string is an isogram
func IsIsogram(phrase string) bool {
	var uniqueLetters = make(map[rune]struct{})

	for _, char := range phrase {

		if !unicode.IsLetter(char) {
			continue
		}

		char := unicode.ToLower(char)

		if _, found := uniqueLetters[char]; found {
			return false
		}

		uniqueLetters[char] = struct{}{}

	}
	return true
}
