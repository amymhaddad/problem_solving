package isogram

var letters = make(map[rune]bool)

// IsIsogram determines if a given string is an isogram
func IsIsogram(phrase string) bool {

	for _, char := range phrase {
		if _, ok := letters[char]; ok {
			return false
		}
		//	fmt.Println("\nchar: ", char, "\nletters: ", letters)
		letters[char] = true
	}
	return true
}
