package dynamic_programming

func DecodeWays(s string) int {
	// Special case
	if len(s) == 0 {
		return 0
	}
	// Length of the string and previous and current pointers
	n, previous, current := len(s), 1, 0
	// Process the string from right to left
	for i := n - 1; i >= 0; i-- {
		temp := 0
		if s[i] != '0' {
			temp = previous
		}
		if i < n-1 && (s[i] == '1' || s[i] == '2' && s[i+1] < '7') {
			temp += current
		}
		current = previous
		previous = temp
	}
	return previous
}
