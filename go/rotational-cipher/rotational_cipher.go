package rotationalcipher

import (
	"strings"
	"unicode"
)

const alphaLength = 26

//RotationalCipher is a shift cipher that transposes the letters in the alphabet using an integer key between 0 and 26
func RotationalCipher(str string, shiftKey int) string {
	var result strings.Builder

	for _, r := range str {

		if !unicode.IsLetter(r) {
			result.WriteRune(r)
			continue
		}

		base := 'a'
		if unicode.IsUpper(r) {
			base = 'A'
		}

		rotated_value := ((r-base)+rune(shiftKey))%alphaLength + base
		result.WriteRune(rotated_value)
	}
	return result.String()
}
