const range = require('lodash');

const sounds = {
    3: 'Pling',
    5: 'Plang',
    7: 'Plong'
};
export const convert = (number) => {
    let range_arr = range.range(1, number + 1);

    let factors = range_arr.filter((num) => number % num === 0 && num in sounds);
    let words = factors.map((num) => sounds[num]);

    if (words.length === 0) return number.toString();
    return words.join('');
};