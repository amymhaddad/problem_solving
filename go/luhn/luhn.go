package luhn

import (
	"strings"
)

//'0' is a character -- make a comparison to see if each char is >= '0' and <= '9'
//Valid determins if a given number is valid based on the Luhn formula
func Valid(num string) bool {

	// num = strings.ReplaceAll(num, " ", "")

	// if len(num) <= 1 || (num[len(num)-1] >= 97 && num[len(num)-1] <= 122) {
	// 	return false
	// }

	if !isValid(num) {
		return false
	}

	number := []byte(num)
	var newNums []int

	for i := 0; i < len(number); i++ {
		validNums := number[i] >= '0' && number[i] <= '9'
		if validNums {
			newNums = append(newNums, int(number[i]-'0'))
		}

		// if !validNums {
		// 	return false
		// }
		//This works bc only looking for nums 0-9.
		//THis won't work for other chars bc it'll exeed the amount
		//newNums = append(newNums, int(number[i]-'0'))
	}

	luhnNums := make([]int, len(newNums))
	var totalSum int
	indexToDouble := len(newNums) - 2

	for i := len(newNums) - 1; i >= 0; i-- {
		if i == indexToDouble && indexToDouble >= 0 {
			doubleNum := newNums[indexToDouble] * 2
			if doubleNum > 9 {
				doubleNum -= 9
			} else {
				luhnNums[i] = doubleNum
			}
			totalSum += doubleNum
			indexToDouble -= 2
		} else {
			luhnNums[i] = newNums[i]
			totalSum += newNums[i]
		}

	}

	return totalSum%10 == 0
}

func isValid(num string) bool {
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
