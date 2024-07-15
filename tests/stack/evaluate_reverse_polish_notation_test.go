package stack_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/stack"
)

func TestEvaluateReversePolishNotation(t *testing.T) {
	tests := []struct {
		name     string
		tokens   []string
		expected int
	}{
		{
			name:     "Basic Addition",
			tokens:   []string{"2", "1", "+", "3", "*"},
			expected: 9, // (2 + 1) * 3 = 9
		},
		{
			name:     "Basic Subtraction",
			tokens:   []string{"4", "13", "5", "/", "+"},
			expected: 6, // 4 + (13 / 5) = 6
		},
		{
			name:     "Basic Multiplication",
			tokens:   []string{"2", "3", "*", "4", "+"},
			expected: 10, // (2 * 3) + 4 = 10
		},
		{
			name:     "Basic Division",
			tokens:   []string{"10", "6", "/", "9", "+"},
			expected: 10, // (10 / 6) + 9 = 10
		},
		{
			name:     "Single Addition",
			tokens:   []string{"2", "3", "+"},
			expected: 5, // 2 + 3 = 5
		},
		{
			name:     "Single Subtraction",
			tokens:   []string{"5", "3", "-"},
			expected: 2, // 5 - 3 = 2
		},
		{
			name:     "Single Multiplication",
			tokens:   []string{"3", "4", "*"},
			expected: 12, // 3 * 4 = 12
		},
		{
			name:     "Single Division",
			tokens:   []string{"8", "4", "/"},
			expected: 2, // 8 / 4 = 2
		},
		{
			name:     "Negative Numbers",
			tokens:   []string{"-2", "3", "*"},
			expected: -6, // -2 * 3 = -6
		},
		{
			name:     "Negative Result",
			tokens:   []string{"2", "1", "-", "3", "*"},
			expected: 3, // (2 - 1) * 3 = 3
		},
		{
			name:     "Complex Expression",
			tokens:   []string{"3", "4", "+", "2", "*", "7", "/"},
			expected: 2, // (3 + 4) * 2 / 7 = 2
		},
		{
			name:     "Edge Case Single Token",
			tokens:   []string{"1"},
			expected: 1, // Edge case: returns the default value
		},
		{
			name:     "Edge Case Two Tokens",
			tokens:   []string{"1", "2"},
			expected: 2, // Edge case: returns the default value
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := stack.EvaluateReversePolishNotation(test.tokens)
			if result != test.expected {
				t.Errorf("For tokens %v, expected %d, but got %d", test.tokens, test.expected, result)
			}
		})
	}
}
