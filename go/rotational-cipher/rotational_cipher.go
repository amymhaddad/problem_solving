package rotationalcipher

import (
	"strings"
)

const alphaLength = 26

//RotationalCipher is a shift cipher that transposes the letters in the alphabet using an integer key between 0 and 26
func RotationalCipher(str string, shiftKey int) string {
	var newString []rune

	for _, val := range str {
		var shiftVal int
		var distFromStartOfAlpha int
		var rotatedVal rune

		switch {
		case strings.ContainsRune(string(val), 32):
			newString = append(newString, ' ')
		case strings.ContainsAny(string(val), "0123456789"):
			newString = append(newString, val)
		case int(val) >= 65 && int(val) <= 90:
			distFromStartOfAlpha = int(val) - int('A')
			shiftVal = (distFromStartOfAlpha + shiftKey) % alphaLength
			rotatedVal = rune(int('A') + shiftVal)
			newString = append(newString, rotatedVal)
		case int(val) >= 97 && int(val) <= 123:
			distFromStartOfAlpha = int(val) - int('a')
			shiftVal = (distFromStartOfAlpha + shiftKey) % alphaLength
			rotatedVal = rune(int('a') + shiftVal)
			newString = append(newString, rotatedVal)
		default:
			newString = append(newString, val)
		}
	}
	return string(newString)
}
