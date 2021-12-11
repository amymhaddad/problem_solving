package rotationalcipher

const alphaLength = 26

func RotationalCipher(str string, shiftKey int) string {
	var rotatedIndex int

	for _, ch := range str {

		rotationAmt := shiftKey

		distanceFromA := ch - 'a'
		newVal := 'a' + distanceFromA

		if newVal > 'z' {
			rotationAmt = alphaLength % shiftKey
		}
		rotatedIndex = 'a' + rotationAmt
	}
	return string(rune(rotatedIndex))
}
