package array

func ContainerWithMostWater(height []int) int {
	// Special case
	if len(height) == 0 {
		return 0
	}
	// Left and right pointers
	left := 0
	right := len(height) - 1
	// Max area
	maxArea := 0
	// Process array from both ends
	for left <= right {
		h := min(height[left], height[right])
		w := right - left
		maxArea = max(h*w, maxArea)
		// Update the pointers accordingly
		if height[left] < height[right] {
			left++
		} else if height[left] > height[right] {
			right--
		} else {
			left++
			right--
		}
	}
	return maxArea
}
