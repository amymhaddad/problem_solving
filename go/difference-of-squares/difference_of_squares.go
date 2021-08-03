package diffsquares

import "math"

//SquareOfSum sums all numbers from 1 through the given number, and squares the
//total
func SquareOfSum(num int) int {

	var total float64

	for num > 0 {
		total += float64(num)
		num--
	}

	return int(math.Pow(total, 2))
}

//SumOfSquares squares all numbers from 1 to the given number, and adds them
//together
func SumOfSquares(num int) int {
	var total float64

	for num > 0 {
		total += math.Pow(float64(num), 2)
		num--
	}
	return int(total)

}
