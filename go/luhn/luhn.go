//Package luhn determines if a number is valid based on the Luhn formula
package luhn

import (
	"strings"
	"unicode"
)

//Valid determines if a given number is valid based on the Luhn formula
func Valid(nums string) bool {
	nums = strings.ReplaceAll(nums, " ", "")
	altDigit := len(nums) - 2
	var total int

	for i, val := range nums {
		if !unicode.IsNumber(val) || len(nums) <= 1 {
			return false
		}

		currIndex := len(nums) - 1 - i
		currRuneVal := nums[currIndex]
		digitValue := int(currRuneVal - '0')

		if currIndex == altDigit && digitValue != 9 {
			digitValue = (digitValue * 2) % 9
			altDigit -= 2
		}
		total += digitValue
	}

	return total%10 == 0

}
