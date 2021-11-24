package protein

import (
	"errors"
	"fmt"
)

//ErrStop raises an error if a stop codon is encountered
var ErrStop = errors.New("stop")

//ErrInvalidBase raises an error if an invalid base is encountered
var ErrInvalidBase = errors.New("invalid base")

var allCodons = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "",
	"UAG": "",
	"UGA": "",
}

func FromRNA(rna string) ([]string, error) {
	var s []string

	start := 0
	stop := 3
	codonLength := 3
	for start < len(rna) {
		codon := rna[start:stop]
		fmt.Println(start, codon)
		if len(rna)-start < codonLength {
			fmt.Println("here", start, stop)
			codon = rna[start:]
		}

		protein, err := FromCodon(codon)
		if err == nil {
			s = append(s, protein)
		}
		start = stop
		stop += 3
		if start >= len(rna) {
			break
		}
		//	fmt.Println(start, stop, len(rna))

	}
	return s, nil
}

//FromCondon returns the rna protein
func FromCodon(codon string) (string, error) {
	value, found := allCodons[codon]

	if found && value == "" {
		return value, ErrStop
	} else if value != "" {
		return value, nil
	} else {
		return "", ErrInvalidBase
	}
}
