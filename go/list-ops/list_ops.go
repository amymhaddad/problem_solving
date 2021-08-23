package listops

//IntList is type slice
type IntList []int

type binFunc func(int, int) int

type predFunc func(int) bool

type unaryFunc func(int) int

//Foldl takes a function and an initial accumulator and applies
//them to a slice of numbers -- starting from the left
func (nums IntList) Foldl(fn binFunc, initial int) int {
	if len(nums) == 0 {
		return initial
	}

	var total int

	for i := range nums {
		initial = fn(initial, nums[i])
		total = initial
	}
	return total
}

//Foldr takes a function and an initial accumulator and applies
//them to a slice of numbers -- starting from the right
func (nums IntList) Foldr(fn binFunc, initial int) int {
	if len(nums) == 0 {
		return initial
	}

	var total int

	for i := len(nums) - 1; i >= 0; i-- {
		initial = fn(nums[i], initial)
		total = initial
	}
	return total
}

//Filter takes a function and applies it to a slice of numbers, and returns a
//slice of numbers for which the function is true
func (nums IntList) Filter(fn predFunc) IntList {
	filteredNums := IntList{}

	for _, value := range nums {
		if fn(value) {
			filteredNums = append(filteredNums, value)
		}
	}
	return filteredNums
}

//Length returns the length of a slice of numbers
func (nums IntList) Length() int {
	return len(nums)
}

//Map takes a function and a slice of numbers, and returns a slice with the
//function applied to each number in the slice
func (nums IntList) Map(fn unaryFunc) IntList {
	mappedNums := make(IntList, len(nums))

	for i := 0; i < len(mappedNums); i++ {
		mappedNums[i] = fn(nums[i])
	}
	return mappedNums
}

//Reverse reverses a given slice of numbers
func (nums IntList) Reverse() IntList {
	reversed := make(IntList, len(nums))
	var numIndex int

	for i := len(reversed) - 1; i >= 0; i-- {
		reversed[i] = nums[numIndex]
		numIndex++
	}
	return reversed

}

//Append adds all items from the second slice of numbers to the first
func (nums IntList) Append(secondNums IntList) IntList {
	for _, value := range secondNums {
		nums = append(nums, value)
	}
	return nums

}

//Concat combines all elements from multiple slices into a single slice
func (nums IntList) Concat(multiNums []IntList) IntList {
	for _, vals := range multiNums {
		for i := range vals {
			nums = append(nums, vals[i])
		}
	}
	return nums
}
