package backtracking

func LetterCombinations(digits string) []string {
	// List to store the combinations
	var combinations []string
	// Special case
	if len(digits) == 0 {
		return combinations
	}
	// Mappings from digits to letters
	mappings := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	// Perform backtracking
	backtrackForLetterCombinations(digits, 0, "", &mappings, &combinations)
	return combinations
}

func backtrackForLetterCombinations(digits string, index int, combination string, mappings *[]string, combinations *[]string) {
	if index == len(digits) {
		*combinations = append(*combinations, combination)
		return
	}
	// Get letters
	letters := (*mappings)[digits[index]-'0']
	for _, letter := range letters {
		backtrackForLetterCombinations(digits, index+1, combination+string(letter), mappings, combinations)
	}
}
