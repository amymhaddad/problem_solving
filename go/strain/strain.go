package strain

//Ints contains a slice of one or more integers
type Ints []int

//Lists contains one or more integer slices
type Lists [][]int

//Strings contains a slice of one or more strings
type Strings []string

//Keep is a method that returns a new collection containing those elements where the predicate is true
func (nums Ints) Keep(fn func(int) bool) Ints {
	var results Ints

	for _, val := range nums {
		if fn(val) {
			results = append(results, val)
		}
	}
	return results

}

//Discard returns a new collection containing those elements where the predicate is false.
func (nums Ints) Discard(fn func(int) bool) Ints {

	var results Ints

	for _, val := range nums {
		if !fn(val) {
			results = append(results, val)
		}
	}

	return results

}

//Keep Strings is a method that returns a new collection containing those elements where the predicate is true
func (words Strings) Keep(fn func(string) bool) Strings {
	var results Strings

	for _, val := range words {
		if fn(val) {
			results = append(results, val)
		}
	}

	return results
}

//Keep Lists is a method that returns a new collection containing those elements where the predicate is true
func (lists Lists) Keep(fn func(int) bool) Lists {
	var results Lists

	for _, list := range lists {
		i := len(list) - 1
		for i > 0 {
			if !fn(list[i]) {
				break
			}
			i--
		}
		if i == 0 {
			results = append(results, list)
		}

	}
	return results
}
