package leap

// IsLeapYear determines if a year is a leap year
func IsLeapYear(year int) bool {
	divisibleByFour := year%4 == 0
	//fmt.Println("div by 4: ", divisibleByFour)
	return divisibleByFour
}
