export class HighScores {
	constructor(all_scores) {
		this.all_scores = all_scores;
	}

	get scores() {
		return this.all_scores;
	}

	get latest() {
		return parseInt([ this.all_scores[this.all_scores.length - 1] ]);
	}

	get personalBest() {
		return parseInt(Math.max(...this.all_scores));
	}

	get personalTopThree() {
		let topScores = this.all_scores.sort((a, b) => b - a);
		return topScores.slice(0, 3);
	}
}
