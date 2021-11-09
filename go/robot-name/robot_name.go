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

const nameLength = 5

func init() {
	rand.Seed(time.Now().UnixNano())
}

//Name creates a robot name
func (r *Robot) Name() (string, error) {
	if len(r.name) == nameLength {
		return r.name, nil
	}

	r.name = generateName()
	for usedNames[r.name] || len(r.name) < nameLength {
		r.name = generateName()
	}
	usedNames[r.name] = true
	return r.name, nil
}

//Reset resets a robot's name
func (r *Robot) Reset() {
	r.name = ""
}

func generateName() string {
	ch1 := rand.Intn(26) + 'A'
	ch2 := rand.Intn(26) + 'A'
	digits := rand.Intn(1000)
	return fmt.Sprintf("%c%c%d", ch1, ch2, digits)
}
