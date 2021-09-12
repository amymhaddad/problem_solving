package grains

import "errors"

//Square calculates the number of grains on a given square
func Square(num int) (uint64, error) {
	if num <= 0 || num > 64 {
		return 0, errors.New("Invalid number")
	}

	// if num <= 2 {
	// 	return uint64(num), nil
	// }
	//
	// total := 1
	// for i := 2; i <= num; i++ {
	// 	total += i * 2
	// }
	total := 1

	if num == 1 {
		return uint64(total), nil
	}

	for i := 2; i <= num; i++ {
		value := total * 2
		total = value
	}

	return uint64(total), nil
	// if num <= 2 {
	// 	return uint64(num), nil
	// }
	//
}
