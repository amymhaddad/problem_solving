package cars

const carsProducedPerHour = 221

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {

	rate := successRate(speed)
	if speed == 0 {
		return float64(speed)
	}
	if speed < 5 {
		return rate * carsProducedPerHour * float64(speed)
	}
	if speed >= 9 {
		return rate * carsProducedPerHour * float64(speed)
	}
	return float64(0.9 * carsProducedPerHour) * float64(speed)
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	panic("not implemented")
}

// successRate is used to calculate the ratio of an item being created without
// error for a given speed
func successRate(speed int) float64 {
	if speed == 0 {
		return 0.0
	}

	if speed < 5 {
		return 1.0
	}

	if speed >= 9 {
		return 0.77
	}

	return 0.9
}
