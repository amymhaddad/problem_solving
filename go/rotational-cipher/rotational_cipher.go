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
		var letterSize int

		if !unicode.IsLetter(r) {
			result.WriteRune(r)
			continue
		}

		if unicode.IsUpper(r) {
			letterSize = int('A')
		} else {
			letterSize = int('a')
		}
		rotatedVal := rotateLetter(r, shiftKey, letterSize)
		result.WriteRune(rotatedVal)

	}
	return result.String()
}

func rotateLetter(letter rune, shiftKey int, letterSize int) rune {
	var distFromStartOfAlpha, shiftVal int

	distFromStartOfAlpha = int(letter) - letterSize
	shiftVal = (distFromStartOfAlpha + shiftKey) % alphaLength
	return rune(letterSize + shiftVal)
}
