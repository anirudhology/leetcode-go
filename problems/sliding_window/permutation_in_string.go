package sliding_window

func CheckInclusion(s1 string, s2 string) bool {
	// Special case
	if len(s1) > len(s2) {
		return false
	}
	// Array to store the frequencies of characters in string s1
	frequencies := [26]int{}
	for _, c := range s1 {
		frequencies[int(c)-int('a')]++
	}
	// Left and right pointers for the sliding window
	left, right := 0, 0
	// Process the string
	for right < len(s2) {
		frequencies[int(s2[right])-int('a')]--
		for left < len(s2) && frequencies[int(s2[right])-int('a')] < 0 {
			frequencies[int(s2[left])-int('a')]++
			left++
		}
		if right-left+1 == len(s1) {
			return true
		}
		right++
	}
	return false
}
