//Package luhn determines if a number is valid based on the Luhn formula
package luhn

import (
	"fmt"
	"strings"
	"unicode"
)

//Valid determines if a given number is valid based on the Luhn formula
func Valid(nums string) bool {
	nums = strings.ReplaceAll(nums, " ", "")
	double := false
	var total int

	//reg for loop bc going right ot left
	for _, val := range nums {
		if !unicode.IsNumber(val) || len(nums) <= 1 {
			return false
		}

		digitValue := int(val - '0')
		fmt.Println(digitValue)
		if double && digitValue != 9 {
			digitValue = (digitValue * 2) % 9
			double = false
		} else {
			double = true
		}

		total += digitValue
	}
	return total%10 == 0

}

// func Valid(nums string) bool {
// 	nums = strings.ReplaceAll(nums, " ", "")
// 	altDigit := len(nums) - 2
// 	var total int

// 	for i, val := range nums {
// 		if !unicode.IsNumber(val) || len(nums) <= 1 {
// 			return false
// 		}
// 		currIndex := len(nums) - 1 - i

// 		currRuneVal := nums[currIndex]
// 		fmt.Println("here: ", currRuneVal)
// 		digitValue := int(val - '0')

// 		if currIndex == altDigit && digitValue != 9 {
// 			digitValue = (digitValue * 2) % 9
// 			altDigit -= 2
// 		}
// 		total += digitValue

// 	}
// 	return total%10 == 0
// }
