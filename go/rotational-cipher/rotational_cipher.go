package rotationalcipher

const alphaLength = 26

func RotationalCipher(str string, shiftKey int) string {

	distanceFromA := str - 'a'
	newVal := 'a' + distanceFromA

	if newVal > 'z' {
		rotationAmt := alphaLength % shiftKey
	} else {
		rotationAmt := shiftKey
	}

	rotatedIndex := 'a' + rotationAmt
	rotatedVal := string(rotatedIndex)
	return rotatedVal

}
