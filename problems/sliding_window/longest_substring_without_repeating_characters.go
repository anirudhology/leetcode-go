package sliding_window

func LongestSubstringWithoutRepeatingCharacters(s string) int {
	// Special case
	if len(s) == 0 {
		return 0
	}
	// Length of the string
	n := len(s)
	// Slow and fast pointers for sliding window
	slow, fast := 0, 0
	// Set to store unique characters
	uniques := map[byte]int{}
	// Longest substring length
	longestLength := 0
	// Process the string
	for fast < n {
		if _, ok := uniques[s[fast]]; !ok {
			uniques[s[fast]] = fast
			fast++
			longestLength = max(longestLength, fast-slow)
		} else {
			delete(uniques, s[slow])
			slow++
		}
	}
	return longestLength
}
