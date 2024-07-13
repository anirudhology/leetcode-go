package array

func TwoSumII(numbers []int, target int) []int {
	// Special case
	if len(numbers) < 2 {
		return nil
	}
	// Left and right pointers
	left := 0
	right := len(numbers) - 1
	// Process the array from both ends
	for left <= right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return nil
}
