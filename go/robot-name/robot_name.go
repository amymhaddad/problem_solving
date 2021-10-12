package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

//Robot contains a name and exists in a factory
type Robot struct {
	name string
}

var usedNames = make(map[string]bool)
var random = rand.New(rand.NewSource(time.Now().UnixNano()))
var maxNamespace = 26 * 26 * 10 * 10 * 10
var nameLength = 5

//Name gets a new name for a robot
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	if len(usedNames) >= maxNamespace {
		return "", fmt.Errorf("namespace is exhausted")
	}

	r.name = getNewName()

	for usedNames[r.name] || len(r.name) < nameLength {
		r.name = getNewName()
	}

	usedNames[r.name] = true
	return r.name, nil

}

//Reset a robot name to a new name
func (r *Robot) Reset() {
	r.name = ""
}

func getNewName() string {
	var name string

	letter1 := random.Intn(26) + 'A'
	letter2 := random.Intn(26) + 'A'
	digit := random.Intn(1000)
	name = fmt.Sprintf("%c%c%d", letter1, letter2, digit)

	return name
}
