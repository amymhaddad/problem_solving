package listops

//IntList contains a slice of integers
type IntList []int

//must note what all the inputs and outputs are
type binFunc func(int, int) int

func (nums IntList) Foldl(fn binFunc, initial int) int {
	//return 15
	accum := initial
	for i := range nums {
	 
		accum = fn(initial, nums[i]) 
	}
	return accum
}

func (l IntList) Foldr(fn binFunc, initial int) int {
	return 15
}
