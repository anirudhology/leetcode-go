package math_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/math"
)

func TestMultiplyWithZero(t *testing.T) {
	result := math.MultiplyStrings("0", "1234")
	if result != "0" {
		t.Errorf("Expected 0 but got %s", result)
	}

	result = math.MultiplyStrings("5678", "0")
	if result != "0" {
		t.Errorf("Expected 0 but got %s", result)
	}
}

func TestMultiplySingleDigitNumbers(t *testing.T) {
	result := math.MultiplyStrings("3", "3")
	if result != "9" {
		t.Errorf("Expected 9 but got %s", result)
	}

	result = math.MultiplyStrings("5", "9")
	if result != "45" {
		t.Errorf("Expected 45 but got %s", result)
	}
}

func TestMultiplyMultiDigitNumbers(t *testing.T) {
	result := math.MultiplyStrings("123", "456")
	if result != "56088" {
		t.Errorf("Expected 56088 but got %s", result)
	}

	result = math.MultiplyStrings("999", "999")
	if result != "998001" {
		t.Errorf("Expected 998001 but got %s", result)
	}
}
