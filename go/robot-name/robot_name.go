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

//this doesn't give me a new char each time - but I can use as a value
//var random = rand.New(rand.NewSource(time.Now().UnixNano()))
func init() {
	rand.Seed(time.Now().UnixNano())
}

var usedNames = make(map[string]bool)

//Name returns a new Robot name
func (r *Robot) Name() (string, error) {
	name := getName()
	return name, nil
}

func getName() string {
	ch1 := rand.Intn(26) + 'A'
	ch2 := rand.Intn(26) + 'A'
	digits := rand.Intn(1000)

	return fmt.Sprintf("%c%c%d", ch1, ch2, digits)
}
