package rotationalcipher

import (
	"strings"
	"unicode"
)

const alphaLength = 26

//RotationalCipher is a shift cipher that transposes the letters in the alphabet using an integer key between 0 and 26
func RotationalCipher(str string, shiftKey int) string {
	var result strings.Builder

	for _, val := range str {
		var letterSize int
		isLower := unicode.IsLower(val)
		isUpper := unicode.IsUpper(val)
		isNumber := strings.ContainsAny(string(val), "0123456789")
		isSpace := strings.ContainsRune(string(val), 32)

		switch {
		case isSpace:
			result.WriteRune(' ')
		case isNumber || (!isUpper && !isLower):
			result.WriteRune(val)
		default:
			if isUpper {
				letterSize = int('A')
			} else {
				letterSize = int('a')
			}
			rotatedVal := rotateLetter(val, shiftKey, letterSize)
			result.WriteRune(rotatedVal)
		}
	}
	return result.String()
}

func rotateLetter(letter rune, shiftKey int, letterSize int) rune {
	var distFromStartOfAlpha, shiftVal int

	distFromStartOfAlpha = int(letter) - letterSize
	shiftVal = (distFromStartOfAlpha + shiftKey) % alphaLength
	return rune(letterSize + shiftVal)
}
