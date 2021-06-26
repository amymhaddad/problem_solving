
const fp = require('lodash/fp')

export const classify = (number) => {
	if (number < 1) {
		throw new Error('Classification is only possible for natural numbers.')
	}

	const sum_of_factors = sumOfFactors(number);

	return (
		isPerfect(sum_of_factors, number) || isAbundant(sum_of_factors, number) || isDeficient(sum_of_factors, number)
	)

};

function sumOfFactors(number) {
	return fp.sum(fp.range(1, number).filter(num => number % num === 0))
}


function isPerfect(sum_of_factors, number) {
	if (sum_of_factors === number) return "perfect";
}

function isAbundant(sum_of_factors, number) {
	if (sum_of_factors > number) return "abundant";
}

function isDeficient(sum_of_factors, number) {
	if (sum_of_factors < number) return "deficient";
}
