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

//SHould this be Name or getName()
//Name creates a robot name
func (r *Robot) Name() (string, error) {

	name := getName()

	_, found := usedNames[name]

	for found || len(name) < nameLength {
		name = getName()
	}
	usedNames[name] = true
	fmt.Println(name)
	return name, nil

}

func getName() string {
	ch1 := rand.Intn(26) + 'A'
	ch2 := rand.Intn(26) + 'A'
	digits := rand.Intn(1000)
	fmt.Println("chrs", ch1, digits)
	return fmt.Sprintf("%c%c%d", ch1, ch2, digits)

}
