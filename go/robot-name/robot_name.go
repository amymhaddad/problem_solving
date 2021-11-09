package robotname

import (
	"math/rand"
	"time"
)

//Robot contains a name and exists in a factory
type Robot struct {
	name string
}

var usedNames = make(map[string]bool)

func init() {
	rand.Seed(time.Now().UnixNano())
}

//SHould this be Name or getName()
//Name creates a robot name
func (r *Robot) Name() string {

	name := getName()

}

func getName() string {

}
