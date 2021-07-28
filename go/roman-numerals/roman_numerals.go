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

var mappings = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

// ToRomanNumeral receives an integer and returns its Roman Numeral equivalent
func ToRomanNumeral(num int) (string, error) {
	invalidArabicNumber := num <= 0 || num > 3000

	if invalidArabicNumber {
		return "", errors.New("invalid arabic number")
	}

	var romanLetters string
	for num > 0 {
		for _, val := range numsToLetters {
			if val.arabic == num {
				romanLetters += val.roman
				num -= val.arabic
			}
		}

		if num == 0 {
			break
		}

		if num < 4 {
			romanLetters += mappings[1]
			num--
		}

		if num > 5 && num < 9 {
			romanLetters += mappings[5]
			num -= 5
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
