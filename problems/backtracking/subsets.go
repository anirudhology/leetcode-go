package backtracking

import "sort"

func Subsets(nums []int) [][]int {
	// List to store the final output
	var powerSet [][]int
	// Special case
	if len(nums) == 0 {
		return powerSet
	}
	// Sort the array
	sort.Ints(nums)
	// Perform backtracking
	backtrack(nums, 0, []int{}, &powerSet)
	return powerSet
}

func backtrack(nums []int, index int, subset []int, powerSet *[][]int) {
	// Add a copy of the current subset to the powerSet
	temp := make([]int, len(subset))
	copy(temp, subset)
	*powerSet = append(*powerSet, temp)
	for i := index; i < len(nums); i++ {
		subset = append(subset, nums[i])
		backtrack(nums, i+1, subset, powerSet)
		subset = subset[:len(subset)-1]
	}
}
