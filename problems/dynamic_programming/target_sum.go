package dynamic_programming

import "strconv"

func TargetSum(nums []int, target int) int {
	// Special case
	if len(nums) == 0 {
		return 0
	}
	// Map to store results of overlapping subproblems
	lookup := make(map[string]int)
	return evaluate(nums, target, 0, lookup)
}

func evaluate(nums []int, target int, index int, lookup map[string]int) int {
	// Base case
	if index >= len(nums) {
		if target == 0 {
			return 1
		}
		return 0
	}
	// Already solved the overlapping subproblem
	key := strconv.Itoa(index) + "-" + strconv.Itoa(target)
	if value, exists := lookup[key]; exists {
		return value
	}
	// Evaluate ways for both positive and negative signs
	ways := evaluate(nums, target-nums[index], index+1, lookup) + evaluate(nums, target+nums[index], index+1, lookup)
	lookup[key] = ways
	return ways
}
