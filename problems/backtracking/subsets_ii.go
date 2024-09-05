package backtracking

import "sort"

func SubsetsWithDup(nums []int) [][]int {
	// List to store final output
	var powerSet [][]int
	// Special case
	if len(nums) == 0 {
		return powerSet
	}
	// Sort the array
	sort.Ints(nums)
	// Perform backtracking
	backtrackSubsetsII(nums, 0, []int{}, &powerSet)
	return powerSet
}

func backtrackSubsetsII(nums []int, index int, subset []int, powerSet *[][]int) {
	temp := make([]int, len(subset))
	copy(temp, subset)
	*powerSet = append(*powerSet, temp)
	for i := index; i < len(nums); i++ {
		if i > index && nums[i] == nums[i-1] {
			continue
		}
		subset = append(subset, nums[i])
		backtrackSubsetsII(nums, i+1, subset, powerSet)
		subset = subset[:len(subset)-1]
	}
}
