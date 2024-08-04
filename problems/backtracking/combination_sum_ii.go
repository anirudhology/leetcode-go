package backtracking

import "sort"

func CombinationSum2(candidates []int, target int) [][]int {
	// List to store final output
	var combinations [][]int
	// Special case
	if len(candidates) == 0 {
		return combinations
	}
	// Sort the array
	sort.Ints(candidates)
	// Perform backtracking
	backtrackForCombinationSumII(candidates, 0, []int{}, &combinations, target)
	return combinations
}

func backtrackForCombinationSumII(candidates []int, index int, combination []int, combinations *[][]int, target int) {
	if target == 0 {
		var temp = make([]int, len(combination))
		copy(temp, combination)
		*combinations = append(*combinations, temp)
		return
	}
	if target < 0 {
		return
	}
	for i := index; i < len(candidates); i++ {
		if i > index && candidates[i] == candidates[i-1] {
			continue
		}
		combination = append(combination, candidates[i])
		backtrackForCombinationSumII(candidates, i+1, combination, combinations, target-candidates[i])
		combination = combination[:len(combination)-1]
	}
}
