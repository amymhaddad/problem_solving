package lasagna

const (
	noodleGrams = 50
	sauceLiters = 0.2
)

//PreparationTime determines the amount of time the provided layers will take to
//prepare
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return len(layers) * time

}

//Quantities determines the amount of of noodles and sauce needed
func Quantities(ingredients []string) (int, float64) {
	var totalNoodleLayers, totalSauceLayers int
	
	for _, ingredient := range ingredients {
		switch ingredient {
		case "noodles":
			totalNoodleLayers++
		case "sauce":
			totalSauceLayers++
		}
	}

	return totalNoodleLayers * noodleGrams, float64(totalSauceLayers) * sauceLiters
}

//AddSecretIngredient creates a new slice of ingredients that includes a
//friend's special ingredient
func AddSecretIngredient(friendList []string, myList []string) []string {
	myList = append(myList, friendList[len(friendList)-1])
	return myList
}

//ScaleRecipe scales the lasagna recipe
func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaledQuantites := make([]float64, len(quantities))
	portion := float64(portions) / 2

	for i, quantity := range quantities {
		scaledQuantites[i] = quantity * portion
	}
	return scaledQuantites
}
