package triangle

import (
	"math"
)

// Kind refers to the kind of triangle that's returned by KindFromSides()
type Kind string

// Declare the constants to return kind of triangle
const (
	NaT = "NaT"
	Equ = "Equ"
	Iso = "Iso"
	Sca = "Sca"
)

// KindFromSides returns the kind of triangle
func KindFromSides(a, b, c float64) Kind {

	switch {
	case !ValidTriangle(a, b, c):
		return NaT
	case a == b && b == c:
		return Equ
	case a == b || b == c || a == c:
		return Iso
	default:
		return Sca
	}

}

// ValidTriangle determines if a triangle is valid
func ValidTriangle(a, b, c float64) bool {

	if math.IsNaN(a+b+c) || math.IsInf(a+b+c, 0) {
		return false
	}

	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}

	if a+b < c || b+c < a || c+a < b {
		return false
	}

	return true
}
