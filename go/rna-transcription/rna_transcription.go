package strand

var dnaRNA = map[rune]rune{'G': 'C', 'C': 'G', 'T': 'A', 'A': 'U'}

// ToRNA receives a dna string and returns its rna complement
func ToRNA(dna string) string {

	var rna = make([]rune, len(dna))

	for i, letter := range dna {
		rna[i] = dnaRNA[letter]
	}

	return string(rna)

}
