package luhn

import (
	"fmt"
	"regexp"
	"strconv"
)

//Valid determins if a given number is valid based on the Luhn formula
//'0' is a character -- make a comparison to see if each char is >= '0' and <= '9'
func Valid(number string) bool {
	//var isNumber, _ = regexp.MatchString("[0-9]+")
	
		var numbers []byte
		
		for _, val := range number {
			numString := string(val)
			var isNumber, _ = regexp.MatchString("[0-9]+", numString)
			
			if !isNumber {
				continue
			}
			updatedVal, _ := strconv.Atoi(numString)
			numbers = append(numbers, byte(updatedVal))
		}
		fmt.Println(numbers)
		return true
	
	}

// func getNumbers(number string) []string {
// 	matchedNums, _ := regexp.Compile("[0-9]+")
// 	found := matchedNums.FindAllString(number, -1)
// 	return found

// }
