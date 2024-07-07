package strings

func IsAnagram(s string, t string) bool {
	// Special case
	if len(s) != len(t) {
		return false
	}
	// Array to store the frequencies of characters
	var frequencies [26]int
	// Process each character of both strings
	for i := 0; i < len(s); i++ {
		frequencies[int(s[i])-int('a')]++
		frequencies[int(t[i])-int('a')]--
	}
	// Check if there is any non zero frequency
	for _, frequency := range frequencies {
		if frequency != 0 {
			return false
		}
	}
	return true
}
