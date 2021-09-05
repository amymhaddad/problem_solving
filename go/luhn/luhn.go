//Package luhn determines if a number is valid based on the Luhn formula
package luhn

import (
	"strings"
)

//Valid determines if a given number is valid based on the Luhn formula
func Valid(nums string) bool {
	nums = strings.ReplaceAll(nums, " ", "")
	double := false
	var total int

	for i := len(nums) - 1; i >= 0; i-- {
		validNum := nums[i] >= 48 && nums[i] <= 57

		if !validNum || len(nums) <= 1 {
			return false
		}

		digit := int(nums[i] - 48)

		if double && digit != 9 {
			digit = (digit * 2) % 9
			double = false
		} else {
			double = true
		}

		total += digit
	}
	return total%10 == 0

}
