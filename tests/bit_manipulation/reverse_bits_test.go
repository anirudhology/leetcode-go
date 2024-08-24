package bit_manipulation_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/bit_manipulation"
)

func TestReverseBitsWithZero(t *testing.T) {
	var expected uint32 = 0
	result := bit_manipulation.ReverseBits(0)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestReverseBitsWithOne(t *testing.T) {
	var expected uint32 = 0b10000000000000000000000000000000
	result := bit_manipulation.ReverseBits(1)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestReverseBitsWithMaxInteger(t *testing.T) {
	var expected uint32 = 0xFFFFFFFF
	result := bit_manipulation.ReverseBits(0xFFFFFFFF)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestReverseBitsWithMinInteger(t *testing.T) {
	var expected uint32 = 1
	result := bit_manipulation.ReverseBits(0x80000000)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestReverseBitsWithSpecificNumber(t *testing.T) {
	var expected uint32 = 964176192
	result := bit_manipulation.ReverseBits(43261596)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
