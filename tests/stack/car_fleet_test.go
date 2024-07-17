package stack_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/stack"
)

func TestCarFleet(t *testing.T) {
	tests := []struct {
		target   int
		position []int
		speed    []int
		expected int
	}{
		{12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3}, 3},
		{10, []int{6, 8}, []int{3, 2}, 2},
		{100, []int{0, 2, 4}, []int{4, 2, 1}, 1},
		{10, []int{}, []int{}, 0},
		{10, nil, nil, 0},
		{10, []int{3, 6, 9}, []int{3, 3, 3}, 3},
		{12, []int{0, 4, 2}, []int{2, 1, 3}, 1},
		{15, []int{5, 10, 1}, []int{1, 2, 1}, 3},
	}

	for _, test := range tests {
		result := stack.CarFleet(test.target, test.position, test.speed)
		if result != test.expected {
			t.Errorf("For target %d, position %v, and speed %v; expected %d but got %d",
				test.target, test.position, test.speed, test.expected, result)
		}
	}
}
