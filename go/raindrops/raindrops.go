package raindrops

import (
	"strconv"
)

var sounds = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

// Convert converts a number into a string of raindrop sounds
func Convert(num int) string {
	factors := []int{3, 5, 7}

	var raindropSounds string

	for _, factor := range factors {
		if num%factor == 0 {
			raindropSounds += sounds[factor]
		}
	}

	if len(raindropSounds) == 0 {
		return strconv.Itoa(num)
	}

	return raindropSounds

}
