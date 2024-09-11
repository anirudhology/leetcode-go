package sliding_window

func MinimumSizeSubarraySum(target int, nums []int) int {
	// Special case
	if len(nums) == 0 {
		return 0
	}
	// Left and right pointers of the window
	left, right := 0, 0
	// Minimum length of the window
	minLength := 1<<31 - 1
	// Window sum
	windowSum := 0
	// Process the array
	for right < len(nums) {
		windowSum = windowSum + nums[right]
		right++
		// Squeeze the window, if possible
		for windowSum >= target {
			minLength = min(minLength, right-left)
			windowSum = windowSum - nums[left]
			left++
		}
	}
	if minLength == (1<<31 - 1) {
		return 0
	}
	return minLength
}
