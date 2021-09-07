package listops

//IntList contains a slice of integers
type IntList []int

//must note what all the inputs and outputs are
type binFunc func(int, int) int

type predFunc func(int) bool

type unaryFunc func(int) int

func (nums IntList) Foldl(fn binFunc, initial int) int {
	if nums.Length() == 0 {
		return initial
	}

	for i := range nums {
		initial = fn(initial, nums[i])
	}
	return initial
}

func (nums IntList) Foldr(fn binFunc, initial int) int {
	if nums.Length() == 0 {
		return initial
	}

	for i := len(nums) - 1; i >= 0; i-- {
		initial = fn(nums[i], initial)
	}

	return initial
}

func (nums IntList) Filter(fn predFunc) IntList {
	filteredVals := IntList{}
	//var filteredVals IntList doesnt' work: it LOOKS like (ie, prints) [] if empty list is encoutnered BUT
	//its type is nil, and that's not what I want. I need to return type IntList. So I must declare and init IntList

	for i := range nums {
		if fn(nums[i]) {
			filteredVals = append(filteredVals, nums[i])
		}
	}

	return filteredVals

}

func (list IntList) Length() int {
	return len(list)
}

func (list IntList) Map(fn unaryFunc) IntList {
	mappedList := make(IntList, len(list))

	for i, val := range list {
		mappedList[i] = fn(val)
	}

	return mappedList

}

func (list IntList) Reverse() IntList {
	reversed := make(IntList, len(list))

	for i := range list {
		reversed[i] = list[list.Length()-1-i]
	}

	return reversed

}
