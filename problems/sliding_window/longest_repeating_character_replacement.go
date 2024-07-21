package sliding_window

func CharacterReplacement(s string, k int) int {
	// Special case
	if len(s) == 0 || k < 0 {
		return 0
	}
	// Left and right pointers for the sliding window
	left, right := 0, 0
	// Array to store the frequencies of characters in the string
	frequencies := [26]int{}
	// Count of most popular character so far
	maxCount := 0
	// Longest length
	longestLength := 0
	// Process the string
	for right < len(s) {
		// Frequency of the current character
		frequencies[s[right]-'A']++
		frequency := frequencies[s[right]-'A']
		// Update max count, if required
		maxCount = max(maxCount, frequency)
		// If there are more than k characters as there are most
		// popular character, we shift the window
		for right-left+1-maxCount > k {
			frequencies[s[left]-'A']--
			left++
		}
		longestLength = max(longestLength, right-left+1)
		right++
	}
	return longestLength
}
