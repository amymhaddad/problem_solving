package strain

//Ints contains a slice of one or more integers
type Ints []int

//Lists contains one or more integer slices
type Lists [][]int

//Strings contains a slice of one or more strings
type Strings []string

//Keep filters integers and returns a new collection containing those elements where the predicate is true
func (nums Ints) Keep(fn func(int) bool) Ints {
	var result Ints

	for _, num := range nums {
		if fn(num) {
			result = append(result, num)
		}
	}

	return result

}

//Discard filters integers and returns a new collection containing those elements where the predicate is false.
func (nums Ints) Discard(fn func(int) bool) Ints {
	var result Ints

	for _, num := range nums {
		if !fn(num) {
			result = append(result, num)
		}
	}

	return result

}

//Keep filters strings and returns a new collection containing those elements where the predicate is true
func (words Strings) Keep(fn func(string) bool) Strings {
	var result Strings

	for _, word := range words {
		if fn(word) {
			result = append(result, word)
		}
	}

	return result
}

//Keep filters slices and returns a new collection containing those elements where the predicate is true
func (lists Lists) Keep(fn func([]int) bool) Lists {
	var result Lists

	for _, list := range lists {
		if fn(list) {
			result = append(result, list)
		}
	}

	return result
}
