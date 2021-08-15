package luhn

import (
	"regexp"
)

//Valid determins if a given is valid based on the Luhn formula
func Valid(number string) bool {

	numbers := getNumbers(number)
	if len(numbers) <= 1 {
		return false
	}
	return true
}

func getNumbers(number string) []string {
	matchedNums, _ := regexp.Compile("[0-9]+")
	found := matchedNums.FindAllString(number, -1)
	return found

}
