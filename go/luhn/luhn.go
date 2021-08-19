package luhn

import (
	"strings"
)

//'0' is a character -- make a comparison to see if each char is >= '0' and <= '9'
//Valid determins if a given number is valid based on the Luhn formula
func Valid(num string) bool {

	if !isValidString(num) {
		return false
	}

	validNumbers := getValidNums(num)

	luhnNums := make([]int, len(validNumbers))
	var totalSum int
	indexToDouble := len(validNumbers) - 2

	for i := len(validNumbers) - 1; i >= 0; i-- {
		if i == indexToDouble && indexToDouble >= 0 {
			doubleNum := validNumbers[indexToDouble] * 2
			if doubleNum > 9 {
				doubleNum -= 9
			} else {
				luhnNums[i] = doubleNum
			}
			totalSum += doubleNum
			indexToDouble -= 2
		} else {
			luhnNums[i] = validNumbers[i]
			totalSum += validNumbers[i]
		}

	}

	return totalSum%10 == 0
}

func isValidString(num string) bool {
	num = strings.ReplaceAll(num, " ", "")

	if len(num) <= 1 || (num[len(num)-1] >= 97 && num[len(num)-1] <= 122) {
		return false
	}

	number := []byte(num)

	for i := 0; i < len(number); i++ {
		validNums := number[i] >= '0' && number[i] <= '9'

		if !validNums {
			return false
		}
	}
	return true

}

func getValidNums(num string) []int {
	number := []byte(num)
	var validNumbers []int

	for i := 0; i < len(number); i++ {
		validNums := number[i] >= '0' && number[i] <= '9'
		if validNums {
			validNumbers = append(validNumbers, int(number[i]-'0'))
		}
	}
	return validNumbers
}
