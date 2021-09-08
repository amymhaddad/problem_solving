package listops

//IntList contains a slice of integers
type IntList []int

type binFunc func(int, int) int

type predFunc func(int) bool

type unaryFunc func(int) int

//Foldl takes a function and an initial accumulator and applies
//them to a slice of numbers -- starting from the left
func (nums IntList) Foldl(fn binFunc, initial int) int {
	if nums.Length() == 0 {
		return initial
	}

	for i := range nums {
		initial = fn(initial, nums[i])
	}
	return initial
}

//Foldr takes a function and an initial accumulator and applies
//them to a slice of numbers -- starting from the right
func (nums IntList) Foldr(fn binFunc, initial int) int {
	if nums.Length() == 0 {
		return initial
	}

	for i := len(nums) - 1; i >= 0; i-- {
		initial = fn(nums[i], initial)
	}

	return initial
}

//Filter takes a function and applies it to a slice of numbers, and returns a
//slice of numbers for which the function is true
func (nums IntList) Filter(fn predFunc) IntList {
	filteredVals := IntList{}

	for i := range nums {
		if fn(nums[i]) {
			filteredVals = append(filteredVals, nums[i])
		}
	}

	return filteredVals

}

//Length returns the length of a slice of numbers
func (nums IntList) Length() int {
	return len(nums)
}

//Map takes a function and a slice of numbers, and returns a slice with the
//function applied to each number in the slice
func (nums IntList) Map(fn unaryFunc) IntList {
	mappedList := make(IntList, len(nums))

	for i, val := range nums {
		mappedList[i] = fn(val)
	}

	return mappedList

}

//Reverse reverses a given slice of numbers
func (nums IntList) Reverse() IntList {
	reversed := make(IntList, len(nums))

	for i := range nums {
		reversed[i] = nums[nums.Length()-1-i]
	}

	return reversed

}

//Append adds all items from the second slice of numbers to the first
func (nums IntList) Append(moreNums IntList) IntList {
	for i := range moreNums {
		nums = append(nums, moreNums[i])
	}

	return nums

}

//Concat combines all elements from multiple slices into a single slice
func (nums IntList) Concat(args []IntList) IntList {
	for _, list := range args {
		nums = append(nums, list...)
	}

	return nums
}
