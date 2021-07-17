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

	orbitPeriod, ok := planetOrbitalPeriod[planet]
	if !ok {
		return seconds / earthSeconds
	}
	return seconds / (orbitPeriod * earthSeconds)
}
