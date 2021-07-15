package raindrops

import (
	"strconv"
	"strings"
)

//Raindrops converts a number into a string of raindrop sounds
func Raindrops(num int) string {
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

// func main() {
// 	fmt.Println(Raindrops(1))

// }
