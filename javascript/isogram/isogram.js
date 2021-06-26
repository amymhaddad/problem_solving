export const isIsogram = (word) => {
	const updatedWord = normalizeWord(word);

	const unique = new Set();
	for (let letter of updatedWord) {
		unique.add(letter);
	}
	return unique.size === updatedWord.length;
};

function normalizeWord(word) {
	const normalizeChars = /[^\s\-]/;
	let normalizedWord = '';
	Array.from(word).forEach((char) => {
		let lowerCaseChar = char.toLowerCase().trim();
		if (normalizeChars.test(lowerCaseChar)) {
			normalizedWord += lowerCaseChar;
		}
	});
	return normalizedWord;
}
