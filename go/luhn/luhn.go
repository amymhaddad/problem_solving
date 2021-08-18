package luhn

import (
	"regexp"
	"strings"
)

//'0' is a character -- make a comparison to see if each char is >= '0' and <= '9'
//Valid determins if a given number is valid based on the Luhn formula
func Valid(num string) bool {
	totalNumbers := getNumbers(num)

	if totalNumbers < 1 {
		return false
	}
	number := []byte(num)
	var newNums []int 
	  
	for i := 0; i < len(number); i++ {
		validNums := number[i] >= '0' && number[i] <= '9'

		if !validNums {
			continue
		}
		newNums = append(newNums, int(number[i] - '0'))
	}
	
	luhnNums := make([]int, len(newNums))
	var totalSum int
	for i := len(newNums) -1; i >= 0; i-- {
		if i == 0 || i % 2 == 0 {
			doubleNum := newNums[i] * 2
			if doubleNum > 9{
				doubleNum -= 9
			} else {
				luhnNums[i] = doubleNum
			}
			totalSum += doubleNum
		} else {
			luhnNums[i] = newNums[i]
			totalSum += newNums[i]
		}
		
	}

	return totalSum % 10 == 0
}

func getNumbers(number string) int {
	matchedNums, _ := regexp.Compile("[0-9]+")
	found := matchedNums.FindAllString(number, -1)
	return len(strings.Join(found, ""))
}