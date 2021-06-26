export const reverseString = (word) => {
	if (!word) return '';

	let wordList = Array.from(word);
	let left = 0;
	let right = word.length - 1;

	while (left < right) {
		[ wordList[left], wordList[right] ] = [ wordList[right], wordList[left] ];
		right -= 1;
		left += 1;
	}
	return wordList.join('');
};
