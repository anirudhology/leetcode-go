package backtracking

func Partition(s string) [][]string {
	// List to store final partitions
	var partitions [][]string
	// Special case
	if len(s) == 0 {
		return partitions
	}
	// Perform backtracking
	backtrackForPartition(s, 0, []string{}, &partitions)
	return partitions
}

func backtrackForPartition(s string, index int, current []string, partitions *[][]string) {
	if index == len(s) {
		temp := make([]string, len(current))
		copy(temp, current)
		*partitions = append(*partitions, temp)
		return
	}
	for i := index; i < len(s); i++ {
		if isPalindrome(s, index, i) {
			current = append(current, s[index:i+1])
			backtrackForPartition(s, i+1, current, partitions)
			current = current[:len(current)-1]
		}
	}
}

func isPalindrome(s string, left int, right int) bool {
	for left <= right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
