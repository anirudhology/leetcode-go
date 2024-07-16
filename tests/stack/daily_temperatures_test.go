package stack_test

import (
	"reflect"
	"testing"

	"github.com/anirudhology/leetcode-go/problems/stack"
)

func TestDailyTemperatures(t *testing.T) {
	tests := []struct {
		temperatures []int
		expected     []int
	}{
		// Edge case: Empty input
		{nil, nil},
		{[]int{}, []int{}},

		// Edge case: Single temperature
		{[]int{30}, []int{0}},

		// Increasing temperatures
		{[]int{30, 40, 50, 60}, []int{1, 1, 1, 0}},

		// Decreasing temperatures
		{[]int{60, 50, 40, 30}, []int{0, 0, 0, 0}},

		// Mixed temperatures with a peak
		{[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},

		// All same temperatures
		{[]int{50, 50, 50, 50}, []int{0, 0, 0, 0}},

		// Sequential increasing temperatures
		{[]int{30, 31, 32, 33}, []int{1, 1, 1, 0}},

		// Random temperatures with multiple peaks
		{[]int{20, 22, 21, 23, 25, 24, 26, 27}, []int{1, 2, 1, 1, 2, 1, 1, 0}},
	}

	for _, test := range tests {
		result := stack.DailyTemperatures(test.temperatures)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %v, expected %v, but got %v", test.temperatures, test.expected, result)
		}
	}
}
