package triangle

import (
	"math"
)

// Kind indicates the kind of triangle
type Kind string

const (
	NaT = "NaT"
	Equ = "Equ"
	Iso = "Iso"
	Sca = "Sca"
)

// KindFromSides returns the kind of triangle
func KindFromSides(a, b, c float64) Kind {
	validTriangle := containsCorrectLengths(a, b, c) && containsTriangleEquality(a, b, c) && containsValidNumbers(a, b, c)

	if !validTriangle {
		return "NaT"
	}

	triangleSides := []float64{a, b, c}
	allSides := map[float64]bool{}

	for _, side := range triangleSides {
		allSides[side] = true
	}

	uniqueSides := len(allSides)
	switch uniqueSides {
	case 3:
		return Sca
	case 2:
		return Iso
	default:
		return Equ
	}
}

func containsValidNumbers(a, b, c float64) bool {
	nums := []float64{a, b, c}

	for i := range nums {
		if math.IsInf(nums[i], 0) || math.IsNaN(nums[i]) {
			return false
		}
	}

	return true
}

func containsCorrectLengths(a, b, c float64) bool {
	return a > 0 && b > 0 && c > 0
}

func containsTriangleEquality(a, b, c float64) bool {
	isValid := true

	sides := []float64{a, b, c}

	if sides[0]+sides[1] < sides[2] {
		isValid = false
	} else if sides[0]+sides[2] < sides[1] {
		isValid = false
	} else if sides[1]+sides[2] < sides[0] {
		isValid = false
	} else if sides[2]+sides[1] < sides[0] {
		isValid = false
	}

	return isValid

}
