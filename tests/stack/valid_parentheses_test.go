package stack_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/stack"
)

func TestValidParentheses(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Empty string", "", true},
		{"Single pair of parentheses", "()", true},
		{"Single pair of curly braces", "{}", true},
		{"Single pair of square brackets", "[]", true},
		{"Mixed pairs", "(){}[]", true},
		{"Nested pairs", "({[]})", true},
		{"Incorrect order", "([)]", false},
		{"Single left parenthesis", "(", false},
		{"Single right parenthesis", ")", false},
		{"Left without right", "(((", false},
		{"Right without left", ")))", false},
		{"Mixed correct and incorrect", "{[()()]}", true},
		{"Mixed incorrect and correct", "{[()]}", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := stack.ValidParentheses(test.input)
			if result != test.expected {
				t.Errorf("For input '%s', expected %v, but got %v", test.input, test.expected, result)
			}
		})
	}
}
