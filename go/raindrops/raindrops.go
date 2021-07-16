package raindrops

import (
	"strconv"
	"strings"
)

// Convert converts a number into a string of raindrop sounds
func Convert(num int) string {
	sounds := map[int]string{
		3: "Pling",
		5: "Plang",
		7: "Plong",
	}

	var raindropSounds []string

	for factor, sound := range sounds {
		if num%factor == 0 {
			raindropSounds = append(raindropSounds, sound)
		}
	}

	if len(raindropSounds) == 0 {
		return strconv.Itoa(num)
	}

	return strings.Join(raindropSounds, "")

}
