package diffsquares

import "math"

//SquareOfSum sums all numbers from 1 through the given number, and squares the
//total
func SquareOfSum(num int) int {

	var total int

	for num > 0 {
		total += num
		num--
	}

	sq := math.Sqrt(float64(total))
	return int(sq)
}
