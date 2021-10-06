package lasagna

//	"strings"

const (
	noodleGrams = 50
	sauceLiters = 0.2
)

//PreparationTime determines the amount of time the provided layers will take to
//prepare
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		return 2 * len(layers)
	}
	return len(layers) * time

}

//Quantities determines the amount of of noodles and sauce needed
func Quantities(ingredients []string) (int, float64) {
	var totalNoodleLayers int
	var totalSauceLayers int

	for _, ingredient := range ingredients {
		if ingredient == "noodles" {
			totalNoodleLayers++
		}
		if ingredient == "sauce" {
			totalSauceLayers++
		}
	}

	return totalNoodleLayers * noodleGrams, float64(totalSauceLayers) * sauceLiters
}

//AddSecretIngredient creates a new slice of ingredients that includes a
//friend's special ingredient
func AddSecretIngredient(friendList []string, myList []string) []string {
	specialIngredient := friendList[len(friendList)-1]
	newIngredients := make([]string, len(myList))

	for i, ingredient := range myList {
		newIngredients[i] = ingredient
	}
	newIngredients = append(newIngredients, specialIngredient)
	return newIngredients

}

// TODO: define the 'ScaleRecipe()' function
