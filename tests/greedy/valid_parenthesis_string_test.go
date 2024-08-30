package greedy_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/greedy"
)

func TestCheckValidString(t *testing.T) {
	// Test null input
	if !greedy.ValidParenthesisString("") {
		t.Errorf("Expected true for empty string")
	}

	// Test valid strings
	if !greedy.ValidParenthesisString("()") {
		t.Errorf("Expected true for ()")
	}
	if !greedy.ValidParenthesisString("(*)") {
		t.Errorf("Expected true for (*)")
	}
	if !greedy.ValidParenthesisString("(*))") {
		t.Errorf("Expected true for (*))")
	}

	// Test invalid strings
	if greedy.ValidParenthesisString(")") {
		t.Errorf("Expected false for )")
	}
	if greedy.ValidParenthesisString(")(") {
		t.Errorf("Expected false for )(")
	}
	if !greedy.ValidParenthesisString("((*)") {
		t.Errorf("Expected false for ((*)")
	}
	if greedy.ValidParenthesisString("((*))(") {
		t.Errorf("Expected false for ((*))(")
	}
}

func TestEdgeCases(t *testing.T) {
	// Test string with only one *
	if !greedy.ValidParenthesisString("*") {
		t.Errorf("Expected true for *")
	}

	// Test string with multiple *
	if !greedy.ValidParenthesisString("***") {
		t.Errorf("Expected true for ***")
	}

	// Test strings with nested parentheses
	if !greedy.ValidParenthesisString("(())") {
		t.Errorf("Expected true for (())")
	}
	if greedy.ValidParenthesisString("(()))") {
		t.Errorf("Expected false for (()))")
	}
}
