package greedy

func ValidParenthesisString(s string) bool {
	// Minimum and maximum count of open parentheses
	minOpen, maxOpen := 0, 0
	// Traverse through the string
	for _, c := range s {

		if c == '(' {
			// If it is a left parenthesis, it means you have one
			// more open parenthesis to worry about
			minOpen++
			maxOpen++
		} else if c == ')' {
			// If it is right parenthesis, it means you have one
			// less open parenthesis to worry about
			minOpen--
			maxOpen--
		} else if c == '*' {
			// If it is *, it means there are three cases
			// 1. If we make it as ), then we have one less open parenthesis to worry about
			// 2. If we make it as (, then we have one more open parentheses
			// to worry about
			// 3. No impact at all
			minOpen--
			maxOpen++
		}

		// There are more closing parentheses, then opening ones
		if maxOpen < 0 {
			return false
		}

		// We can't have negative open parenthesis count
		minOpen = max(minOpen, 0)
	}
	return minOpen == 0
}
