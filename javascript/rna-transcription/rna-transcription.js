const dnaToRna = {
	G: 'C',
	C: 'G',
	T: 'A',
	A: 'U'
};

export const toRna = (dna) => {
	return Array.from(dna).map((n) => dnaToRna[n]).join('');
};
