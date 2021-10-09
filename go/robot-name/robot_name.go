package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

//Robot contains a name and exists in a factory
type Robot struct {
	name string
}

//storing each name, not each robot instance of name
//Can't add to an uninit map. Must use make() or literal syntax
var cache = make(map[string]bool)
var nameLength = 5
var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var digits = []rune("1234567890")

func init() {
	rand.Seed(time.Now().UnixNano())
}

//Name gets a new name for a robot
//Extract the random name logic into a separate function -- don't repeatidly
//call the method bc I only need to repeatdily call a portion of what's
//currently in the method
func (r *Robot) Name() (string, error) {
	name := getNewName()
	_, found := cache[string(name)]
	fmt.Println("H: ", found)
	if found {
		delete(cache, name)
		return "", errors.New("name is not unique")
	}

	cache[string(name)] = true
	r.name = string(name)
	return string(name), nil

}

func getNewName() string {
	name := make([]rune, nameLength)
	for i := range name {
		if i < 2 {
			name[i] = letters[rand.Intn(len(letters))]
		} else {
			name[i] = digits[rand.Intn(len(digits))]
		}
	}
	return string(name)
}
