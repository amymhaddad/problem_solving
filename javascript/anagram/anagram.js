export const findAnagrams = (givenWord, words) => {
	const parsedWords = parseAllWords(givenWord, words);

	let anagrams = parsedWords.filter((word) => {
		const sortedGivenWord = Array.from(givenWord.toLowerCase()).sort().join('');
		let sortedWord = Array.from(word.toLowerCase()).sort().join('');
		if (sortedWord === sortedGivenWord) {
			return word;
		}
	});
	return anagrams;
};

function parseAllWords(givenWord, words) {
	return words.join(' ').split(/[\s+]/).filter((word) => word.toLowerCase() != givenWord.toLowerCase());
}
