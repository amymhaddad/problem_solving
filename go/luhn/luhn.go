//Package luhn determines if a number is valid based on the Luhn formula
package luhn

import (
	"strings"
)

//Valid determines if a given number is valid based on the Luhn formula
func Valid(nums string) bool {
	nums = strings.ReplaceAll(nums, " ", "")
	double := len(nums)%2 == 0
	var total int

	if len(nums) <= 1 {
		return false
	}
	for _, num := range nums {
		validNum := num-'0' >= 0 && num-'0' <= 9
		if !validNum {
			return false
		}

		digit := int(num - '0')

		if double && digit != 9 {
			digit = (digit * 2) % 9
		}

		double = !double
		total += digit
	}

	return total%10 == 0

}
