package rotationalcipher

import (
	"strings"
)

const alphaLength = 26

func RotationalCipher(str string, shiftKey int) string {
	var newString []rune

	for _, val := range str {
		var shiftVal int
		var distFromA int
		var newVal rune

		//isAlpha := int(val) >= 97 && int(val) <= 123 || val >= 65 && int(val) <= 90

		if strings.ContainsRune(string(val), 32) {
			//	emptyString := ' '
			newString = append(newString, ' ')
		}

		if int(val) >= 65 && int(val) <= 90 {
			distFromA = int(val) - int('A')
			shiftVal = (distFromA + shiftKey) % alphaLength
			newVal = rune(int('A') + shiftVal)
		} else {
			distFromA = int(val) - int('a')
			shiftVal = (distFromA + shiftKey) % alphaLength
			newVal = rune(int('a') + shiftVal)
		}
		newString = append(newString, newVal)

	}
	return string(newString)
}
