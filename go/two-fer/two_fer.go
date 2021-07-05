package twofer

//Given a name, return a string with the message
func ShareWith(name string) string {
	if len(name) > 0 {
		return "One for " + name + ", one for me."
	} else {
		return "One for you, one for me."
	}
}
