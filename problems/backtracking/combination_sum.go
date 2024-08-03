package backtracking

import "sort"

func CombinationSum(candidates []int, target int) [][]int {
	// List to store final output
	var result [][]int
	// Special case
	if len(candidates) == 0 {
		return result
	}
	// Sort the array
	sort.Ints(candidates)
	backtrackCombinationSum(candidates, 0, []int{}, target, &result)
	return result
}

func backtrackCombinationSum(candidates []int, index int, combination []int, target int, result *[][]int) {
	// Base case
	if target == 0 {
		temp := make([]int, len(combination))
		copy(temp, combination)
		*result = append(*result, temp)
		return
	}
	if target < 0 {
		return
	}
	for i := index; i < len(candidates); i++ {
		combination = append(combination, candidates[i])
		backtrackCombinationSum(candidates, i, combination, target-candidates[i], result)
		combination = combination[:len(combination)-1]
	}
}
