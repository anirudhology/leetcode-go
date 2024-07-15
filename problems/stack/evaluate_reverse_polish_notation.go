package stack

import "strconv"

func EvaluateReversePolishNotation(tokens []string) int {
	// Stack to store operands
	operands := []int{}
	// Process tokens in the string
	for _, token := range tokens {
		switch token {
		case "+":
			x := operands[len(operands)-1]
			operands = operands[:len(operands)-1]
			y := operands[len(operands)-1]
			operands = operands[:len(operands)-1]
			operands = append(operands, x+y)
		case "-":
			x := operands[len(operands)-1]
			operands = operands[:len(operands)-1]
			y := operands[len(operands)-1]
			operands = operands[:len(operands)-1]
			operands = append(operands, y-x)
		case "*":
			x := operands[len(operands)-1]
			operands = operands[:len(operands)-1]
			y := operands[len(operands)-1]
			operands = operands[:len(operands)-1]
			operands = append(operands, x*y)
		case "/":
			x := operands[len(operands)-1]
			operands = operands[:len(operands)-1]
			y := operands[len(operands)-1]
			operands = operands[:len(operands)-1]
			operands = append(operands, y/x)
		default:
			num, _ := strconv.Atoi(token)
			operands = append(operands, num)
		}
	}
	return operands[len(operands)-1]
}
