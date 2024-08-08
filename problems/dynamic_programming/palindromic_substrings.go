package dynamic_programming

func PalindromicSubstrings(s string) int {
	// Special case
	if len(s) == 0 {
		return 0
	}
	// Count to store total number of palindromic substrings
	count := 0
	// Process each string by keeping each index as middle and
	// expand both sides from it
	for i := 0; i < len(s); i++ {
		// Odd length
		count += expandFromCenterForPalindromicSubstrings(s, i, i)
		// Even length
		count += expandFromCenterForPalindromicSubstrings(s, i, i+1)
	}
	return count
}

func expandFromCenterForPalindromicSubstrings(s string, left int, right int) int {
	count := 0
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
		count++
	}
	return count
}
