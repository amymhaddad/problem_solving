package romannumerals

import (
	"errors"
)

type arabicToRoman struct {
	arabic int
	roman  string
}

var numsToLetters = [13]arabicToRoman{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// ToRomanNumeral receives an integer and returns its Roman Numeral equivalent
func ToRomanNumeral(num int) (string, error) {
	invalidArabicNumber := num <= 0 || num > 3000

	if invalidArabicNumber {
		return "", errors.New("arabic number is out of range")
	}

	var romanLetters string
	for num > 0 {
		for _, val := range numsToLetters {
			if num >= val.arabic {
				romanLetters += val.roman
				num -= val.arabic
				break
			}
		}
	}
	return romanLetters, nil
}
