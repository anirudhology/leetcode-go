package strings_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/strings"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{[]string{"hello", "world"}, "5/hello5/world"},
		{[]string{"foo", "bar", "baz"}, "3/foo3/bar3/baz"},
		{[]string{"a", "bb", "ccc"}, "1/a2/bb3/ccc"},
		{[]string{}, ""},
		{[]string{""}, "0/"},
		{[]string{"a/b", "c/d", "e/f"}, "3/a/b3/c/d3/e/f"},
	}

	for _, test := range tests {
		result := strings.Encode(test.input)
		if result != test.expected {
			t.Errorf("Encode(%v) = %s; expected %s", test.input, result, test.expected)
		}
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"5/hello5/world", []string{"hello", "world"}},
		{"3/foo3/bar3/baz", []string{"foo", "bar", "baz"}},
		{"1/a2/bb3/ccc", []string{"a", "bb", "ccc"}},
		{"", []string{}},
		{"0/", []string{""}},
		{"3/a/b3/c/d3/e/f", []string{"a/b", "c/d", "e/f"}},
	}

	for _, test := range tests {
		result := strings.Decode(test.input)
		if len(result) != len(test.expected) {
			t.Errorf("Decode(%s) = %v; expected %v", test.input, result, test.expected)
			continue
		}
		for i := range result {
			if result[i] != test.expected[i] {
				t.Errorf("Decode(%s) = %v; expected %v", test.input, result, test.expected)
				break
			}
		}
	}
}

func TestEncodeDecode(t *testing.T) {
	tests := [][]string{
		{"hello", "world"},
		{"foo", "bar", "baz"},
		{"a", "bb", "ccc"},
		{},
		{""},
		{"a/b", "c/d", "e/f"},
	}

	for _, test := range tests {
		encoded := strings.Encode(test)
		decoded := strings.Decode(encoded)
		if len(decoded) != len(test) {
			t.Errorf("EncodeDecode(%v) = %v; expected %v", test, decoded, test)
			continue
		}
		for i := range decoded {
			if decoded[i] != test[i] {
				t.Errorf("EncodeDecode(%v) = %v; expected %v", test, decoded, test)
				break
			}
		}
	}
}
