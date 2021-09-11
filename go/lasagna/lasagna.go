package lasagna

// TODO: define the 'OvenTime()' function

// TODO: define the 'RemainingOvenTime()' function

// TODO: define the 'PreparationTime()' function

// TODO: define the 'ElapsedTime()' function

//OvenTime returns the amount of time the lasagna should be in the oven
func OvenTime() int {
	return 40
}

//RemainingOvenTime calculates the remaining time in oven
func RemainingOvenTime(givenTime int) int {
	totalOvenTime := OvenTime()
	return totalOvenTime - givenTime
}
