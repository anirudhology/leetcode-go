package math_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/math"
)

func TestMyPowZeroExponent(t *testing.T) {
	result := math.PowXN(2.0, 0)
	expected := 1.0
	if expected != result {
		t.Errorf("Expected %f, but got %f", expected, result)
	}
}

func TestMyPowNegativeExponent(t *testing.T) {
	result := math.PowXN(2.0, -2)
	expected := 0.25
	if expected != result {
		t.Errorf("Expected %f, but got %f", expected, result)
	}
}

func TestMyPowPositiveExponentEven(t *testing.T) {
	result := math.PowXN(2.0, 4)
	expected := 16.0
	if expected != result {
		t.Errorf("Expected %f, but got %f", expected, result)
	}
}

func TestMyPowPositiveExponentOdd(t *testing.T) {
	result := math.PowXN(2.0, 3)
	expected := 8.0
	if expected != result {
		t.Errorf("Expected %f, but got %f", expected, result)
	}
}

func TestMyPowBaseZero(t *testing.T) {
	result := math.PowXN(0.0, 5)
	expected := 0.0
	if expected != result {
		t.Errorf("Expected %f, but got %f", expected, result)
	}
}

func TestMyPowBaseOne(t *testing.T) {
	result := math.PowXN(1.0, 5)
	expected := 1.0
	if expected != result {
		t.Errorf("Expected %f, but got %f", expected, result)
	}
}

func TestMyPowBaseNegativeOneOdd(t *testing.T) {
	result := math.PowXN(-1.0, 3)
	expected := -1.0
	if expected != result {
		t.Errorf("Expected %f, but got %f", expected, result)
	}
}

func TestMyPowBaseNegativeOneEven(t *testing.T) {
	result := math.PowXN(-1.0, 2)
	expected := 1.0
	if expected != result {
		t.Errorf("Expected %f, but got %f", expected, result)
	}
}
