package protein

import (
	"errors"
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

//FromRNA returns a slice of protein strings
func FromRNA(rna string) ([]string, error) {
	var proteins []string
	seen := make(map[string]bool)

	for i := 0; i <= len(rna)-3; i += 3 {
		codon, err := FromCodon(rna[i : i+3])

		switch err {
		case ErrInvalidBase:
			return proteins, err
		case ErrStop:
			return proteins, nil
		}

		if seen[codon] || codon == "" {
			continue
		}

		proteins = append(proteins, codon)
		seen[codon] = true
	}
	return proteins, nil
}

//FromCodon returns the rna protein
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
