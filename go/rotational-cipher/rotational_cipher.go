package rotationalcipher

import "fmt"

const alphaLength = 26

func RotationalCipher(str string, shiftKey int) string {
	var rotatedIndex int

	for _, ch := range str {

		rotationAmt := shiftKey

		distanceFromA := ch - 'a'
		newVal := 'a' + distanceFromA

		fmt.Println("Here: ", ch+distanceFromA, newVal, distanceFromA)
		switch {
		case newVal > 'z':
			rotationAmt = alphaLength % shiftKey
		case newVal == 'a' && rotationAmt == alphaLength:
			rotationAmt = 0
		}
		rotatedIndex = 'a' + rotationAmt
	}
	return string(rune(rotatedIndex))
}
