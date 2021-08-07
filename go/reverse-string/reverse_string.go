package reverse

//Reverse receives a string and reverses it
func Reverse(word string) string {
	runes := []rune(word)

	p1 := 0
	p2 := len(runes) - 1

	for p2-p1 >= 1 {
		runes[p1], runes[p2] = runes[p2], runes[p1]
		p2--
		p1++
	}

	return string(runes)

}
