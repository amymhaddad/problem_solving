package listops

type IntList []int

//always look at the test to determine what to write (ie, function or method)
// adding behavior to IntList type

//a function can be a type -- just naming something particular
//type is binFunc; func(type of params) return type
type binFunc func(int, int) int

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
		result := fn(initial, nums[i])
		initial = result
		total = result
	}
	return total

}
