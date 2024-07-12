package array

func LongestConsecutiveSequence(nums []int) int {
	// Special case
	if len(nums) == 0 {
		return 0
	}
	// Hash set to store all elements in the array for O(1) retrieval
	elements := make(map[int]bool)
	for _, num := range nums {
		elements[num] = true
	}
	// Length of the longest consecutive sequence
	lcsLength := 0
	// Process elements in the array
	for _, num := range nums {
		// Current element
		currentElement := num
		// If the previous number to the currentElement is not in the set,
		// it means this element must be the first number of a set
		if !elements[currentElement-1] {
			currentLength := 1
			for elements[currentElement+1] {
				currentLength++
				currentElement++
			}
			lcsLength = max(lcsLength, currentLength)
		}
	}
	return lcsLength
}
