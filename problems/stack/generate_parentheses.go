package stack

func GenerateParentheses(n int) []string {
	// List to store the combinations
	combinations := []string{}
	// Special case
	if n <= 0 {
		return combinations
	}
	// Stack for storing a combination
	stack := []rune{}
	generate(n, 0, 0, stack, &combinations)
	return combinations
}

func generate(n int, left int, right int, stack []rune, combinations *[]string) {
	// Base case
	if left == right && left == n {
		combination := string(stack)
		*combinations = append(*combinations, combination)
		return
	}
	// Add left parentheses if possible
	if left < n {
		stack = append(stack, '(')
		generate(n, left+1, right, stack, combinations)
		stack = stack[:len(stack)-1]
	}
	// Match the right parentheses
	if right < left {
		stack = append(stack, ')')
		generate(n, left, right+1, stack, combinations)
		stack = stack[:len(stack)-1]
	}
}
