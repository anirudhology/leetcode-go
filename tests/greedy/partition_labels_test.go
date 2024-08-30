package greedy_test

import (
	"reflect"
	"testing"

	"github.com/anirudhology/leetcode-go/problems/greedy"
)

func TestPartitionLabels(t *testing.T) {
	// Test string with all unique characters
	if !reflect.DeepEqual(greedy.PartitionLabels("abcd"), []int{1, 1, 1, 1}) {
		t.Errorf("Expected [1, 1, 1, 1] for 'abcd'")
	}

	// Test string with all same characters
	if !reflect.DeepEqual(greedy.PartitionLabels("aaaa"), []int{4}) {
		t.Errorf("Expected [4] for 'aaaa'")
	}

	// Test string with multiple partitions
	if !reflect.DeepEqual(greedy.PartitionLabels("ababcbacadefegdehijhklij"), []int{9, 7, 8}) {
		t.Errorf("Expected [9, 7, 8] for 'ababcbacadefegdehijhklij'")
	}

	// Test string with overlapping partitions
	if !reflect.DeepEqual(greedy.PartitionLabels("eababcbaca"), []int{1, 9}) {
		t.Errorf("Expected [1, 9] for 'eababcbaca'")
	}
}

func TestPartitionLabelEdgeCases(t *testing.T) {
	// Test single character string
	if !reflect.DeepEqual(greedy.PartitionLabels("a"), []int{1}) {
		t.Errorf("Expected [1] for 'a'")
	}

	// Test string with repeating characters
	if !reflect.DeepEqual(greedy.PartitionLabels("abacaba"), []int{7}) {
		t.Errorf("Expected [3, 4] for 'abacaba'")
	}
}
