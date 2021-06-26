export class Matrix {
	constructor(string) {
		this.rows = this.processRows(string.split('\n'));
		this.columns = this.processColumns(this.rows);
	}

	processRows(numbers) {
		return numbers.map((num) => {
			let row = num.split(' ');

			let rowAsInt = row.map((int) => {
				return parseInt(int);
			});
			return rowAsInt;
		});
	}

	processColumns(rows) {
		const eachColumn = [];

		let i = 0;
		while (i <= rows.length) {
			let numberAtIndex = rows.map((row) => {
				return row[i];
			});
			eachColumn.push(numberAtIndex);
			i += 1;
		}
		return eachColumn;
	}
}
