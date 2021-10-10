package robotname

import (
	"errors"
	"math/rand"
	"time"
)

//Robot contains a name and exists in a factory
type Robot struct {
	name string
}

const nameLength = 5

var cache = make(map[string]bool)
var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var digits = []rune("1234567890")

func init() {
	rand.Seed(time.Now().UnixNano())
}

//Name gets a new name for a robot
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}
	name := getNewName()

	repeatedName := getOrSetCache(name)
	if repeatedName {
		return r.name, errors.New("name is not unique")
	}
	r.name = string(name)
	return string(name), nil

}

//Reset a robot name to a new name
func (r *Robot) Reset() string {
	name := getNewName()

	repeatedName := getOrSetCache(name)
	if repeatedName {
		name = getNewName()
	}
	r.name = string(name)
	return string(name)
}

func getOrSetCache(name string) bool {
	_, found := cache[string(name)]
	if !found {
		cache[string(name)] = true
	}
	return found
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
