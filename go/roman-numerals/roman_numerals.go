package romannumerals

import (
	"bytes"
	"errors"
)

//Solution 1
var decimals = [13]int{
	1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1,
}

var romanLetters = [13]string{
	"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I",
}

// ToRomanNumeral receives an integer and returns its Roman Numeral equivalent
func ToRomanNumeral(num int) (string, error) {
	if num <= 0 || num > 3000 {
		return "", errors.New("arabic number is out of range")
	}

	var i int
	var buffer bytes.Buffer

	for num > 0 {
		for decimals[i] > num {
			i++
		}

		buffer.WriteString(romanLetters[i])
		num -= decimals[i]
	}
	return buffer.String(), nil

}

//Solution 1
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
