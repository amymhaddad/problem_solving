package romannumerals

import (
	"errors"
)

// ArabicToRoman contains an arabic number and its roman numeral equivalent
type ArabicToRoman struct {
	arabic int
	roman  string
}

var numsToLetters = [12]ArabicToRoman{
	ArabicToRoman{1000, "M"},
	ArabicToRoman{900, "CM"},
	ArabicToRoman{500, "D"},
	ArabicToRoman{100, "C"},
	ArabicToRoman{90, "XC"},
	ArabicToRoman{50, "L"},
	ArabicToRoman{40, "XL"},
	ArabicToRoman{10, "X"},
	ArabicToRoman{9, "IX"},
	ArabicToRoman{5, "V"},
	ArabicToRoman{4, "IV"},
	ArabicToRoman{1, "I"},
}

// ToRomanNumeral receives an integer and returns its Roman Numeral equivalent
func ToRomanNumeral(num int) (string, error) {
	invalidArabicNumber := num <= 0 || num > 3000

	if invalidArabicNumber {
		return "", errors.New("invalid arabic number")
	}

	//The loop doesn't restart at the top of the struct
	//I want to restart the loop each time I add a letter to romanLetters
	//I even used "continue" and that didn't restart the loop
	var romanLetters string
	for num > 0 {
		for _, val := range numsToLetters {
			if num >= val.arabic {
				// fmt.Println("\nnum before: ", num)
				// fmt.Println("\narabic: ", val.arabic)
				romanLetters += val.roman
				num -= val.arabic
				break
			}
		}

		if num == 0 {
			break
		}

	}
	return romanLetters, nil
}

/*
Iterate thru the array of structs, bc everything is ordered in a struct -- not in a map. Plus
I move from largest to smallest: 1000 to 1.

On each iteartion, check if val.arabic (ie, the current number of iteration in array of structs)
is >= to the current number.
IF so, add the letter and reduce the number by the val.arabic amount.

Delete the map. Keep the break.

*/

// var mappings = map[int]string{
// 	1000: "M",
// 	900:  "CM",
// 	500:  "D",
// 	100:  "C",
// 	90:   "XC",
// 	50:   "L",
// 	40:   "XL",
// 	10:   "X",
// 	9:    "IX",
// 	5:    "V",
// 	4:    "IV",
// 	1:    "I",
// }
