package strings_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/strings"
)

func TestValidAnagram(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{"anagrams listen/silent", "listen", "silent", true},
		{"anagrams triangle/integral", "triangle", "integral", true},
		{"anagrams anagram/nagaram", "anagram", "nagaram", true},
		{"non-anagrams hello/world", "hello", "world", false},
		{"non-anagrams rat/car", "rat", "car", false},
		{"non-anagrams abcd/abce", "abcd", "abce", false},
		{"different lengths abc/abcd", "abc", "abcd", false},
		{"different lengths a/empty", "a", "", false},
		{"empty strings", "", "", true},
		{"repeated characters aabbcc/abcabc", "aabbcc", "abcabc", true},
		{"repeated characters aabbcc/aabbc", "aabbcc", "aabbc", false},
		{"single characters a/a", "a", "a", true},
		{"single characters a/b", "a", "b", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strings.IsAnagram(tt.s, tt.t); got != tt.want {
				t.Errorf("IsAnagram(%q, %q) = %v, want %v", tt.s, tt.t, got, tt.want)
			}
		})
	}
}
