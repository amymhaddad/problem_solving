package lasagna

//PreparationTime determines the amount of time the provided layers will take to
//prepare
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		return 2 * len(layers)
	}
	return len(layers) * time

}

// TODO: define the 'Quantities()' function

// TODO: define the 'AddSecretIngredient()' function

// TODO: define the 'ScaleRecipe()' function
