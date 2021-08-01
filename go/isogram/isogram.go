// Package isogram determines if a string is an isogram
package isogram

import (
	"unicode"
)

// IsIsogram determines if a given string is an isogram
func IsIsogram(phrase string) bool {
	var letters = make(map[rune]struct{})

	for _, char := range phrase {

		char := unicode.ToLower(char)

		if !unicode.IsLower(char) {
			continue
		}

		if _, found := letters[char]; found {
			return false
		}

		letters[char] = struct{}{}

	}
	return true
}
