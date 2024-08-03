package backtracking

func Permute(nums []int) [][]int {
	// List to store final output
	var permutations [][]int
	// Special case
	if len(nums) == 0 {
		return permutations
	}
	// Perform backtracking
	backtrackForPermutations(nums, []int{}, &permutations)
	return permutations
}

func backtrackForPermutations(nums []int, permutation []int, permutations *[][]int) {
	if len(permutation) == len(nums) {
		// Add a copy of the current subset to the powerSet
		temp := make([]int, len(permutation))
		copy(temp, permutation)
		*permutations = append(*permutations, temp)
	} else {
		for i := 0; i < len(nums); i++ {
			if contains(permutation[:], nums[i]) {
				continue
			}
			permutation = append(permutation, nums[i])
			backtrackForPermutations(nums, permutation, permutations)
			permutation = permutation[:len(permutation)-1]
		}
	}
}

func contains(arr []int, element int) bool {
	for _, x := range arr {
		if x == element {
			return true
		}
	}
	return false
}
