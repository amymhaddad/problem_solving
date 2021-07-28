package romannumerals

import (
	"errors"
)

// ArabicToRoman contains an arabic number and its roman numeral equivalent
// type ArabicToRoman struct {
// 	arabic int
// 	roman  string
// }

// var mappings = [9]ArabicToRoman{
// 	ArabicToRoman{1000, "M"},
// 	ArabicToRoman{500, "D"},
// 	ArabicToRoman{100, "C"},
// 	ArabicToRoman{50, "L"},
// 	ArabicToRoman{10, "X"},
// 	ArabicToRoman{9, "IV"},
// 	ArabicToRoman{5, "V"},
// 	ArabicToRoman{4, "IV"},
// 	ArabicToRoman{1, "I"},
// }

var mappings = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IV",
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
		// if roman, ok := mappings[num]; ok {
		// 	romanLetters += roman
		// 	num -= num
		// }

		if num < 4 {
			romanLetters += mappings[1]
			num--
		}

	}
	return romanLetters, nil
}
