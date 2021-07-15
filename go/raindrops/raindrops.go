package main

import (
	"fmt"
	//"strconv"
)

//Raindrops converts a number into a string of raindrop sounds
func Raindrops(num int) []string {
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
	fmt.Println(raindropSounds)

	return raindropSounds
	// switch {
	// 	case num % 3 == 0:
	// 		raindropSounds = append(raindropSounds, sounds[3])
	// 		fallthrough
	// 	case num % 5 == 0:
	// 		raindropSounds = append(raindropSounds, sounds[3])
	// 	case num
	//
	// }

}

func main() {
	fmt.Println(Raindrops(28))

}
