//Package luhn determines if a number is valid based on the Luhn formula
package luhn

import (
	"strings"
)

//Valid determines if a given number is valid based on the Luhn formula
func Valid(num string) bool {

	if !isValidString(num) {
		return false
	}

	validNumbers := getValidNums(num)

	return isLuhn(validNumbers)

}

func isValidString(num string) bool {
	num = strings.ReplaceAll(num, " ", "")
	letterChar := num[len(num)-1] >= 97 && num[len(num)-1] <= 122

	if len(num) <= 1 || letterChar {
		return false
	}

	for i := 0; i < len(num); i++ {
		validNums := num[i] >= '0' && num[i] <= '9'

		if !validNums {
			return false
		}
	}

	return true

}

func getValidNums(num string) []int {
	var validNumbers []int

	for i := 0; i < len(num); i++ {
		validNums := num[i] >= '0' && num[i] <= '9'
		if validNums {
			validNumbers = append(validNumbers, int(num[i]-'0'))
		}
	}

	return validNumbers

}

func isLuhn(validNumbers []int) bool {
	luhnNums := make([]int, len(validNumbers))
	var totalSum int
	indexToDouble := len(validNumbers) - 2

	for i := len(validNumbers) - 1; i >= 0; i-- {
		if i == indexToDouble && indexToDouble >= 0 {
			doubleNum := validNumbers[indexToDouble] * 2
			if doubleNum > 9 {
				doubleNum -= 9
			}

			luhnNums[i] = doubleNum
			totalSum += doubleNum
			indexToDouble -= 2

		} else {
			luhnNums[i] = validNumbers[i]
			totalSum += validNumbers[i]
		}
	}

	return totalSum%10 == 0
}
