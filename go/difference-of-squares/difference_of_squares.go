//Package diffsquares finds the difference between the square of the sum and the
//sum of the squares from n to 1
package diffsquares

//SquareOfSum returns the square of the sum from 1 to the given number
func SquareOfSum(num int) int {
	sum := num * (num + 1) / 2
	return sum * sum
}

//SumOfSquares returns the sum of the squares from 1 to the given number
func SumOfSquares(num int) int {
	return num * (num + 1) * (2*num + 1) / 6
}

//Difference returns the difference between the square of the sum and the sum
//of the squares for a given number
func Difference(num int) int {
	return SquareOfSum(num) - SumOfSquares(num)
}
