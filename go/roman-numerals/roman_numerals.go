package romannumerals

import "errors"

var arabicToRoman = map[int]string{
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M",
}

// ToRomanNumeral receives an integer and returns its Roman Numeral equivalent
func ToRomanNumeral(num int) (string, error) {
	invalidArabicNumber := num <= 0 || num > 3000

	if invalidArabicNumber {
		return "", errors.New("Invalid Arabic number")
	}

	if num < 4 {
		return underFour(num), nil
	}

	return arabicToRoman[num], nil
}

func underFour(num int) string {
	result := ""
	for i := 0; i < num; i++ {
		result += arabicToRoman[1]
	}

	return result
}
