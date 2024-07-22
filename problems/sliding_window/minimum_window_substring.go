package sliding_window

import "math"

func MinimumWindowSubstring(s string, t string) string {
	// Array to store the frequencies of characters in t
	frequencies := [128]int{}
	for _, c := range t {
		frequencies[int(c)]++
	}
	// Left and right pointers for the sliding window
	left, right := 0, 0
	// Lengths of both strings
	m, n := len(s), len(t)
	// Minimum length of the window
	minLength := math.MaxInt
	// Start index of the window
	start := 0
	// Process the string
	for right < m {
		// If the current character exists in t as well
		if frequencies[int(s[right])] > 0 {
			n--
		}
		frequencies[int(s[right])]--
		// Expand window
		right++
		for n == 0 {
			if right-left < minLength {
				minLength = right - left
				start = left
			}
			// Shrink the window, if possible
			if frequencies[int(s[left])] == 0 {
				n++
			}
			frequencies[int(s[left])]++
			left++
		}
	}
	if minLength == math.MaxInt {
		return ""
	}
	return s[start : start+minLength]
}
