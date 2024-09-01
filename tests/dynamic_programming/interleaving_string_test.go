package dynamic_programming_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/dynamic_programming"
)

func TestIsInterleave(t *testing.T) {

	if !dynamic_programming.InterleavingString("aabcc", "dbbca", "aadbbcbcac") {
		t.Error("Expected true, but got false")
	}

	if dynamic_programming.InterleavingString("aabcc", "dbbca", "aadbbbaccc") {
		t.Error("Expected false, but got true")
	}

	if !dynamic_programming.InterleavingString("", "", "") {
		t.Error("Expected true, but got false")
	}

	if !dynamic_programming.InterleavingString("a", "", "a") {
		t.Error("Expected true, but got false")
	}

	if !dynamic_programming.InterleavingString("", "b", "b") {
		t.Error("Expected true, but got false")
	}

	if dynamic_programming.InterleavingString("abc", "def", "abdccfe") {
		t.Error("Expected false, but got true")
	}
}

func TestEdgeCases(t *testing.T) {
	if dynamic_programming.InterleavingString("abc", "def", "abcdefg") {
		t.Error("Expected false, but got true")
	}

	if !dynamic_programming.InterleavingString("abc", "def", "adbcef") {
		t.Error("Expected true, but got false")
	}

	if dynamic_programming.InterleavingString("", "a", "b") {
		t.Error("Expected false, but got true")
	}
}
