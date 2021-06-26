export const isPangram = (sentence) => {
	if (!sentence) return false;

	const letter_counter = create_letter_counter();
	const letters = [ ...sentence.match(/[a-z]/gi) ];

	for (let letter of letters) {
		letter_counter[letter.toLowerCase()] += 1;
	}

	if (Object.values(letter_counter).includes(0)) {
		return false;
	}
	return true;
};

function create_letter_counter() {
	let index = 97;
	let letter_counter = {};

	for (let i = 0; i < 26; i++) {
		let eachLetter = String.fromCharCode(index);
		letter_counter[eachLetter] = 0;
		index += 1;
	}
	return letter_counter;
}
