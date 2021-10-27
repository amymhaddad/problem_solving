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

//Name returns a new Robot name
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	name := getName()

	for {
		_, found := usedNames[name]
		if found || len(name) < 5 {
			name = getName()
			continue
		}
		usedNames[name] = true
		r.name = name
		break
	}
	return name, nil
}

//Reset resets a robot's name to be empty
func (r *Robot) Reset() {
	r.name = ""
}

func getName() string {
	ch1 := rand.Intn(26) + 'A'
	ch2 := rand.Intn(26) + 'A'
	digits := rand.Intn(1000)
	return fmt.Sprintf("%c%c%d", ch1, ch2, digits)
}
