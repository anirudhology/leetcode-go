package dynamic_programming_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/dynamic_programming"
)

func TestMinDistanceEmptyStrings(t *testing.T) {
	result := dynamic_programming.EditDistance("", "")
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}

func TestMinDistanceOneEmptyString(t *testing.T) {
	if result := dynamic_programming.EditDistance("hello", ""); result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}
	if result := dynamic_programming.EditDistance("", "java"); result != 4 {
		t.Errorf("Expected 4, got %d", result)
	}
}

func TestMinDistanceSameStrings(t *testing.T) {
	if result := dynamic_programming.EditDistance("same", "same"); result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}

func TestMinDistanceDifferentStrings(t *testing.T) {
	if result := dynamic_programming.EditDistance("horse", "ros"); result != 3 {
		t.Errorf("Expected 3, got %d", result)
	}
	if result := dynamic_programming.EditDistance("intention", "execution"); result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}
	if result := dynamic_programming.EditDistance("flaw", "lawn"); result != 2 {
		t.Errorf("Expected 2, got %d", result)
	}
}
