package dynamic_programming

func HouseRobber(nums []int) int {
	// Total number of houses
	n := len(nums)
	// Special cases
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	// Lookup table for remaining houses
	lookup := make([]int, n)
	lookup[0] = nums[0]
	lookup[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		lookup[i] = max(lookup[i-2]+nums[i], lookup[i-1])
	}
	return lookup[n-1]
}
