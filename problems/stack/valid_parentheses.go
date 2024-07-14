package stack

func ValidParentheses(s string) bool {
	// Special case
	if len(s) == 0 {
		return true
	}
	// Pairs of brackets
	pairs := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	// Stack to keep track of left parentheses
	stack := []rune{}
	// Process the string
	for _, c := range s {
		// Add left brackets to the stack
		if _, ok := pairs[c]; ok {
			stack = append(stack, c)
		} else if len(stack) == 0 || pairs[stack[len(stack)-1]] != c {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
