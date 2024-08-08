package dynamic_programming

func HouseRobberII(nums []int) int {
	// Total number of houses
	n := len(nums)
	// Special case
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	return max(robHelper(nums, 0, n-2), robHelper(nums, 1, n-1))
}

func robHelper(nums []int, start int, end int) int {
	previous, current := 0, 0
	for i := start; i <= end; i++ {
		temp := max(previous+nums[i], current)
		previous = current
		current = temp
	}
	return current
}
