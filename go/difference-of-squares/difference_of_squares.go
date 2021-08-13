//Package diffsquares finds the difference between the square of the sum and the
//sum of the squares from n to 1
package diffsquares

//SquareOfSum sums all numbers from the given number to one, and squares the
//total
func SquareOfSum(num int) int {
	sum := num * (num + 1) / 2
	return sum * sum
}

//SumOfSquares squares each number from the given number to one, and adds them
//together
func SumOfSquares(num int) int {

	squaredNum := num * num
	sum := squaredNum * (squaredNum + 1) / 2
	return sum
	// var total int

	// for num > 0 {
	// 	total += num * num
	// 	num--
	// }
	// return total

}

//Difference calculates the difference between the square of the sum and the sum
//of the squares for a given number
func Difference(num int) int {

	return SquareOfSum(num) - SumOfSquares(num)

}
