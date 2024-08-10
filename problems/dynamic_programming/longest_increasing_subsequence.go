package dynamic_programming

func LongestIncreasingSubsequence(nums []int) int {
	// Special case
	if len(nums) == 0 {
		return 0
	}
	// Lookup table to store length of LIS until current index
	lookup := make([]int, len(nums))
	// Process the array
	for i := 0; i < len(nums); i++ {
		lookup[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				lookup[i] = max(lookup[i], 1+lookup[j])
			}
		}
	}
	// Max length of subsequences
	maxLength := 0
	for _, element := range lookup {
		maxLength = max(maxLength, element)
	}
	return maxLength
}
