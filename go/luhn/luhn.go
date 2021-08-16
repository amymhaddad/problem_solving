package luhn

import (
	"fmt"
	"regexp"
	"strings"
)

//Valid determins if a given number is valid based on the Luhn formula
//'0' is a character -- make a comparison to see if each char is >= '0' and <= '9'

func Valid(num string) bool {
	totalNumbers := getNumbers(num)
	if totalNumbers < 1 {
		return false
	}
	number := []byte(num)
	newNum := make([]int, totalNumbers)

	for i := totalNumbers - 1; i >= 0; i-- {
		validNums := number[i] >= '0' && number[i] <= '9'
		// isLetter := number[i] >= 'a' && number[i] <= 'z'

		// if isLetter {
		// 	return false
		// }

		if !validNums {
			continue
		}

		if i == 0 || i%2 == 0 {
			n := number[i] - '0'
			doubleNum := n * 2
			if doubleNum > 9 {
				doubleNum -= 9
			}
			newNum[i] = int(doubleNum)

		} else {
			newNum[i] = int(number[i] - '0')
		}

		fmt.Println(newNum)
      //[8 5 6 9 0 1 8 8 7 0 0 3 8 3 0 6]


		// }
		// 	if validNums {
		// 		newNum = append(newNum, int(number[i]-'0'))
		// 	}
		// }
		// fmt.Println(newNum)

		// reversed := make([]int, len(newNum))

		// counter := 0
		// for i := len(newNum) - 1; i >= 0; i-- {
		// 	reversed[i] = newNum[counter]
		// 	counter++
		// }
		// fmt.Println(reversed)
	}
	return true
}

func getNumbers(number string) int {
	matchedNums, _ := regexp.Compile("[0-9]+")
	found := matchedNums.FindAllString(number, -1)
	return len(strings.Join(found, ""))
}

func main() {
	x := "4539 1488 0343 6467"
	y := Valid(x)
	//y := Valid(x)
	fmt.Println(y)
}


// func getNumbers(number string) []string {
// 	matchedNums, _ := regexp.Compile("[0-9]+")
// 	found := matchedNums.FindAllString(number, -1)
// 	return found

// }
