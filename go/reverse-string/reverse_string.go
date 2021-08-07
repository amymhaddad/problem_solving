package reverse

//Reverse receives a string and reverses it
func Reverse(word string) string {
	wordBytes := []byte(word)

	p1 := 0
	p2 := len(wordBytes) - 1

	for p2-p1 >= 1 {
		wordBytes[p1], wordBytes[p2] = wordBytes[p2], wordBytes[p1]
		p2--
		p1++
	}
	return string(wordBytes)

}
