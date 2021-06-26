const rnaToProtein = {
	AUG: 'Methionine',
	UUU: 'Phenylalanine',
	UUC: 'Phenylalanine',
	UUA: 'Leucine',
	UUG: 'Leucine',
	UCU: 'Serine',
	UCC: 'Serine',
	UCA: 'Serine',
	UCG: 'Serine',
	UAU: 'Tyrosine',
	UAC: 'Tyrosine',
	UGU: 'Cysteine',
	UGC: 'Cysteine',
	UGG: 'Tryptophan',
	UAA: 'STOP',
	UAG: 'STOP',
	UGA: 'STOP'
};

export const translate = (string) => {
	if (!string) return [];

	let list = Array.from(string);

	let allProteins = [];
	let i = 0;
	while (i <= list.length - 1) {
		let protein = rnaToProtein[list.slice(i, i + 3).join('')];

		checkProtein(protein, i);
		if (protein === 'STOP') return allProteins;

		allProteins.push(protein);
		i += 3;
	}
	return allProteins;
};

function checkProtein(protein, i) {
	if (protein === undefined) {
		throw new Error('Invalid codon');
	}
	if (protein === 'STOP' && i === 0) {
		return [];
	}
}
