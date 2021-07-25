package strand

import "strings"

var dnaRna = map[rune]rune{'G': 'C', 'C': 'G', 'T': 'A', 'A': 'U'}

// ToRNA receives a dna string and returns its rna complement
func ToRNA(dna string) string {

	var rna = make([]string, len(dna))

	for _, letter := range dna {
		rna = append(rna, string(dnaRna[letter]))
	}

	return strings.Join(rna, "")
}

// 	var rna = make([]string, len(dna))

// 	for _, letter := range dna {
// 		rna = append(rna, string(dnaRna[letter]))
// 	}

// 	return strings.Join(rna, "")
// }
