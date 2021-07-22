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
	validTriangle := containsCorrectLengths(a, b, c) && containsTriangleEquality(a, b, c)
}

func containsCorrectLengths(a, b, c float64) bool {
	return a >= 0 && b >= 0 && c >= 0
}

func containsTriangleEquality(a, b, c float64) bool {

	isValid := true

	arr := []float64{3, 4, 5}

	if arr[0]+arr[1] < arr[2] {
		isValid = false
	}

	if arr[0]+arr[2] < arr[1] {
		isValid = false
	}

	if arr[1]+arr[2] < arr[0] {
		isValid = false
	}

	if arr[2]+arr[1] < arr[0] {
		isValid = false
	}
	return isValid

}
