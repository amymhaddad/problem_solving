package leap

// IsLeapYear determines if a year is a leap year
func IsLeapYear(year int) bool {
	divByFour := year%4 == 0
	divByOneHundred := year%100 == 0
	divByFourHundred := year%400 == 0

	switch {
	case !divByFour:
		return false
	case divByFour && divByOneHundred && !divByFourHundred:
		return false
	default:
		return true
	}

}
