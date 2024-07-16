package stack_test

import (
	"reflect"
	"testing"

	"github.com/anirudhology/leetcode-go/problems/stack"
)

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected []string
	}{
		{"Zero pairs", 0, []string{}},
		{"One pair", 1, []string{"()"}},
		{"Two pairs", 2, []string{"(())", "()()"}},
		{"Three pairs", 3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := stack.GenerateParentheses(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("For input %d, expected %v, but got %v", test.input, test.expected, result)
			}
		})
	}
}
