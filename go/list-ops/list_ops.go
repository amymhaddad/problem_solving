package listops

import "fmt"

//IntList contains a slice of integers
type IntList []int

//must note what all the inputs and outputs are
type binFunc func(int, int) int

func (nums IntList) Foldl(fn binFunc, initial int) int {
	if len(nums) == 0 {
		return initial
	}

	for i := range nums {
		initial = fn(initial, nums[i])
	}
	return initial
}

func (nums IntList) Foldr(fn binFunc, initial int) int {
	if len(nums) == 0 {
		return initial
	}

	for i := len(nums) - 1; i >= 0; i-- {
		fmt.Println("HERE", i)
		initial = fn(initial, nums[i])
	}

	return initial

}
