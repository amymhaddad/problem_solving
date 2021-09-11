package lasagna

//OvenTime returns the amount of time the lasagna should be in the oven
func OvenTime() int {
	return 40
}

//RemainingOvenTime calculates the remaining time in oven
func RemainingOvenTime(currBakingTime int) int {
	return OvenTime() - currBakingTime
}

//PreparationTime returns the amount of time it takes to prep each layer
func PreparationTime(layers int) int {
	return layers * 2
}

//ElapsedTime calculates the total working time in minutes
func ElapsedTime(layers, time int) int {
	return PreparationTime(layers) + time
}
