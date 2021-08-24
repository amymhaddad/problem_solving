package strain

//Ints contains a slice of ints
type Ints []int

//Lists contains a slice of slices
type Lists [][]int

//Strings contains a slice of strings
type Strings []string

func (nums Ints) Keep(fn func(int) bool) Ints {
	var result Ints

	for _, val := range nums {
		if fn(val) {
			result = append(result, val)
		}
	}

	return result
}

