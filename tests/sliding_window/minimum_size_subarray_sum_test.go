package sliding_window_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/sliding_window"
)

func TestMinSubArrayLen_EmptyArray(t *testing.T) {
	result := sliding_window.MinimumSizeSubarraySum(7, []int{})
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}

func TestMinSubArrayLen_NullArray(t *testing.T) {
	result := sliding_window.MinimumSizeSubarraySum(7, nil)
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}

func TestMinSubArrayLen_NoSubarrayExists(t *testing.T) {
	result := sliding_window.MinimumSizeSubarraySum(15, []int{1, 2, 3, 4, 5})
	if result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}
}

func TestMinSubArrayLen_ExactMatch(t *testing.T) {
	result := sliding_window.MinimumSizeSubarraySum(7, []int{2, 3, 1, 2, 4, 3})
	if result != 2 {
		t.Errorf("Expected 2, got %d", result)
	}
}

func TestMinSubArrayLen_MultipleMatches(t *testing.T) {
	result := sliding_window.MinimumSizeSubarraySum(4, []int{1, 4, 4})
	if result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}
}

func TestMinSubArrayLen_SingleElement(t *testing.T) {
	result := sliding_window.MinimumSizeSubarraySum(3, []int{3})
	if result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}
}
