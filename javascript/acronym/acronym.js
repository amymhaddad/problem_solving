export const parse = (phrase) => {
  let matches = phrase
    .match(/[A-Za-z\-a-z]|[A-Z,]|[A-Za-z\s]+/g)
    .join('')
    .split(/[\- ,]+/);

	let acronym = matches.map((word) => word[0].toUpperCase());
	return acronym.join('');
};
