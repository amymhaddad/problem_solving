package dna

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
func (d DNA) Counts() (Histogram, error) {
	hist := Histogram{'G': 0, 'T': 0, 'A': 0, 'C': 0}

	for _, n := range d {
		if _, found := hist[n]; found {
			hist[n]++
		} else {
			return hist, errors.New("invalid strand")
		}
	}

	return hist, nil
}
