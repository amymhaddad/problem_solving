package strand

var dnaRna = map[string]string{"G": "C", "C": "G", "T": "A", "A": "U"}

// ToRNA recieves a dna string and returns its rna complement
func ToRNA(dna string) string {

	var rna string

	for _, letter := range dna {
		rna += dnaRna[string(letter)]
	}
	return rna
}
