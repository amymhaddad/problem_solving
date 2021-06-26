const range = require('lodash');

export class Squares {
    constructor(number) {
        this.number_range = range.range(1, number + 1);
    }

    get sumOfSquares() {
        return this.number_range.map((num) => num ** 2).reduce((acc, curr) => (acc += curr));
    }

    get squareOfSum() {
        return this.number_range.reduce((acc, curr) => (acc += curr)) ** 2;
    }

    get difference() {
        return this.squareOfSum - this.sumOfSquares;
    }
}
