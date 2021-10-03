package robotname

import (
	"math/rand"
	"time"
	"errors"
)

//Robot contains a name and exists in a factory
type Robot struct {
	name string
}

var cache map[string]bool
var nameLength = 5
var letters    = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var digits = []rune("1234567890")


func init() {
	rand.Seed(time.Now().UnixNano())
}
//Read common_test
func Name() string {
	name := make([]rune, nameLength)
	for i := range name {
		if i <= 2 {
			name[i] = letters[rand.Intn(len(letters))]
		} else {
			name[i] = digits[rand.Intn(len(digits))]
		}
}

//Check set up of method.
func (r *Robot) getName(name string) (string, error) {
	_, found := cache[string(name)]
	if found {
		return "", errors.New("name is not unique")
	}
	r.name = name
	return string(name), nil

}