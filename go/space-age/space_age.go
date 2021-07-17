package space

//using a type Planet here to be specific -- I want a planet: Mars, Earth, etc
//also use for temperatures (F or C) and time
//Planet is the planet name
type Planet string

const earthSeconds = 31557600

//Age calculates the age of a person on a particular planet
func Age(seconds float64, planet Planet) float64 {
	planetOrbitalPeriod := map[Planet]float64{
		"Mercury": 0.2408467,
	}

	for name, orbitalPeriod := range planetOrbitalPeriod {
		switch name {
		case "Mercury":
			return calculation(seconds, orbitalPeriod)
		default:
			return onEarth(seconds)
		}
	}
	//why do I need another return when I have multi returns in the for loop?
	//	return onEarth(seconds)
}

func onEarth(seconds float64) float64 {
	return seconds / earthSeconds
}

func calculation(seconds float64, orbitalPeriod float64) float64 {
	return seconds / (orbitalPeriod * earthSeconds)
}
