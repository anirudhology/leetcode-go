package math_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/math"
)

func TestIsHappy(t *testing.T) {
	// Test happy numbers
	if !math.HappyNumber(1) {
		t.Errorf("Expected 1 to be a happy number")
	}
	if !math.HappyNumber(7) {
		t.Errorf("Expected 7 to be a happy number")
	}
	if !math.HappyNumber(19) {
		t.Errorf("Expected 19 to be a happy number")
	}

	// Test unhappy numbers
	if math.HappyNumber(2) {
		t.Errorf("Expected 2 to be an unhappy number")
	}
	if math.HappyNumber(4) {
		t.Errorf("Expected 4 to be an unhappy number")
	}
	if math.HappyNumber(20) {
		t.Errorf("Expected 20 to be an unhappy number")
	}
	if math.HappyNumber(21) {
		t.Errorf("Expected 21 to be an unhappy number")
	}
}
