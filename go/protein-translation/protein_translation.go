package protein

import "errors"

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
	panic("Please implement the FromRNA function")
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
