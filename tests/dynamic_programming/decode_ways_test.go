package dynamic_programming_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/dynamic_programming"
)

func TestNumDecodings(t *testing.T) {
	// Test case for null string
	if result := dynamic_programming.DecodeWays(""); result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}

	// Test case for string with a single character
	if result := dynamic_programming.DecodeWays("1"); result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}

	// Test case for string with multiple characters
	if result := dynamic_programming.DecodeWays("12"); result != 2 {
		t.Errorf("Expected 2, got %d", result)
	}

	// Test case for string with leading zeros
	if result := dynamic_programming.DecodeWays("012"); result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}

	// Test case for string with multiple decoding possibilities
	if result := dynamic_programming.DecodeWays("226"); result != 3 {
		t.Errorf("Expected 3, got %d", result)
	}

	// Test case for string with no valid decodings
	if result := dynamic_programming.DecodeWays("30"); result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}

	// Test case for string with multiple valid decodings
	if result := dynamic_programming.DecodeWays("11106"); result != 2 {
		t.Errorf("Expected 2, got %d", result)
	}
}
