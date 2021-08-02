package diffsquares

//import "math"

//SquareOfSum sums all numbers from 1 through the given number, and squares the
//total
func SquareOfSum(num int) int {

	var total int

	for num > 0 {
		total += num
		num--
	}

	return total * total
}
