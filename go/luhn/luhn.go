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

		digitValue := int(val - '0')
		if i == altDigit {
			digitValue = (digitValue * 2) % 9
			altDigit -= 2
		}
		total += digitValue
	}

	return total%10 == 0

}
