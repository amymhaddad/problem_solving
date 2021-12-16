package rotationalcipher

import (
	"fmt"
	"strings"
)

const alphaLength = 26

func RotationalCipher(str string, shiftKey int) string {
	var newString []rune

	for _, val := range str {
		var shiftVal int
		var distFromA int
		var newVal rune

		switch {
		case strings.ContainsRune(string(val), 32):
			newString = append(newString, ' ')
			continue
		case strings.ContainsAny(string(val), "0123456789"):
			newString = append(newString, val)
			continue
		case int(val) >= 65 && int(val) <= 90:
			distFromA = int(val) - int('A')
			shiftVal = (distFromA + shiftKey) % alphaLength
			newVal = rune(int('A') + shiftVal)
		case int(val) >= 97 && int(val) <= 123:
			distFromA = int(val) - int('a')
			shiftVal = (distFromA + shiftKey) % alphaLength
			newVal = rune(int('a') + shiftVal)
		default:
			newString = append(newString, val)
		}
		newString = append(newString, newVal)
		// if strings.ContainsRune(string(val), 32) {
		// 	newString = append(newString, ' ')
		// 	continue
		// }
		//
		// if strings.ContainsAny(string(val), "0123456789") {
		// 	newString = append(newString, val)
		// 	continue
		//
		// }
		//
		// if int(val) >= 65 && int(val) <= 90 {
		// 	distFromA = int(val) - int('A')
		// 	shiftVal = (distFromA + shiftKey) % alphaLength
		// 	newVal = rune(int('A') + shiftVal)
		// } else {
		// 	distFromA = int(val) - int('a')
		// 	shiftVal = (distFromA + shiftKey) % alphaLength
		// 	newVal = rune(int('a') + shiftVal)
		// }
		//	newString = append(newString, newVal)

	}
	fmt.Println("H: ", string(newString))
	return string(newString)
}
