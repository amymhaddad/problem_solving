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
	return "NaT"
}

func containsCorrectLengths(a, b, c float64) bool {

}
