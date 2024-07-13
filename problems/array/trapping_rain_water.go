package array

func TrappingRainWater(height []int) int {
	// Special case
	if len(height) == 0 {
		return 0
	}
	// Left and right pointers
	left := 0
	right := len(height) - 1
	// Variables to keep track of maximum height to the left
	// and maximum height to the right of the current element
	leftHeight := 0
	rightHeight := 0
	// Area of trapped water
	area := 0
	// Process the array from both ends
	for left <= right {
		if height[left] <= height[right] {
			area += max(0, leftHeight-height[left])
			leftHeight = max(leftHeight, height[left])
			left++
		} else {
			area += max(0, rightHeight-height[right])
			rightHeight = max(rightHeight, height[right])
			right--
		}
	}
	return area
}
