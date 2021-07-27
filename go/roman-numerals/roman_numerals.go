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

	switch {
	case invalidArabicNumber:
		return "", errors.New("Invalid Arabic number")
	case num < 5:
		return underFour(num), nil

	case num > 5 && num < 10:
		return underTen(num), nil

	case num > 10 && num < 50:
		return underFifty(num), nil

	default:
		return arabicToRoman[num], nil
	}
}

func underFour(num int) string {
	result := ""
	for i := 0; i < num; i++ {
		result += arabicToRoman[1]
	}

	if len(result) == 4 {
		return "IV"
	}

	return result
}

func underTen(num int) string {
	if 10-num == 1 {
		return "IX"
	}
	return "VI"
}

func underFifty(num int) string {

}
