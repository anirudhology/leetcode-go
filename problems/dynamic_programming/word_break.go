package dynamic_programming

func WordBreak(s string, wordDict []string) bool {
	// Special case
	if len(s) == 0 || len(wordDict) == 0 {
		return false
	}
	n := len(s)
	// Convert list to set for O(1) search
	words := make(map[string]bool)
	for _, word := range wordDict {
		words[word] = true
	}
	// Lookup table to store if the substring until current index
	// lies in the dictionary or not
	lookup := make([]bool, n+1)
	// Empty string is always present
	lookup[0] = true
	// Process the string
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if _, exists := words[s[j:i]]; exists && lookup[j] {
				lookup[i] = true
				break
			}
		}
	}
	return lookup[n]
}
