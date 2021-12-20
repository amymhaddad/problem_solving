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
		var rotate int

		if !unicode.IsLetter(r) {
			result.WriteRune(r)
			continue
		}

		if unicode.IsUpper(r) {
			rotate = ((int(r)-int('A'))+shiftKey)%alphaLength + 'A'
		} else {
			rotate = ((int(r)-int('a'))+shiftKey)%alphaLength + 'a'
		}

		result.WriteRune(rune(rotate))
	}
	return result.String()
}
