package strings

import "strings"

func ValidPalindrome(s string) bool {
	// Special case
	if len(s) == 0 {
		return true
	}
	// Left and right pointers
	left := 0
	right := len(s) - 1
	// Process array from both ends
	for left <= right {
		if !isAlphanumeric(s[left]) {
			left++
			continue
		}
		if !isAlphanumeric(s[right]) {
			right--
			continue
		}
		if strings.ToLower(string(s[left])) != strings.ToLower(string(s[right])) {
			return false
		}
		left++
		right--
	}
	return true
}

func isAlphanumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}
