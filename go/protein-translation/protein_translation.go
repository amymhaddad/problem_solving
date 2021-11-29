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
}
var StopCodons = [3]string{"UAA", "UAG", "UGA"}

//FromRNA returns a slice of protein strings
func FromRNA(rna string) ([]string, error) {
	var proteins []string

	for i := 0; i <= len(rna)-3; i += 3 {
		codon := rna[i : i+3]

		if isStopCodon(codon) {
			return proteins, nil
		}

		protein, err := FromCodon(codon)

		if err != nil {
			return proteins, err
		}
		proteins = append(proteins, protein)
	}
	return proteins, nil
}

func isStopCodon(codon string) bool {
	for _, stopCodon := range StopCodons {
		if codon == stopCodon {
			return true
		}
	}
	return false
}

//FromCodon returns the rna protein
func FromCodon(codon string) (string, error) {
	value := allCodons[codon]

	if isStopCodon(codon) {
		return "", ErrStop
	}
	if value == "" {
		return "", ErrInvalidBase
	}
	return value, nil
}
