package strain

//Ints contains a slice of ints
type Ints []int

//Lists contains a slice of slices
type Lists [][]int

//Strings contains a slice of strings
type Strings []string

//Keep returns a new slice of numbers where the predicate is true
func (nums Ints) Keep(fn func(int) bool) Ints {
	var result Ints

	for _, val := range nums {
		if fn(val) {
			result = append(result, val)
		}
	}

	return result
}

//Discard returns a new slice of numbers where the predicate is false
func (nums Ints) Discard(fn func(int) bool) Ints {
	var result Ints

	for _, val := range nums {
		if !fn(val) {
			result = append(result, val)
		}
	}

	return result
}

//Keep returns a new slice of strings where the predicate is true
func (words Strings) Keep(fn func(string) bool) Strings {
	var results Strings

	for _, word := range words {
		if fn(word) {
			results = append(results, word)
		}
	}

	return results
}

//Keep returns a new slice of one or more slices where the predicate is true
func (lists Lists) Keep(fn func([]int) bool) Lists {
	var result Lists

	for _, list := range lists {
		if fn(list) {
			result = append(result, list)
		}
	}

	return result

}
