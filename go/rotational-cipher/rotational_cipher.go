package rotationalcipher

import (
	"strings"
	"unicode"
)

const alphaLength = 26

//RotationalCipher is a shift cipher that transposes the letters in the alphabet using an integer key between 0 and 26
func RotationalCipher(str string, shiftKey int) string {
	var newVals strings.Builder

	for _, val := range str {
		switch {
		case strings.ContainsRune(string(val), 32):
			newVals.WriteRune(' ')
		case strings.ContainsAny(string(val), "0123456789"):
			newVals.WriteRune(val)
		case unicode.IsUpper(val):
			rotatedVal := rotateLetter(val, shiftKey)
			// distFromStartOfAlpha = int(val) - int('A')
			// shiftVal = (distFromStartOfAlpha + shiftKey) % alphaLength
			// rotatedVal = rune(int('A') + shiftVal)
			newVals.WriteRune(rotatedVal)

		case unicode.IsLower(val):

			rotatedVal := rotateLetter(val, shiftKey)
			// distFromStartOfAlpha = int(val) - int('a')
			// shiftVal = (distFromStartOfAlpha + shiftKey) % alphaLength
			// rotatedVal = rune(int('a') + shiftVal)
			newVals.WriteRune(rotatedVal)

		default:
			newVals.WriteRune(val)
		}
	}
	return newVals.String()
}

//Add bool as aparm on refactor
func rotateLetter(letter rune, shiftKey int) rune {
	var distFromStartOfAlpha int
	var shiftVal int
	var rotatedVal rune

	if unicode.IsUpper(letter) {
		distFromStartOfAlpha = int(letter) - int('A')
		shiftVal = (distFromStartOfAlpha + shiftKey) % alphaLength
		rotatedVal = rune(int('A') + shiftVal)
	} else {
		distFromStartOfAlpha = int(letter) - int('a')
		shiftVal = (distFromStartOfAlpha + shiftKey) % alphaLength
		rotatedVal = rune(int('a') + shiftVal)
	}
	return rotatedVal

}

// func RotationalCipher(str string, shiftKey int) string {
// 	var newString []rune
//
// 	for _, val := range str {
// 		var shiftVal int
// 		var distFromStartOfAlpha int
// 		var rotatedVal rune
//
// 		switch {
// 		case strings.ContainsRune(string(val), 32):
// 			newString = append(newString, ' ')
// 		case strings.ContainsAny(string(val), "0123456789"):
// 			newString = append(newString, val)
// 		case int(val) >= 65 && int(val) <= 90:
// 			distFromStartOfAlpha = int(val) - int('A')
// 			shiftVal = (distFromStartOfAlpha + shiftKey) % alphaLength
// 			rotatedVal = rune(int('A') + shiftVal)
// 			newString = append(newString, rotatedVal)
// 		case int(val) >= 97 && int(val) <= 123:
// 			distFromStartOfAlpha = int(val) - int('a')
// 			shiftVal = (distFromStartOfAlpha + shiftKey) % alphaLength
// 			rotatedVal = rune(int('a') + shiftVal)
// 			newString = append(newString, rotatedVal)
// 		default:
// 			newString = append(newString, val)
// 		}
// 	}
// 	return string(newString)
// }
