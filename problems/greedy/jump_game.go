package greedy

func JumpGame(nums []int) bool {
	// Current jump
	jump := nums[0]
	// Process through the remaining array
	for i := 1; i < len(nums); i++ {
		if jump < i {
			return false
		}
		jump = max(jump, i+nums[i])
	}
	return true
}
