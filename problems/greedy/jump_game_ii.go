package greedy

func JumpGameII(nums []int) int {
	// Total number of jumps requires
	jumps := 0
	// Pointers for the window
	left, right := 0, 0
	// Process the elements in the array
	for right < len(nums)-1 {
		// longest jump from the current position
		longest := 0
		// Now, we will try every position in this window
		for i := left; i <= right; i++ {
			longest = max(longest, i+nums[i])
		}
		// Update the pointers for new window
		left = right + 1
		right = longest
		jumps++
	}
	return jumps
}
