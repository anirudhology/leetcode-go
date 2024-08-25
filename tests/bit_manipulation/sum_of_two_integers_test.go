package bit_manipulation_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/bit_manipulation"
)

func TestGetSumBothPositive(t *testing.T) {
	if got := bit_manipulation.SumOfTwoIntegers(2, 3); got != 5 {
		t.Errorf("getSum(2, 3) = %d; want 5", got)
	}
}

func TestGetSumOnePositiveOneNegative(t *testing.T) {
	if got := bit_manipulation.SumOfTwoIntegers(3, -2); got != 1 {
		t.Errorf("getSum(3, -2) = %d; want 1", got)
	}
}

func TestGetSumBothNegative(t *testing.T) {
	if got := bit_manipulation.SumOfTwoIntegers(-2, -3); got != -5 {
		t.Errorf("getSum(-2, -3) = %d; want -5", got)
	}
}

func TestGetSumFirstZero(t *testing.T) {
	if got := bit_manipulation.SumOfTwoIntegers(0, 4); got != 4 {
		t.Errorf("getSum(0, 4) = %d; want 4", got)
	}
}

func TestGetSumSecondZero(t *testing.T) {
	if got := bit_manipulation.SumOfTwoIntegers(-4, 0); got != -4 {
		t.Errorf("getSum(-4, 0) = %d; want -4", got)
	}
}

func TestGetSumBothZero(t *testing.T) {
	if got := bit_manipulation.SumOfTwoIntegers(0, 0); got != 0 {
		t.Errorf("getSum(0, 0) = %d; want 0", got)
	}
}
