package bit_manipulation_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/bit_manipulation"
)

func TestHammingWeightZero(t *testing.T) {
	result := bit_manipulation.NumberOf1Bits(0)
	if result != 0 {
		t.Errorf("Expected 0, but got %d", result)
	}
}

func TestHammingWeightOne(t *testing.T) {
	result := bit_manipulation.NumberOf1Bits(1)
	if result != 1 {
		t.Errorf("Expected 1, but got %d", result)
	}
}

func TestHammingWeightLargeNumber(t *testing.T) {
	result := bit_manipulation.NumberOf1Bits(11) // binary 1011
	if result != 3 {
		t.Errorf("Expected 3, but got %d", result)
	}
}

func TestHammingWeightPowerOfTwo(t *testing.T) {
	result := bit_manipulation.NumberOf1Bits(16) // binary 10000
	if result != 1 {
		t.Errorf("Expected 1, but got %d", result)
	}
}

func TestHammingWeightMixedBits(t *testing.T) {
	result := bit_manipulation.NumberOf1Bits(29) // binary 11101
	if result != 4 {
		t.Errorf("Expected 4, but got %d", result)
	}
}
