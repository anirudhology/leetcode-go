package dynamic_programming

func LongestPalindromicSubstring(s string) string {
	// Special case
	if len(s) == 0 {
		return s
	}
	// Start and end pointers for palindromic substring
	start, end := 0, 0
	// Process the string
	for i := 0; i < len(s); i++ {
		// Expand from center for both odd and even indices
		oddLength := expandFromCenter(s, i, i)
		evenLength := expandFromCenter(s, i, i+1)
		// Max of both
		maxLength := max(oddLength, evenLength)
		if maxLength > end-start {
			start = i - int((maxLength-1)/2)
			end = i + int(maxLength/2)
		}
	}
	return s[start : end+1]
}

func expandFromCenter(s string, left int, right int) int {
	if left > right {
		return 0
	}
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}
