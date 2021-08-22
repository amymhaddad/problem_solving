package listops

type IntList []int

//always look at the test to determine what to write (ie, function or method)
// adding behavior to IntList type

//a function can be a type -- just naming something particular
//type is binFunc; func(type of params) return type
type binFunc func(int, int) int

type predFunc func(int) bool

type unaryFunc func(int) int

//inner function takes 2 ints and returns an int
//the entire method returns an int
//param name is fn and the type binFund
func (nums IntList) Foldl(fn binFunc, initial int) int {
	if len(nums) == 0 {
		return initial
	}

	var total int
	for i := range nums {

		result := fn(initial, nums[i])
		initial = result
		total = result
	}
	return total
}

func (nums IntList) Foldr(fn binFunc, initial int) int {
	if len(nums) == 0 {
		return initial
	}

	var total int
	for i := len(nums) - 1; i >= 0; i-- {
		result := fn(nums[i], initial)
		initial = result
		total = result
	}
	return total

}

//if use IntList in the program, must use {} to initialize it
//It's a method on IntList, so I'm probably going to do something w/it and
//return it
func (nums IntList) Filter(fn predFunc) IntList {
	filteredNums := IntList{}
	if len(nums) == 0 {
		return filteredNums
	}

	for _, value := range nums {
		if fn(value) {
			filteredNums = append(filteredNums, value)
		}
	}
	return filteredNums
}

func (nums IntList) Length() int {
	return len(nums)
}

func (nums IntList) Map(fn unaryFunc) IntList {
	mappedNums := make(IntList, len(nums))

	if len(nums) == 0 {
		return mappedNums
	}

	for i := 0; i < len(mappedNums); i++ {
		mappedNums[i] = fn(nums[i])
	}
	return mappedNums
}

func (nums IntList) Reverse() IntList {
	reversed := make(IntList, len(nums))
	var numIndex int

	for i := len(reversed) - 1; i >= 0; i-- {
		reversed[i] = nums[numIndex]
		numIndex++
	}
	return reversed

}

func (firstNums IntList) Append(secondNums IntList) IntList {
	for _, value := range secondNums {
		firstNums = append(firstNums, value)
	}
	return firstNums

}
