package triangle

import "fmt"

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

	if !validTriangle {

		return "NaT"
	}

	allSides := []float64{a, b, c}
	matches := map[float64]bool{}

	for _, side := range allSides {
		fmt.Printf("\n%T SIDE: ", side)
		matches[side] = true
	}

	uniqueSides := len(matches)
	//fmt.Println("\nUNIQUE: ", uniqueSides)
	switch uniqueSides {
	case 3:
		return Sca
	case 2:
		return Iso
	default:
		return Equ

	}

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
