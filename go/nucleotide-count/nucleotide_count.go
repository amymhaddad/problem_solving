package dna

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.

//Can declar empty map but can't init it
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
///
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {
	hist := Histogram{'G': 0, 'T': 0, 'A':0, 'C':0}


	var h Histogram

	for _, letter := range d {

	}

	return h, nil
}
