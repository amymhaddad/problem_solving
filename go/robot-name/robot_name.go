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

//var random = rand.Seed(time.Now().UnixNano())

//Name returns a new Robot name
func (r *Robot) Name() (string, error) {

	rand.Seed(time.Now().UnixNano())
	ch := rand.Intn(65 - 99)
	fmt.Println(ch)

	return "hi", nil
}
