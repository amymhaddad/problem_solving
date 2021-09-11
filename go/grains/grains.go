package grains

import "errors"

func Square(num int) (uint64, error) {
	if num <= 0 || num > 64 {
		return 0, errors.New("Invalid number")
	}

	if num <= 2 {
		return uint64(num), nil
	}

}
