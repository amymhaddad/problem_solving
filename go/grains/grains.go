package grains

import (
	"errors"
)

//Square calculates the number of grains on a given square
func Square(num int) (uint64, error) {
	if num <= 0 || num > 64 {
		return 0, errors.New("Invalid square")
	}

	if num == 1 {
		return uint64(num), nil
	}

	total := 1

	for i := 2; i <= num; i++ {
		total = total << 1
	}

	return uint64(total), nil
}

//Total calculates the total number of grains on the chessboard
func Total() uint64 {

	var totalGrains uint64

	for squareNum := 1; squareNum <= 64; squareNum++ {
		grainsOnSquare, _ := Square(squareNum)
		totalGrains += grainsOnSquare
	}
	return totalGrains
}
