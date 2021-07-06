//twofer problem
package twofer

import "fmt"

//Given a name, return a string with the message
func ShareWith(name string) string {
	if len(name) == 0 {
		return "One for you, one for me."
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
