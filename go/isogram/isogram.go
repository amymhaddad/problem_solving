package isogram

import "strings"

var letters = make(map[rune]bool)

// IsIsogram determines if a given string is an isogram
func IsIsogram(phrase string) bool {

	for _, char := range strings.ToLower(phrase) {

		isAlphabetic := char >= 97 && char <= 122
		_, isNotUnique := letters[char]

		if isAlphabetic && isNotUnique {
			return false
		}
		letters[char] = true

		// if _, ok := letters[char]; ok {
		// 	return false
		// }
		// letters[char] = true
	}
	return true
}
