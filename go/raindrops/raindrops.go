package raindrops

import (
	"strconv"
)

// Convert converts a number into a string of raindrop sounds
func Convert(num int) string {
	var raindropSounds string

	if num%3 == 0 {
		raindropSounds += "Pling"
	}

	if num%5 == 0 {
		raindropSounds += "Plang"
	}

	if num%7 == 0 {
		raindropSounds += "Plong"
	}

	if len(raindropSounds) == 0 {
		return strconv.Itoa(num)
	}

	return raindropSounds

}
