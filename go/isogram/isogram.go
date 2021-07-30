package isogram

import "strings"

// IsIsogram determines if a given string is an isogram
func IsIsogram(phrase string) bool {
	var letters = make(map[rune]bool)

	for _, char := range strings.ToLower(phrase) {

		isAlphabetic := char >= 97 && char <= 122

		if !isAlphabetic {
			continue
		}

		if letters[char] {
			return false
		}

		letters[char] = true

	}
	return true
}
