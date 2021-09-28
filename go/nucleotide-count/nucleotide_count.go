package dna

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram struct {
	G int
	A int
	C int
	T int
}

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
///
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {
	//	nucleotides := map[string]bool{"G": true, "A": true, "C": true, "T": true}
	nucleotides := map[rune]string{65: "A", 67: "C", 71: "G", 84: "T"}

	var h Histogram

	for _, letter := range d {

	}

	return h, nil
}
