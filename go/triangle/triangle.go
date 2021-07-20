package triangle

// Kind indicates the kind of triangle
type Kind string

const (
	//NaT contains sides that don't form a triangle
	NaT = "NaT" // not a triangle
	//Equ indicates an equilateral triangle (three sides of the same lenth)
	Equ = "Equ"
	//Iso indicates an isosceles triangle (two sides of the same length)
	Iso = "Iso"
	//Sca indicates a scalene triangle (three different side lengths)
	Sca = "Sca"
)

// KindFromSides returns the kind of triangle
func KindFromSides(a, b, c float64) Kind {
	validTriangle := twoSidesLargerThanThird(a, b, c)
}

func containsCorrectLengths(a, b, c float64) bool {
	return a >= 0 && b >= 0 && c >= 0
}

func twoSidesLargerThanThird(a, b, c float64) bool {
	//add length to this slice?
	sides := []float64{a, b, c}

	count := 0

	for count < 3 {
		s1 := 0
		s2 := 1
		s3 := 2

		if sides[s1]+sides[s2] < sides[s3] || sides[s1]+sides[s3] < sides[s2] {
			return false
		}

		firstNum := sides[0]
		sides = sides[1:]
		sides = append(sides, firstNum)
	}
	return true

}
