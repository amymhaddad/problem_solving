package grains

import (
	"errors"
)

const (
	minSquare int = 0
	maxSquare int = 64
)

//Square calculates the number of grains on a given square
func Square(num int) (uint64, error) {
	if num <= minSquare || num > maxSquare {
		return 0, errors.New("Invalid square")
	}

	return 1 << (num - 1), nil
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
